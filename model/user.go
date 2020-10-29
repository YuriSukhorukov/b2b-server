package model

import "errors"

type AddUser struct {
	Email   	string  `json:"email" example:"my_email@gmail.com" format:"string" header:"email"`
	Password 	string 	`json:"password" example:"pa!ssvv0rD" format:"string" header:"password"`
}

type AuthUser struct {
	Email   	string 	`db:"email" json:"email" example:"my_email@gmail.com" format:"string" header:"email"`
	Password 	string 	`db:"password" json:"password" example:"pa!ssvv0rD" format:"string" header:"password"`
}

var (
	ErrEmailInvalid = errors.New("email is empty")
	ErrPasswordInvalid = errors.New("password is empty")
)

func (a AddUser) Validation() error {
	switch {
		case len(a.Email) == 0:
			return ErrEmailInvalid
		case len(a.Password) == 0:
			return ErrPasswordInvalid
		default:
			return nil
	}
}

func (a AuthUser) Validation() error {
	switch {
		case len(a.Email) == 0:
			return ErrEmailInvalid
		case len(a.Password) == 0:
			return ErrPasswordInvalid
		default:
			return nil
	}
}

type User struct {
    UserID       string `db:"user_id"`
    Email		 string `db:"email"`
    Password 	 string `db:"password"`
    CreatedOn	 string `db:"created_on"`
}