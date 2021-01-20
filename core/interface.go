package core

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/domain"
)

type UserRepository interface {
	GetById(id uuid.UUID) (domain.User, error)

	GetClaims(id uuid.UUID) (domain.Authorization, error)

	Insert(user *domain.User) error

	InsertClaims(payload *domain.Authorization) error
}

type RoomRepository interface {
	GetRoomsFor(userId uuid.UUID) ([]domain.Room, error)

	GetUsersInRoom(room domain.Room) (domain.Room, error)

	Insert(room *domain.Room) error

	InsertMapping(values []interface{}) error
}

type MessageRepository interface {
	GetByRoomId(roomId uuid.UUID) ([]domain.Message, error)

	Insert(message *domain.Message) error
}
