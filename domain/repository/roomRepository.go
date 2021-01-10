package repository

import (
	"github.com/google/uuid"
	"github.com/mateigraura/wirebo-api/domain"
	"github.com/mateigraura/wirebo-api/storage"
)

type RoomRepositoryImpl struct {
}

func (r *RoomRepositoryImpl) GetRoomsFor(userId uuid.UUID) ([]domain.Room, error) {
	conn := storage.Connection()

	rooms := new([]domain.Room)
	err := conn.Model(rooms).
		Column("room.*").
		Join("inner join user_rooms as ur on room.id = ur.room_id").
		Where("ur.user_id = ?", userId).
		Select()

	if err != nil {
		return []domain.Room{}, err
	}

	return *rooms, nil
}

func (r *RoomRepositoryImpl) GetUsersInRoom(room domain.Room) (domain.Room, error) {
	conn := storage.Connection()

	err := conn.Model(&room.Users).
		Column("user.*").
		Join(`inner join user_rooms as ur on "user"."id" = ur.user_id`).
		Where("ur.room_id = ?", room.Id).
		Select()

	if err != nil {
		return domain.Room{}, err
	}

	return room, nil
}

func (r *RoomRepositoryImpl) Insert(room *domain.Room) error {
	conn := storage.Connection()

	_, err := conn.Model(room).
		Returning("id").
		Insert()

	return err
}

func (r *RoomRepositoryImpl) InsertMapping(values []interface{}) error {
	conn := storage.Connection()

	for _, v := range values {
		_, err := conn.Model(v).Insert()
		if err != nil {
			return err
		}
	}

	return nil
}
