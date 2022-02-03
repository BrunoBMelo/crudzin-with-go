package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type postgreSql struct {
	provider   string
	dataSource string
}

func (p *postgreSql) GetConnection() *sql.DB {
	if db, err := sql.Open(p.provider, p.dataSource); err == nil {
		return db
	} else {
		panic(err)
	}
}

func NewPostgreSql(provider string, dataSource string) *postgreSql {
	return &postgreSql{provider: provider, dataSource: dataSource}
}
