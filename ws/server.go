package ws

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/core"
	"github.com/mateigraura/wirebo-api/core/converters"
	"github.com/mateigraura/wirebo-api/models"
)

var ctx = context.Background()

type server struct {
	clients           map[uuid.UUID]*Client
	roomHandlers      map[uuid.UUID]*RoomHandler
	register          chan *Client
	unregister        chan *Client
	roomRepository    core.RoomRepository
	messageRepository core.MessageRepository
}

type ServerArgs struct {
	RoomRepository    core.RoomRepository
	MessageRepository core.MessageRepository
}

func NewWsServer(args ServerArgs) *server {
	return &server{
		clients:           make(map[uuid.UUID]*Client),
		roomHandlers:      make(map[uuid.UUID]*RoomHandler),
		register:          make(chan *Client),
		unregister:        make(chan *Client),
		roomRepository:    args.RoomRepository,
		messageRepository: args.MessageRepository,
	}
}

func (s *server) Run() {
	for {
		select {
		case client := <-s.register:
			s.registerClient(client)

		case client := <-s.unregister:
			s.unregisterClient(client)
		}
	}
}

func (s *server) registerClient(client *Client) {
	if _, ok := s.clients[client.id]; ok {
		log.Println("abort. client already connected", client.id)
		return
	}
	s.clients[client.id] = client
	rooms, err := s.roomRepository.GetRoomsFor(client.id)
	if err != nil {
		log.Println(err)
		return
	}

	for _, room := range rooms {
		roomHandler, ok := s.findRoomHandler(room.Id)
		if ok {
			roomHandler.registerClient(client)
		} else {
			newRoomHandler := NewRoomHandler(room.Id)
			s.roomHandlers[room.Id] = newRoomHandler
			newRoomHandler.registerClient(client)
			go newRoomHandler.RunRoomHandler()
		}
	}
	log.Println("connection established for client", client.id)
}

func (s *server) unregisterClient(client *Client) {
	if _, ok := s.clients[client.id]; ok {
		delete(s.clients, client.id)
	}

	rooms, err := s.roomRepository.GetRoomsFor(client.id)
	if err != nil {
		log.Println(err)
	}

	for _, room := range rooms {
		roomHandler, ok := s.findRoomHandler(room.Id)
		if ok {
			roomHandler.unregisterClient(client)
		}
	}
	log.Println("connection lost for client", client.id)
}

func (s *server) handleMessage(message models.Message) {
	roomHandler, ok := s.findRoomHandler(message.RoomId)
	if ok {
		message.CreatedAt = time.Now().UTC()
		err := s.messageRepository.Insert(&message)
		if err != nil {
			log.Println(fmt.Sprintf("wsServer: message insertion %s", err.Error()))
			// TODO: push msg to msg_queue and broadcast, rather than returning
			return
		}
		msgBytes, err := converters.Marshal(message)
		if err != nil {
			log.Println(fmt.Sprintf("wsServer: message marshaling %s", err.Error()))
			return
		}
		roomHandler.broadcast <- msgBytes
	}
}

func (s *server) findRoomHandler(roomId uuid.UUID) (*RoomHandler, bool) {
	if room, ok := s.roomHandlers[roomId]; ok {
		return room, room != nil
	}

	return nil, false
}
