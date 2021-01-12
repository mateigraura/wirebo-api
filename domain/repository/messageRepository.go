package repository

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/domain"
	"github.com/mateigraura/wirebo-api/storage"
)

type MessageRepositoryImpl struct {
}

func (m *MessageRepositoryImpl) GetByRoomId(roomId uuid.UUID) ([]domain.Message, error) {
	conn := storage.Connection()

	var messages []domain.Message
	err := conn.Model(&messages).
		Where(`"message"."room_id" = ?`, roomId).
		Relation("Sender").
		Select()

	if err != nil {
		return []domain.Message{}, err
	}

	return messages, nil
}

func (m *MessageRepositoryImpl) Insert(message *domain.Message) error {
	conn := storage.Connection()

	_, err := conn.Model(message).
		Returning("id").
		Insert()

	return err
}
