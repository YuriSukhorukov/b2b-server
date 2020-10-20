package repository

import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
)

type OfferRepository struct {
	db *sqlx.DB
}

func (r OfferRepository) Insert() {
    fmt.Println("OfferRepository: Store()")

    fmt.Println(r.db)

    err := r.db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func (r OfferRepository) Fetch() {
    fmt.Println("OfferRepository: Fetch()")
}