package repository

import (
	"fmt"
)

func (r Repository) StoreAccount() {
    fmt.Println("Repository: Store()")
    fmt.Println(r.postgresql)
    fmt.Println(r.cassandra)
    fmt.Println(r.hbase)
}

func (r Repository) FetchAccount() {
    fmt.Println("Repository: Fetch()")
}