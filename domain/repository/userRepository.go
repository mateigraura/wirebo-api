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

func (ur *UserRepositoryImpl) Insert(user domain.User) (uuid.UUID, error) {
	conn := storage.Connection()

	_, err := conn.Model(&user).
		Returning("id").
		Insert()

	if err != nil {
		return [16]byte{}, err
	}

	return user.Id, nil
}
