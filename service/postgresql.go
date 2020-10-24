package service

import (
	_ "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

func NewPostgresql(host_name string, host_port int, user_name string, password string, database_name string, ssl_mode string) *sqlx.DB {
	pg_con_string 	:= fmt.Sprintf("port=%d host=%s user=%s "+"password=%s dbname=%s sslmode=%s", host_port, host_name, user_name, password, database_name, ssl_mode)
	db, err 		:= sqlx.Connect("postgres", pg_con_string)

	// TODO обработка ошибок без прекращения работы сервера
	if err != nil {
        log.Fatalln(err)
    }
    return db
}