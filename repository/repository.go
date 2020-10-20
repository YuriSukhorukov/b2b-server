package repository

import (
	"github.com/b2b-server/service"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	postgresql *sqlx.DB
	cassandra string
	hbase string
}

func NewRepository() *Repository {
 	psql := service.NewPostgresql()

	return &Repository{psql, "csdr", "hb"}
}