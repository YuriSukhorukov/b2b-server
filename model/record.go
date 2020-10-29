package model

import "errors"

type Record struct {
	ID 			string `json:"id" example:"1d586b05-7b80-4a3a-bf2c-ce48169d4e85"`
	CreatedOn 	string `json:"created_on" example:"2020-10-23 18:02:35.745565"`
}

var (
	ErrIDInvalid = errors.New("id is empty")
	ErrCreatedOnInvalid = errors.New("created_on is empty")
)

func (r Record) Validation() error {
	switch {
		case len(r.ID) == 0:
			return ErrIDInvalid
		case len(r.CreatedOn) == 0:
			return ErrCreatedOnInvalid
		default:
			return nil
	}
}