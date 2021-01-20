package repository

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/domain"
	"github.com/mateigraura/wirebo-api/storage"
)

type UserRepositoryImpl struct {
}

func (ur *UserRepositoryImpl) GetById(id uuid.UUID) (domain.User, error) {
	conn := storage.Connection()

	user := new(domain.User)
	err := conn.Model(user).
		Where("id = ?", id).
		Select()

	if err != nil {
		return domain.User{}, err
	}

	return *user, nil
}

func (ur *UserRepositoryImpl) GetClaims(id uuid.UUID) (domain.Authorization, error) {
	conn := storage.Connection()

	var authPayload domain.Authorization
	err := conn.Model(&authPayload).
		Where(`"authorization"."owner_id" = ?`, id).
		Select()

	if err != nil {
		return domain.Authorization{}, err
	}

	return authPayload, nil
}

func (ur *UserRepositoryImpl) Insert(user *domain.User) error {
	conn := storage.Connection()

	_, err := conn.Model(user).
		Returning("id").
		Insert()

	return err
}

func (ur *UserRepositoryImpl) InsertClaims(payload *domain.Authorization) error {
	conn := storage.Connection()

	_, err := conn.Model(payload).
		Returning("id").
		Insert()

	return err
}
