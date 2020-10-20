package repository

import (
	"github.com/b2b-server/service"
)

type Repository struct {
	Accounts 	AccountRepository
	Offers 		OfferRepository
}

func NewRepository() *Repository {
	psgql 		:= service.NewPostgresql()

	accounts 	:= AccountRepository{psgql}
	offers		:= OfferRepository{psgql}

	return &Repository{accounts, offers}
}