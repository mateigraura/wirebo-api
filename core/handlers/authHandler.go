package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/mateigraura/wirebo-api/crypto/authorization"

	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/core"
	"github.com/mateigraura/wirebo-api/models"
)

type AuthHandler struct {
	userRepository core.UserRepository
}

func NewAuthHandler(userRepository core.UserRepository) AuthHandler {
	return AuthHandler{
		userRepository: userRepository,
	}
}

func (ah *AuthHandler) Register(request models.RegisterRequest) (bool, error) {
	pswHash, err := authorization.HashPassword(request.Password)
	if err != nil {
		return false, err
	}

	user := models.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: pswHash,
		Username:     request.Username,
		AvatarUrl:    "//avatar",
	}

	if err = ah.userRepository.Insert(&user); err != nil {
		return false, err
	}

	return true, nil
}

func (ah *AuthHandler) Login(request models.LoginRequest) (string, error) {
	user, err := ah.userRepository.GetByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if !authorization.CheckEqual(request.Password, user.PasswordHash) {
		return "", errors.New("invalid password")
	}

	authorization, err := makeAuthorization(user.Id)
	if err != nil {
		return "", err
	}

	// TODO: refactor
	_, err = ah.userRepository.GetClaims(user.Id)
	if err != nil {
		err = ah.userRepository.InsertClaims(&authorization)
		return authorization.Token, err
	}

	err = ah.userRepository.UpdateClaims(&authorization)

	return authorization.Token, err
}

func (ah *AuthHandler) Refresh(jwtToken string) (string, error) {
	tokenClaims, err := authorization.GetClaims(jwtToken, false)
	if err != nil {
		return "", err
	}

	userId, _ := uuid.Parse(tokenClaims.Id)
	claims, err := ah.userRepository.GetClaims(userId)
	if err != nil {
		return "", err
	}

	refreshFromJwt := sha256.Sum256([]byte(jwtToken))
	refreshFromClaims, _ := hex.DecodeString(claims.RefreshToken)
	if !bytes.Equal(refreshFromJwt[:], refreshFromClaims) {
		return "", errors.New("invalid refresh token")
	}

	authorization, err := makeAuthorization(claims.OwnerId)
	if err != nil {
		return "", err
	}
	err = ah.userRepository.UpdateClaims(&authorization)

	return authorization.Token, err
}

func makeAuthorization(ownerId uuid.UUID) (models.Authorization, error) {
	jwt, err := authorization.GenerateJwt(ownerId.String())
	if err != nil {
		return models.Authorization{}, err
	}

	refreshToken := sha256.Sum256([]byte(jwt))

	return models.Authorization{
		Token:        jwt,
		RefreshToken: hex.EncodeToString(refreshToken[:]),
		OwnerId:      ownerId,
	}, nil
}
