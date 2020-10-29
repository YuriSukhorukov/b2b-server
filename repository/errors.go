package repository

import "errors"

var (
	ErrDublicateEmail = errors.New("duplicate email")
	ErrUserNotFound = errors.New("user not found")
	ErrSomethingWrong = errors.New("something went wrong")
)