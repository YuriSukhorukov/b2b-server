package model

type Account struct {
    UserID       string `db:"user_id"`
    Email        string `db:"email"`
    Password 	 string `db:"password"`
}