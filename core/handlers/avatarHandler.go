package handlers

import (
	"encoding/hex"
	"io"

	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/core"
	"github.com/mateigraura/wirebo-api/core/converters"
	"github.com/mateigraura/wirebo-api/crypto"
	"github.com/mateigraura/wirebo-api/models"
)

type AvatarHandler struct {
	avatarRepository core.AvatarRepository
	userRepository   core.UserRepository
	hasher           crypto.Hasher
}

func NewAvatarHandler(
	avatarRepo core.AvatarRepository,
	userRepo core.UserRepository,
	hasher crypto.Hasher,
) AvatarHandler {
	return AvatarHandler{
		avatarRepository: avatarRepo,
		userRepository:   userRepo,
		hasher:           hasher,
	}
}

func (ah *AvatarHandler) Save(file io.Reader, userId string) (string, error) {
	imgBytes, err := converters.ResizeFromFile(file, 0)
	if err != nil {
		return "", InvalidInputProvidedErr
	}

	userIdUuid, err := uuid.Parse(userId)
	if err != nil {
		return "", InvalidInputProvidedErr
	}

	avatar := &models.Avatar{
		UserId:  userIdUuid,
		Content: imgBytes,
	}
	hash, err := ah.hasher.HashObj(avatar)
	if err != nil {
		return "", SerializationErr
	}

	avatar.Hash = hex.EncodeToString(hash[:])
	_, err = ah.avatarRepository.GetByUserId(avatar.UserId)
	if err != nil {
		if err = ah.avatarRepository.Insert(avatar); err != nil {
			return "", EntityInsertionErr
		}
	} else {
		if err = ah.avatarRepository.Update(avatar); err != nil {
			return "", EntityUpdateErr
		}
	}

	user, err := ah.userRepository.GetById(userIdUuid)
	if err != nil {
		return "", EntityNotFoundErr
	}

	user.AvatarHash = avatar.Hash
	if err = ah.userRepository.Update(&user); err != nil {
		return "", EntityInsertionErr
	}

	return avatar.Hash, nil
}

func (ah *AvatarHandler) GetByHash(hash string) ([]byte, string, error) {
	avatar, err := ah.avatarRepository.GetByHash(hash)
	if err != nil {
		return nil, "", EntityNotFoundErr
	}
	imgType, err := converters.GetImageTypeFromBytes(avatar.Content)
	if err != nil {
		return nil, "", err
	}
	return avatar.Content, imgType, nil
}
