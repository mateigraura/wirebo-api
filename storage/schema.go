package storage

import (
	"log"

	"github.com/go-pg/pg/v10/orm"
	"github.com/mateigraura/wirebo-api/domain"
)

func CreateSchema() {
	models := []interface{}{
		(*domain.UserRoom)(nil),
		(*domain.User)(nil),
		(*domain.Message)(nil),
		(*domain.Room)(nil),
		(*domain.Authorization)(nil),
	}

	db := Connection()

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			log.Fatal(err)
		}
	}
}
