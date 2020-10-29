package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/b2b-server/model"
)

type OfferRepository struct {
	db *sqlx.DB
}

func NewOfferRepository(db *sqlx.DB) *OfferRepository {
	return &OfferRepository{db}
}

func (r OfferRepository) InsertOffer(offer model.Offer) (error, *model.Created) {
	o := model.Offer{}
	s := `
        INSERT INTO offers(user_id, title, description, price, amount, currency_code, offer_type, measure_unit_code, date_expires, country, city) 
        VALUES (:user_id, :title, :description, :price, :amount, :currency_code, :offer_type, :measure_unit_code, :date_expires, :country, :city)
        RETURNING offer_id, created_on;
    `
    rows, err := r.db.NamedQuery(s, offer)

    if err != nil {
    	return ErrSomethingWrong, nil
    }
	for rows.Next() {
   		err = rows.StructScan(&o)
	}
	if err != nil {
    	return ErrSomethingWrong, nil
    }

	return nil, &model.Created{o.OfferID, o.CreatedOn}
}