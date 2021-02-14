package repository

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/models"
	"github.com/mateigraura/wirebo-api/storage"
)

type KeyMapRepositoryImpl struct {
}

func (kmr *KeyMapRepositoryImpl) Get(id uuid.UUID) (models.KeyMapping, error) {
	conn := storage.Connection()

	keyMapping := new(models.KeyMapping)
	err := conn.Model(keyMapping).
		Where("id = ?", id).
		Select()

	return *keyMapping, err
}

func (kmr *KeyMapRepositoryImpl) Update(keyMapping *models.KeyMapping) error {
	return nil
}

func (kmr *KeyMapRepositoryImpl) Insert(keyMapping *models.KeyMapping) error {
	conn := storage.Connection()

	_, err := conn.Model(keyMapping).
		Returning("id").
		Insert()

	return err
}
