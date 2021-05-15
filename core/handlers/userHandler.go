package handlers

import (
	"errors"

	"github.com/mateigraura/wirebo-api/core"
	"github.com/mateigraura/wirebo-api/models"
)

type UserHandler struct {
	userRepository core.UserRepository
}

func NewUserHandler(userRepository core.UserRepository) UserHandler {
	return UserHandler{
		userRepository: userRepository,
	}
}

func (uh *UserHandler) Search(query string) ([]models.User, error) {
	results, err := uh.userRepository.Search(query)
	if err != nil {
		return []models.User{}, errors.New("no results found for query")
	}
	return results, err
}
