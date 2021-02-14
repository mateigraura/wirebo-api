package storage

import (
	"log"

	"github.com/go-pg/pg/v10/orm"
	"github.com/mateigraura/wirebo-api/models"
)

func CreateSchema() {
	entities := []interface{}{
		(*models.UserRoom)(nil),
		(*models.User)(nil),
		(*models.Message)(nil),
		(*models.Room)(nil),
		(*models.Authorization)(nil),
		(*models.KeyMapping)(nil),
	}

	db := Connection()

	for _, model := range entities {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			log.Fatal(err)
		}
	}
}
