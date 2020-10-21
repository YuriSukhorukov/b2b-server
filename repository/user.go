package repository

import (
	"fmt"
	"log"
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
		SELECT email FROM users
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

func (r UserRepository) Insert() {
    fmt.Println("UserRepository: Store()")

    fmt.Println(r.db)

    err := r.db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	// result, err := r.db.Exec(`CREATE TABLE place (
	//     country text,
	//     city text NULL,
	//     telcode integer);
	// `)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)

	a := model.User{}
	err = r.db.Get(&a, `
		SELECT email FROM users
        WHERE email='user_xxx@gmail.com'
        ORDER BY user_id ASC;
	`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}