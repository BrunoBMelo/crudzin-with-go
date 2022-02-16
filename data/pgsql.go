package data

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type postgreSql struct {
	provider   string
	dataSource string
}

func (p *postgreSql) GetConnection() *sql.DB {
	if db, err := sql.Open(p.provider, p.dataSource); err == nil {
		db.SetMaxIdleConns(100)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(1 * time.Nanosecond)
		return db
	} else {
		panic(err)
	}
}

func NewPostgreSql(provider string, dataSource string) *postgreSql {
	return &postgreSql{provider: provider, dataSource: dataSource}
}
