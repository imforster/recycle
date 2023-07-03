package handlers

import (
	"fmt"

	"github.com/imforster/recycle/models"
)

// Container will hold all dependencies for your application.
type Container struct {
	Customers map[int]models.Customer
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer() (Container, error) {
	c := Container{Customers: map[int]models.Customer{
        0:models.Customer{Id: 0, Name: "Ian", Accounts: []models.Account{models.Account{Id: 0}}},
        1:models.Customer{Id: 1, Name: "Tad", Accounts: []models.Account{models.Account{Id: 0}}},
    }}
	fmt.Println(c.Customers)
	return c, nil
}