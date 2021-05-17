package core

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/models"
)

type UserRepository interface {
	GetById(id uuid.UUID) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetClaims(id uuid.UUID) (models.Authorization, error)
	UpdateClaims(claims *models.Authorization) error
	Insert(user *models.User) error
	InsertClaims(payload *models.Authorization) error
	Update(user *models.User) error
	Search(input string) ([]models.User, error)
}

type RoomRepository interface {
	GetRoomByName(name string) (models.Room, error)
	GetRoomsFor(userId uuid.UUID) ([]models.Room, error)
	GetUsersInRoom(room *models.Room) error
	Insert(room *models.Room) error
	InsertMapping(values []*models.UserRoom) error
}

type MessageRepository interface {
	GetByRoomId(roomId uuid.UUID) ([]models.Message, error)
	Insert(message *models.Message) error
}

type KeyMapperRepository interface {
	Get(id uuid.UUID) (models.KeyMapping, error)
	Update(keyMapping *models.KeyMapping) error
	Insert(keyMapping *models.KeyMapping) error
}

type AvatarRepository interface {
	GetByHash(hash string) (models.Avatar, error)
	Insert(avatar *models.Avatar) error
}
