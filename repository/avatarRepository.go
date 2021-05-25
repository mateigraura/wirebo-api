package repository

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/models"
	"github.com/mateigraura/wirebo-api/storage"
	"time"
)

type AvatarRepository struct {
}

func (ar *AvatarRepository) GetByUserId(userId uuid.UUID) (models.Avatar, error) {
	conn := storage.Connection()
	avatar := new(models.Avatar)
	err := conn.Model(avatar).Where(`"avatar"."user_id" = ?`, userId).Select()
	return *avatar, err
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

func (ar *AvatarRepository) Update(avatar *models.Avatar) error {
	conn := storage.Connection()
	avatar.UpdatedAt = time.Now().UTC()
	_, err := conn.Model(avatar).
		Where(`"avatar"."user_id" = ?`, avatar.UserId).
		Update()

	return err
}
