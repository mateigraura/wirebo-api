package storage

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/mateigraura/wirebo-api/utils"
)

var once sync.Once
var connection *pg.DB

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent)  error {
	log, _ := q.FormattedQuery()
	fmt.Println(string(log))
	return nil
}

func Connection() *pg.DB {
	once.Do(func() {
		env := utils.GetEnvFile()

		minConn, _ := strconv.Atoi(env[utils.MinConn])
		maxConn, _ := strconv.Atoi(env[utils.MaxConn])
		connection = pg.Connect(
			&pg.Options{
				Addr:         env[utils.DbHost],
				User:         env[utils.DbUser],
				Password:     env[utils.DbPsw],
				Database:     env[utils.DbName],
				MinIdleConns: minConn,
				PoolSize:     maxConn,
			})
	})
	connection.AddQueryHook(dbLogger{})

	return connection
}
