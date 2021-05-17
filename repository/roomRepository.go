package repository

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/models"
	"github.com/mateigraura/wirebo-api/storage"
)

type RoomRepositoryImpl struct {
}

func (r *RoomRepositoryImpl) GetRoomByName(name string) (models.Room, error) {
	conn := storage.Connection()
	room := new(models.Room)
	err := conn.Model(room).Where("name = ?", name).Select()
	if err != nil {
		return models.Room{}, err
	}
	return *room, nil
}

func (r *RoomRepositoryImpl) GetRoomsFor(userId uuid.UUID) ([]models.Room, error) {
	conn := storage.Connection()

	rooms := new([]models.Room)
	err := conn.Model(rooms).
		Column("room.*").
		Join("inner join user_rooms as ur on room.id = ur.room_id").
		Where("ur.user_id = ?", userId).
		Select()

	if err != nil {
		return []models.Room{}, err
	}

	return *rooms, nil
}

func (r *RoomRepositoryImpl) GetUsersInRoom(room *models.Room) error {
	conn := storage.Connection()
	return conn.Model(&room.Users).
		Column("user.*").
		Join(`inner join user_rooms as ur on "user"."id" = ur.user_id`).
		Where("ur.room_id = ?", room.Id).
		Select()
}

func (r *RoomRepositoryImpl) Insert(room *models.Room) error {
	conn := storage.Connection()

	_, err := conn.Model(room).
		Returning("id").
		Insert()

	return err
}

func (r *RoomRepositoryImpl) InsertMapping(mapping []*models.UserRoom) error {
	conn := storage.Connection()

	for _, v := range mapping {
		_, err := conn.Model(v).Insert()
		if err != nil {
			return err
		}
	}

	return nil
}
