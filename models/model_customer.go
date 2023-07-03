package models

type Customer struct {
	Id   int
	Name string
    Accounts []models.Account
}
