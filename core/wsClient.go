package core

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/mateigraura/wirebo-api/domain"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMsgSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WsClient struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	send     chan []byte
	conn     *websocket.Conn
	rooms    map[*Room]bool // roomId fits better ?
	wsServer *WsServer
}

func NewClient(conn *websocket.Conn, wsServer *WsServer, id, name string) *WsClient {
	_id, _ := uuid.Parse(id)

	return &WsClient{
		Id:       _id,
		Name:     name,
		conn:     conn,
		send:     make(chan []byte, 256),
		rooms:    make(map[*Room]bool),
		wsServer: wsServer,
	}
}

func ServeWs(wsServer *WsServer, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, wsServer, "user_id", "user_name")
	client.wsServer.register <- client

	go client.writePump()
	go client.readPump()
}

func (c *WsClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	for {
		select {
		case msg, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err)
				return
			}

			_, _ = w.Write(msg)

			for i := 0; i < len(c.send); i++ {
				_, _ = w.Write(newline)
				_, _ = w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				log.Println(err)
				return
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (c *WsClient) readPump() {
	defer func() {
		c.wsServer.unregister <- c
		if err := c.conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	c.conn.SetReadLimit(maxMsgSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Println(err)
			}
			break
		}

		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		c.handleNewMessage(msg)
	}
}

func (c *WsClient) handleNewMessage(msg []byte) {
	var message domain.Message
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(err)
		return
	}

	// switch actions
	if room := c.wsServer.findRoomById(message.RoomId); room != nil {
		room.broadcast <- &message
	}
}
