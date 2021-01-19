package storage

import (
	"context"
	"log"
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
	query, _ := q.FormattedQuery()
	log.Println(string(query))
	return nil
}

func Connection(withLogs... bool) *pg.DB {
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

	if len(withLogs) > 0 && withLogs[0] {
		connection.AddQueryHook(dbLogger{})
	}

	return connection
}
