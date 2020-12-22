package storage

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/mateigraura/wirebo-api/domain"
)

func CreateSchema(db *pg.DB) {
	models := []interface{}{
		(*domain.User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			log.Fatal(err)
		}
	}
}