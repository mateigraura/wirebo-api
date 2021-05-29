package handlers

import (
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/core"
	"github.com/mateigraura/wirebo-api/crypto"
	"github.com/mateigraura/wirebo-api/models"
)

type RoomHandler struct {
	roomRepository    core.RoomRepository
	messageRepository core.MessageRepository
	hasher            crypto.Hasher
}

func NewRoomHandler(
	roomRepository core.RoomRepository,
	messageRepository core.MessageRepository,
	hasher crypto.Hasher,
) RoomHandler {
	return RoomHandler{
		roomRepository:    roomRepository,
		messageRepository: messageRepository,
		hasher:            hasher,
	}
}

func (rh *RoomHandler) GetPrivateRoomByName(id1, id2 string) (models.Room, error) {
	hash := rh.hasher.Hash([]byte(id1 + id2))
	room, err := rh.roomRepository.GetRoomByName(hex.EncodeToString(hash))
	if err != nil {
		return models.Room{}, EntityNotFoundErr
	}
	err = rh.roomRepository.GetUsersInRoom(&room)
	if err != nil {
		return models.Room{}, EntityNotFoundErr
	}
	return room, nil
}

func (rh *RoomHandler) CreateRoom(request models.CreateRoomRequest) (models.Room, error) {
	roomName := request.Name
	if request.IsPrivate {
		u1Id := request.UsersRefs[0]
		u2Id := request.UsersRefs[1]
		nameHash := rh.hasher.Hash([]byte(u1Id + u2Id))
		roomName = hex.EncodeToString(nameHash)
	}

	room := models.Room{
		Name:      roomName,
		IsPrivate: request.IsPrivate,
	}
	err := rh.roomRepository.Insert(&room)
	if err != nil {
		return models.Room{}, EntityInsertionErr
	}

	err = rh.insertMapping(request.UsersRefs, room.Id)
	return room, err
}

func (rh *RoomHandler) GetRoomsForUser(userId string) ([]models.Room, error) {
	userIdUuid, err := uuid.Parse(userId)
	if err != nil {
		return []models.Room{}, InvalidInputProvidedErr
	}

	rooms, err := rh.roomRepository.GetRoomsFor(userIdUuid)
	if err != nil {
		return []models.Room{}, EntityNotFoundErr
	}

	for idx, room := range rooms {
		err = rh.roomRepository.GetUsersInRoom(&room)
		if err != nil {
			return []models.Room{}, EntityNotFoundErr
		}
		rooms[idx] = room
	}

	return rooms, nil
}

func (rh *RoomHandler) GetRoomMessages(roomId string) ([]models.Message, error) {
	roomIdParsed, err := uuid.Parse(roomId)
	if err != nil {
		return []models.Message{}, err
	}
	messages, err := rh.messageRepository.GetByRoomId(roomIdParsed)
	if err != nil {
		return []models.Message{}, err
	}
	return messages, nil
}

func (rh *RoomHandler) insertMapping(userRefs []string, roomId uuid.UUID) error {
	var roomMapping []*models.UserRoom
	for _, uId := range userRefs {
		userIdParsed, err := uuid.Parse(uId)
		if err != nil {
			return err
		}
		roomMapping = append(roomMapping, &models.UserRoom{
			RoomId: roomId,
			UserId: userIdParsed,
		})
	}
	return rh.roomRepository.InsertMapping(roomMapping)
}
