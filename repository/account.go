package repository

import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func (r AccountRepository) Insert() {
    fmt.Println("AccountRepository: Store()")

    fmt.Println(r.db)

    err := r.db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func (r AccountRepository) Fetch() {
    fmt.Println("AccountRepository: Fetch()")
}