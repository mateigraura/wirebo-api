package core

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/domain"
)

type UserRepository interface {
	GetById(id uuid.UUID) (domain.User, error)
	Insert(user domain.User) (uuid.UUID, error)
}
