package model

type Check struct {
	Success bool `json:"success" example:"true"`
	Message string `json:"message" example:"string"`
}