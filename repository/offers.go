package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/b2b-server/model"
)

type OfferRepository struct {
	db *sqlx.DB
}

func NewOfferRepository(db *sqlx.DB) *OfferRepository {
	return &OfferRepository{db}
}

func (r OfferRepository) InsertOffer(addOffer model.AddOffer) {
	fmt.Println("InsertOffer")
	fmt.Println(addOffer)
}