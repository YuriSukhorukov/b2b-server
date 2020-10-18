package model

type SignupHeader struct {
	Email   	string    	`header:"email"`
	Password 	string 		`header:"password"`
}