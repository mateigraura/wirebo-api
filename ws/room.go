package ws

import (
	"log"

	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/core/converters"
	"github.com/mateigraura/wirebo-api/models"
	"github.com/mateigraura/wirebo-api/storage"
)

type RoomHandler struct {
	id        uuid.UUID
	clients   map[uuid.UUID]*Client
	broadcast chan []byte
}

func NewRoomHandler(id uuid.UUID) *RoomHandler {
	return &RoomHandler{
		id:        id,
		clients:   make(map[uuid.UUID]*Client),
		broadcast: make(chan []byte),
	}
}

func (r *RoomHandler) RunRoomHandler() {
	go r.subscribeToPubSub()

	for {
		select {
		case message := <-r.broadcast:
			r.publishToPubSub(message)
		}
	}
}

func (r *RoomHandler) subscribeToPubSub() {
	pubsub := storage.Redis.Subscribe(ctx, r.id.String())
	ch := pubsub.Channel()

	for msg := range ch {
		r.pushToClients([]byte(msg.Payload))
	}
}

func (r *RoomHandler) publishToPubSub(message []byte) {
	err := storage.Redis.Publish(ctx, r.id.String(), message).Err()
	if err != nil {
		log.Println(err)
	}
}

func (r *RoomHandler) pushToClients(message []byte) {
	msg := models.Message{}
	err := converters.Unmarshal(message, &msg)
	if err != nil {
		return
	}
	for _, client := range r.clients {
		if client.id != msg.SenderId {
			client.send <- message
		}
	}
}

func (r *RoomHandler) registerClient(client *Client) {
	r.clients[client.id] = client
}

func (r *RoomHandler) unregisterClient(client *Client) {
	if _, ok := r.clients[client.id]; ok {
		delete(r.clients, client.id)
	}
}
