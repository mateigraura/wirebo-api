package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `pg:",pk,type:uuid default gen_random_uuid()" json:"id"`
	Name     string    `pg:",unique" json:"name"`
	Username string    `pg:",notnull,unique" json:"userName"`
	Email    string    `pg:",notnull,unique" json:"email"`
}

type Message struct {
}
