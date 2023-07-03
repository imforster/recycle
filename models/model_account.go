package models

type Account struct {
	Id int
	Balance int
	Desposits []DepositItem
}
