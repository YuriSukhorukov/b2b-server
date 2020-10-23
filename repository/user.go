package repository

import (
	"fmt"
	"errors"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/b2b-server/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) IsEmailFree(email string) (error, bool) {
	a := model.User{}
	s := `
		SELECT user_id FROM users
        WHERE email=$1
        ORDER BY user_id ASC;
	`
	err := r.db.Get(&a, s, email)
	switch err {
		case nil:
		    return nil, false
		case sql.ErrNoRows:
		    return nil, true
	}

	return err, false
}

func (r UserRepository) InsertUser(email string, password string) (error, model.User) {
    m := model.User{}
    s := `
        INSERT INTO users(email, password) 
        VALUES ($1, $2)
        RETURNING user_id;
    `
    rows, err := r.db.Queryx(s, email, password)

	if err != nil {
		switch err.Error() {
			case "pq: duplicate key value violates unique constraint \"users_email_key\"":
			    return errors.New("duplicate email"), model.User{}
		    default:
		    	fmt.Println(err.Error())
			    return errors.New("something wrong"), model.User{}
		}
	}

	for rows.Next() {
   		err = rows.StructScan(&m)
	}

    return err, m
}