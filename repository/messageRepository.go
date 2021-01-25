package repository

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/models"
	"github.com/mateigraura/wirebo-api/storage"
)

type MessageRepositoryImpl struct {
}

func (m *MessageRepositoryImpl) GetByRoomId(roomId uuid.UUID) ([]models.Message, error) {
	conn := storage.Connection()

	var messages []models.Message
	err := conn.Model(&messages).
		Where(`"message"."room_id" = ?`, roomId).
		Relation("Sender").
		Select()

	if err != nil {
		return []models.Message{}, err
	}

	return messages, nil
}

func (m *MessageRepositoryImpl) Insert(message *models.Message) error {
	conn := storage.Connection()

	_, err := conn.Model(message).
		Returning("id").
		Insert()

	return err
}
