package repository

import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	DB *sqlx.DB
}

func (r AccountRepository) Insert() {
    fmt.Println("AccountRepository: Store()")

    fmt.Println(r.DB)

    err := r.DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func (r AccountRepository) Fetch() {
    fmt.Println("AccountRepository: Fetch()")
}