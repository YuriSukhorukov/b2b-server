package controller

import "github.com/b2b-server/repository"

// Controller example
type Controller struct {
	AccountRepository 		repository.AccountRepository
	OfferRepository 		repository.OfferRepository
}

// NewController example
func NewController(ar repository.AccountRepository, or repository.OfferRepository) *Controller {
	return &Controller{ar, or}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}
