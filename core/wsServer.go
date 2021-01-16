package core

import (
	"fmt"

	"github.com/google/uuid"
)

const pubSubChannel = "pub-sub-chan"

type WsServer struct {
	clients    map[uuid.UUID]bool
	rooms      map[uuid.UUID]bool
	register   chan *WsClient
	unregister chan *WsClient
}

func NewWsServer() *WsServer {
	return &WsServer{
		clients:    make(map[uuid.UUID]bool),
		rooms:      make(map[uuid.UUID]bool),
		register:   make(chan *WsClient),
		unregister: make(chan *WsClient),
	}
}

func (w *WsServer) Run() {
	for {
		select {
		case client := <-w.register:
			w.registerClient(client)

		case client := <-w.unregister:
			w.unregisterClient(client)
		}
	}
}

func (w *WsServer) registerClient(client *WsClient) {
	w.clients[client.Id] = true
	fmt.Printf("Welcome %s\n", client.Name)
}

func (w *WsServer) unregisterClient(client *WsClient) {
	if _, ok := w.clients[client.Id]; ok {
		delete(w.clients, client.Id)
	}
	fmt.Printf("See you %s\n", client.Name)
}

func (w *WsServer) findRoomById(id uuid.UUID) *RoomHandler {
	return nil
}
