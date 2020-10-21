package controller

import "github.com/b2b-server/repository"

// Controller example
type Controller struct {
	UserRepository 		repository.UserRepository
	OfferRepository 	repository.OfferRepository
}

// NewController example
func NewController(ur repository.UserRepository, or repository.OfferRepository) *Controller {
	return &Controller{ur, or}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}
