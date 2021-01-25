package storage

import (
	models2 "github.com/mateigraura/wirebo-api/models"
	"log"

	"github.com/go-pg/pg/v10/orm"
)

func CreateSchema() {
	models := []interface{}{
		(*models2.UserRoom)(nil),
		(*models2.User)(nil),
		(*models2.Message)(nil),
		(*models2.Room)(nil),
		(*models2.Authorization)(nil),
	}

	db := Connection()

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			log.Fatal(err)
		}
	}
}
