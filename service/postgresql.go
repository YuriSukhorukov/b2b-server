package service

import (
	_ "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

func NewPostgresql() *sqlx.DB {
	const (
		hostname 			= "localhost"
		host_port 			= 5432
		username 			= "postgres"
		password 			= "12345"
		database_name 		= "postgres"
		sslmode 			= "disable"
	)

	pg_con_string 	:= fmt.Sprintf("port=%d host=%s user=%s "+"password=%s dbname=%s sslmode=%s", host_port, hostname, username, password, database_name, sslmode)
	db, err 		:= sqlx.Connect("postgres", pg_con_string)

	// TODO обработка ошибок без прекращения работы сервера
	if err != nil {
        log.Fatalln(err)
    }
    return db
}