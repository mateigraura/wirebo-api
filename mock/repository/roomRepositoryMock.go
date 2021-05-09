package repository

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/models"
)

func getUserRooms() []models.UserRoom {
	user1, _ := uuid.Parse("")
	user2, _ := uuid.Parse("")
	user3, _ := uuid.Parse("")
	room1, _ := uuid.Parse("")
	room2, _ := uuid.Parse("")

	return []models.UserRoom{
		{
			UserId: user1,
			RoomId: room1,
		},
		{
			UserId: user2,
			RoomId: room1,
		},
		{
			UserId: user3,
			RoomId: room2,
		},
		{
			UserId: user1,
			RoomId: room2,
		},
	}
}

func getRooms() []models.Room {
	room1, _ := uuid.Parse("")
	room2, _ := uuid.Parse("")
	return []models.Room{
		{
			BaseModel: models.BaseModel{
				Id: room1,
			},
			Name: "room1",
		},
		{
			BaseModel: models.BaseModel{
				Id: room2,
			},
			Name: "room2",
		},
	}
}

type RoomRepositoryMock struct {
}

func (r *RoomRepositoryMock) GetRoomsFor(userId uuid.UUID) ([]models.Room, error) {
	var rooms []models.Room
	mapping := getUserRooms()
	for _, userRoom := range mapping {
		if userRoom.UserId == userId {
			for _, room := range getRooms() {
				if room.Id == userRoom.RoomId {
					rooms = append(rooms, room)
				}
			}
		}
	}
	return rooms, nil
}

func (r *RoomRepositoryMock) GetUsersInRoom(room models.Room) (models.Room, error) {
	return room, nil
}

func (r *RoomRepositoryMock) Insert(room *models.Room) error {
	return nil
}

func (r *RoomRepositoryMock) InsertMapping(values []interface{}) error {
	return nil
}
