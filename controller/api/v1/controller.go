package controller

import "github.com/b2b-server/repository"

// Controller example
type Controller struct {
	Repository repository.Repository
}

// NewController example
func NewController(r repository.Repository) *Controller {
	return &Controller{r}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}
