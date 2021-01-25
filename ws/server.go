package ws

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/core/handlers"
)

const pubSubChannel = "pub-sub-chan"

type Server struct {
	clients    map[uuid.UUID]bool
	rooms      map[uuid.UUID]bool
	register   chan *Client
	unregister chan *Client
}

func NewWsServer() *Server {
	return &Server{
		clients:    make(map[uuid.UUID]bool),
		rooms:      make(map[uuid.UUID]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (w *Server) Run() {
	for {
		select {
		case client := <-w.register:
			w.registerClient(client)

		case client := <-w.unregister:
			w.unregisterClient(client)
		}
	}
}

func (w *Server) registerClient(client *Client) {
	w.clients[client.Id] = true
	fmt.Printf("Welcome %s\n", client.Name)
}

func (w *Server) unregisterClient(client *Client) {
	if _, ok := w.clients[client.Id]; ok {
		delete(w.clients, client.Id)
	}
	fmt.Printf("See you %s\n", client.Name)
}

func (w *Server) findRoomById(id uuid.UUID) *handlers.RoomHandler {
	return nil
}
