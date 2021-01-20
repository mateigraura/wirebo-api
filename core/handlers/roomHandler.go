package handlers

import (
	"github.com/mateigraura/wirebo-api/domain"
)

type RoomHandler struct {
	broadcast chan *domain.Message
}
