package ws

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/core"
	"github.com/mateigraura/wirebo-api/core/converters"
	"github.com/mateigraura/wirebo-api/models"
)

const pubSubChannel = "pub-sub-chan"

var ctx = context.Background()

type Server struct {
	clients        map[uuid.UUID]*Client
	roomHandlers   map[uuid.UUID]*RoomHandler
	register       chan *Client
	unregister     chan *Client
	roomRepository core.RoomRepository
}

func NewWsServer(roomRepository core.RoomRepository) *Server {
	return &Server{
		clients:        make(map[uuid.UUID]*Client),
		roomHandlers:   make(map[uuid.UUID]*RoomHandler),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		roomRepository: roomRepository,
	}
}

func (s *Server) Run() {
	for {
		select {
		case client := <-s.register:
			s.registerClient(client)

		case client := <-s.unregister:
			s.unregisterClient(client)
		}
	}
}

func (s *Server) registerClient(client *Client) {
	s.clients[client.id] = client

	rooms, err := s.roomRepository.GetRoomsFor(client.id)
	if err != nil {
		log.Println(err)
	}

	for _, room := range rooms {
		roomHandler, ok := s.findRoomHandler(room.Id)
		if ok {
			roomHandler.registerClient(client)
		} else {
			newRoomHandler := NewRoomHandler(room.Id)
			go newRoomHandler.RunRoomHandler()
		}
	}
	log.Printf("Client %s connected\n", client.id)
}

func (s *Server) unregisterClient(client *Client) {
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
	log.Printf("Client %s disconnected\n", client.id)
}

func (s *Server) handleMessage(message models.Message) {
	roomHandler, ok := s.findRoomHandler(message.RoomId)
	if ok {
		msgBytes, err := converters.Marshal(message)
		if err != nil {
			log.Println(err)
			return
		}
		roomHandler.broadcast <- msgBytes
	}
}

func (s *Server) findRoomHandler(roomId uuid.UUID) (*RoomHandler, bool) {
	if room, ok := s.roomHandlers[roomId]; ok {
		return room, true
	}

	return nil, false
}
