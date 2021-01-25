package handlers

import (
	"github.com/mateigraura/wirebo-api/models"
)

type RoomHandler struct {
	broadcast chan *models.Message
}
