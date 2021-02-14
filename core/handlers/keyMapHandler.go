package handlers

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/core"
	"github.com/mateigraura/wirebo-api/models"
)

type KeyMapHandler struct {
	keyMapRepository core.KeyMapperRepository
}

func NewKeyMapHandler(keyMapRepository core.KeyMapperRepository) KeyMapHandler {
	return KeyMapHandler{
		keyMapRepository: keyMapRepository,
	}
}

func (kmh *KeyMapHandler) GetKey(id string) (string, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return kmh.findKey(parsedId)
}

func (kmh *KeyMapHandler) InsertKey(id, pubKey string) (string, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	key, err := kmh.findKey(parsedId)
	if err != nil {
		keyMapping := &models.KeyMapping{
			OwnerId: parsedId,
			PubKey:  pubKey,
		}
		err = kmh.keyMapRepository.Insert(keyMapping)
		if err != nil {
			return "", err
		}

		return keyMapping.PubKey, nil
	}

	return key, nil
}

func (kmh *KeyMapHandler) findKey(id uuid.UUID) (string, error) {
	keyMapping, err := kmh.keyMapRepository.Get(id)
	if err != nil {
		return "", err
	}

	return keyMapping.PubKey, nil
}
