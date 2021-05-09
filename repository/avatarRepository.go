package repository

import (
	"github.com/mateigraura/wirebo-api/models"
	"github.com/mateigraura/wirebo-api/storage"
)

type AvatarRepository struct {
}

func (ar *AvatarRepository) GetByHash(hash string) (models.Avatar, error) {
	conn := storage.Connection()
	avatar := new(models.Avatar)
	err := conn.Model(avatar).Where(`"avatar"."hash" = ?`, hash).Select()
	return *avatar, err
}

func (ar *AvatarRepository) Insert(avatar *models.Avatar) error {
	conn := storage.Connection()
	_, err := conn.Model(avatar).Returning("id").Insert()
	return err
}
