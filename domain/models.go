package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `pg:",pk,type:uuid default gen_random_uuid()" json:"id"`
	Name     string    `pg:",unique" json:"name"`
	Username string    `pg:",notnull,unique" json:"userName"`
	Email    string    `pg:",notnull,unique" json:"email"`
}

// TODO: partial
type Message struct {
	Id     uuid.UUID `pg:",pk,type:uuid default gen_random_uuid()" json:"id"`
	Text   string    `pg:",notnull" json:"text"`
	RoomId uuid.UUID `pg:",notnull" json:"roomId"`
	Sender User      `pg:",notnull" json:"sender"`
}
