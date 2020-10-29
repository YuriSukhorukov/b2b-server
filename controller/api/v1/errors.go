package controller

import "errors"

var (
	ErrEmailOrPasswordNotCorrect = errors.New("email or password is not correct")
	ErrNotAuthorized = errors.New("not authorized")
	ErrEmailNowAvailable = errors.New("not authorized")
	ErrSomethingWrong = errors.New("something went wrong")
)