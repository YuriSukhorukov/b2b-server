package repository

import (
	"fmt"
	"log"
)

func (r Repository) InsertAccount() {
    fmt.Println("Repository: Store()")

    fmt.Println(r.postgresql)
    fmt.Println(r.cassandra)
    fmt.Println(r.hbase)

    err := r.postgresql.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func (r Repository) FetchAccount() {
    fmt.Println("Repository: Fetch()")
}