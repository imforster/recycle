package models

type Account struct {
	Id int
	Balance int
	Deposits []DepositItem
}
