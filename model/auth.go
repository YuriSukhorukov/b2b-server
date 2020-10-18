package model

type AuthHeader struct {
	Email   	string    	`header:"email"`
	Password 	string 		`header:"password"`
}