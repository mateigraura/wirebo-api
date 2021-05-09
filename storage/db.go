package storage

import (
	"context"
	utils2 "github.com/mateigraura/wirebo-api/core/utils"
	"log"
	"strconv"
	"sync"

	"github.com/go-pg/pg/v10"
)

var once sync.Once
var connection *pg.DB

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()
	log.Println(string(query))
	return nil
}

func Connection(withLogs ...bool) *pg.DB {
	once.Do(func() {
		env := utils2.GetEnvFile()

		minConn, _ := strconv.Atoi(env[utils2.MinConn])
		maxConn, _ := strconv.Atoi(env[utils2.MaxConn])
		connection = pg.Connect(
			&pg.Options{
				Addr:         env[utils2.DbHost],
				User:         env[utils2.DbUser],
				Password:     env[utils2.DbPsw],
				Database:     env[utils2.DbName],
				MinIdleConns: minConn,
				PoolSize:     maxConn,
			})
	})

	if len(withLogs) > 0 && withLogs[0] {
		connection.AddQueryHook(dbLogger{})
	}

	return connection
}
