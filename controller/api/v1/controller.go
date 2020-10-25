package controller

import (
	"github.com/b2b-server/service"
	"github.com/b2b-server/repository"
)

// Controller example
type Controller struct {
	JWT 				service.JWT
	UserRepository 		repository.UserRepository
	OfferRepository 	repository.OfferRepository
}

// NewController example
func NewController(jwt service.JWT, ur repository.UserRepository, or repository.OfferRepository) *Controller {
	return &Controller{jwt, ur, or}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}
