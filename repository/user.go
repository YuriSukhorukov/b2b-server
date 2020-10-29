package repository

import (
	"fmt"
	"errors"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/b2b-server/model"
	// jwt "github.com/dgrijalva/jwt-go"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) IsEmailFree(email string) (error, bool) {
	m := model.User{}
	s := `
		SELECT user_id FROM users
        WHERE email=$1
        ORDER BY user_id ASC;
	`
	err := r.db.Get(&m, s, email)
	switch err {
		case nil:
		    return nil, false
		case sql.ErrNoRows:
		    return nil, true
	}

	return err, false
}

func (r UserRepository) InsertUser(user model.User) (error, *model.Record) {
    u := model.User{}
    s := `
        INSERT INTO users(email, password) 
        VALUES (:email, :password)
        RETURNING user_id, email, created_on;
    `
    rows, err := r.db.NamedQuery(s, user)

	if err != nil {
		switch err.Error() {
			case "pq: duplicate key value violates unique constraint \"users_email_key\"":
			    return errors.New("duplicate email"), nil		    
			default:
		    	fmt.Println(err.Error())
			    return errors.New("something wrong"), nil
		}
	}

	for rows.Next() {
   		err = rows.StructScan(&u)
	}

    return err, &model.Record{u.UserID, u.CreatedOn}
}

func (r UserRepository) AuthorizeUser(user model.User) (error, *model.Record) {
	u := model.User{}
    s := `
        SELECT user_id, created_on 
        FROM users
        WHERE email = :email
        AND password = :password
        ORDER BY user_id ASC;
    `
    rows, err := r.db.NamedQuery(s, user)

    if err != nil {
    	return err, nil
    }

	for rows.Next() {
   		err = rows.StructScan(&u)
	}

    return err, &model.Record{u.UserID, u.CreatedOn}
}