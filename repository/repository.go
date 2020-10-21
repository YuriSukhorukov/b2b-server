package repository

type Repository struct {
	Accounts 	AccountRepository
	Offers 		OfferRepository
}

func NewRepository(accounts AccountRepository, offers OfferRepository) *Repository {
	return &Repository{accounts, offers}
}