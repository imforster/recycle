package handlers

import (
    "github.com/imforster/recycle/models"
    "fmt"
)

// Container will hold all dependencies for your application.
type Container struct {
    Customers []models.Customer
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer() (Container, error) {
    a := []models.Account{Id:0}
    c := Container{Customers: []models.Customer{models.Customer{Id:0,Name:"Ian",Accounts:a}}}
    fmt.Println(c.Customers)
    return c, nil
}