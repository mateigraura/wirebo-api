package domain

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `pg:",pk,type:uuid default gen_random_uuid()" json:"id"`
	Name         string    `pg:",unique" json:"name"`
	Email        string    `pg:",notnull,unique" json:"email"`
	PasswordHash string    `pg:",notnull,unique" json:"-"`
	Username     string    `pg:",notnull,unique" json:"userName"`
	AvatarUrl    string    `pg:",notnull" json:"avatarUrl"`
	Rooms        []Room    `pg:"many2many:user_rooms" json:"rooms"`
}

type Message struct {
	Id       uuid.UUID `pg:",pk,type:uuid default gen_random_uuid()" json:"id"`
	Text     string    `pg:",notnull" json:"text"`
	RoomId   uuid.UUID `pg:",notnull" json:"roomId"`
	Room     Room      `pg:"rel:has-one" json:"room"`
	SenderId uuid.UUID `pg:",notnull" json:"senderId"`
	Sender   User      `pg:"rel:has-one" json:"sender"`
}

type UserRoom struct {
	Id     uuid.UUID `pg:",pk,type:uuid default gen_random_uuid()" json:"id"`
	UserId uuid.UUID `pg:",notnull"`
	RoomId uuid.UUID `pg:",notnull"`
}

type Room struct {
	Id       uuid.UUID  `pg:",pk,type:uuid default gen_random_uuid()" json:"id"`
	Name     string     `pg:",notnull" json:"name"`
	Users    []User     `pg:"many2many:user_rooms" json:"users"`
	Messages []*Message `pg:"rel:has-many"`
}
