package ws

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/mateigraura/wirebo-api/core/converters"
	"github.com/mateigraura/wirebo-api/models"
)

const (
	MESSAGE = "message"
	TYPING  = "typing"
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
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type EventMessage struct {
	Action  string         `json:"action"`
	Message models.Message `json:"message"`
}

type Client struct {
	id       uuid.UUID
	send     chan []byte
	conn     *websocket.Conn
	wsServer *server
}

func NewClient(conn *websocket.Conn, wsServer *server, id uuid.UUID) *Client {
	return &Client{
		id:       id,
		conn:     conn,
		send:     make(chan []byte, 256),
		wsServer: wsServer,
	}
}

func ServeWs(wsServer *server, c *gin.Context) {
	id := c.Param("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, wsServer, parsedId)
	client.wsServer.register <- client

	go client.writePump()
	go client.readPump()
}

// writePump listens on the send channel and pushes data on the ws stream
func (c *Client) writePump() {
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

			if err = w.Close(); err != nil {
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

// readPump listens for incoming events and reads the content
func (c *Client) readPump() {
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
		c.pushNewMessage(msg)
	}
}

func (c *Client) pushNewMessage(eventBytes []byte) {
	var event EventMessage
	if err := converters.Unmarshal(eventBytes, &event); err != nil {
		log.Println(fmt.Sprintf("read event unmarshal err %s", err.Error()))
		return
	}

	switch event.Action {
	case MESSAGE:
		c.wsServer.handleMessage(event.Message)
	}
}
