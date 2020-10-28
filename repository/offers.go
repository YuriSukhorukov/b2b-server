package repository

import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
	"github.com/b2b-server/model"
)

type OfferRepository struct {
	db *sqlx.DB
}

func NewOfferRepository(db *sqlx.DB) *OfferRepository {
	return &OfferRepository{db}
}

func (r OfferRepository) InsertOffer() {
    fmt.Println("OfferRepository: Store()")

    fmt.Println(r.db)

    err := r.db.Ping()
	if err != nil {
		log.Fatalln(err)
	}



	m := model.Offer{}
	_ = m
}

func (r OfferRepository) Fetch() {
    fmt.Println("OfferRepository: Fetch()")
}