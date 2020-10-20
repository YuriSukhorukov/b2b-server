package services

import "fmt"

type IPostgresql interface {
	StoreAccount() 	()
	FetchAccount() 	()
}

type PostgresqlService struct {}

func (p PostgresqlService) StoreAccount() {
	fmt.Println("PostgresqlService: StoreAccount()")
}
func (p PostgresqlService) FetchAccount() {
	fmt.Println("PostgresqlService: FetchAccount()")
}