package models

import (
	"github.com/imforster/recycle/models"
)

type Customer struct {
	Id       int
	Name     string
	Accounts []Account
}
