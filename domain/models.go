package domain

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	Id        uuid.UUID `pg:",pk,type:uuid default gen_random_uuid()" json:"id"`
	CreatedAt time.Time `pg:"default:now()" json:"createdAt"`
	UpdatedAt time.Time `pg:"default:now()" json:"updatedAt"`
}

type User struct {
	BaseModel
	Name         string `pg:",unique" json:"name"`
	Email        string `pg:",notnull,unique" json:"email"`
	PasswordHash string `pg:",notnull,unique" json:"-"`
	Username     string `pg:",notnull,unique" json:"userName"`
	AvatarUrl    string `pg:",notnull" json:"avatarUrl"`
	Rooms        []Room `pg:"many2many:user_rooms" json:"-"`
}

type Message struct {
	BaseModel
	Text     string    `pg:",notnull" json:"text"`
	RoomId   uuid.UUID `pg:",notnull,type:uuid" json:"roomId"`
	Room     Room      `pg:"rel:has-one" json:"room"`
	SenderId uuid.UUID `pg:",notnull,type:uuid" json:"senderId"`
	Sender   User      `pg:"rel:has-one" json:"sender"`
}

type UserRoom struct {
	BaseModel
	UserId uuid.UUID `pg:",notnull,type:uuid"`
	RoomId uuid.UUID `pg:",notnull,type:uuid"`
}

type Room struct {
	BaseModel
	Name     string     `pg:",notnull" json:"name"`
	Users    []User     `pg:"many2many:user_rooms" json:"users"`
	Messages []*Message `pg:"rel:has-many"`
}

type Authorization struct {
	BaseModel
	Token        string    `pg:",notnull"`
	RefreshToken string    `pg:",notnull"`
	OwnerId      uuid.UUID `pg:",notnull,type:uuid"`
}
