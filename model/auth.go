package model

type SignupHeader struct {
	Email   	string    	`header:"Email"`
	Password 	string 		`header:"Password"`
}