package db

import (
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/mateigraura/wirebo-api/utils"
)

var once sync.Once
var connection *pg.DB

func Connection() *pg.DB {
	once.Do(func() {
		env := utils.GetEnvFile()
		connection = pg.Connect(
			&pg.Options{
				Addr:     env[utils.DbHost],
				User:     env[utils.DbUser],
				Password: env[utils.DbPsw],
				Database: env[utils.DbName],
			})
	})

	return connection
}
