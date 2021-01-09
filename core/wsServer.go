package core

import "github.com/google/uuid"

type WsServer struct {
	register   chan *WsClient
	unregister chan *WsClient
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

}

func (w *WsServer) unregisterClient(client *WsClient) {

}

func (w *WsServer) findRoomById(id uuid.UUID) *Room {
	return nil
}
