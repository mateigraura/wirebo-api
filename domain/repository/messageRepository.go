package repository

import (
	"github.com/mateigraura/wirebo-api/domain"
	"github.com/mateigraura/wirebo-api/storage"
)

type MessageRepositoryImpl struct {
}

func (m *MessageRepositoryImpl) Insert(message *domain.Message) error {
	conn := storage.Connection()

	_, err := conn.Model(message).
		Returning("id").
		Insert()

	return err
}
