package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/models"
	"github.com/mateigraura/wirebo-api/storage"
)

type UserRepositoryImpl struct {
}

func (ur *UserRepositoryImpl) GetById(id uuid.UUID) (models.User, error) {
	conn := storage.Connection()

	user := new(models.User)
	err := conn.Model(user).
		Where("id = ?", id).
		Select()

	return *user, err
}

func (ur *UserRepositoryImpl) GetByEmail(email string) (models.User, error) {
	conn := storage.Connection()

	user := new(models.User)
	err := conn.Model(user).
		Where("email = ?", email).
		Select()

	return *user, err
}

func (ur *UserRepositoryImpl) GetClaims(id uuid.UUID) (models.Authorization, error) {
	conn := storage.Connection()

	authPayload := new(models.Authorization)
	err := conn.Model(authPayload).
		Where(`"authorization"."owner_id" = ?`, id).
		Select()

	return *authPayload, err
}

func (ur *UserRepositoryImpl) UpdateClaims(claims *models.Authorization) error {
	conn := storage.Connection()
	_, err := conn.Model(claims).
		Where(`"authorization"."owner_id" = ?`, claims.OwnerId).
		Update()

	return err
}

func (ur *UserRepositoryImpl) Insert(user *models.User) error {
	conn := storage.Connection()
	_, err := conn.Model(user).
		Returning("id").
		Insert()

	return err
}

func (ur *UserRepositoryImpl) InsertClaims(payload *models.Authorization) error {
	conn := storage.Connection()
	_, err := conn.Model(payload).
		Returning("id").
		Insert()

	return err
}

func (ur *UserRepositoryImpl) Update(user *models.User) error {
	conn := storage.Connection()
	_, err := conn.Model(user).
		Where(`"user"."id" = ?`, user.Id).
		Update()

	return err
}

func (ur *UserRepositoryImpl) Search(input string) ([]models.User, error) {
	conn := storage.Connection()
	users := new([]models.User)
	err := conn.Model(users).
		Where(
			`"user"."username" like ?`,
			fmt.Sprintf("%%%s%%", input),
		).
		Select()

	return *users, err
}
