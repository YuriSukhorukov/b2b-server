package model

type Error struct {
	Success bool `json:"success" example:"false"`
	Error string `json:"error" example:"something wrong"`
}