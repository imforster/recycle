package handlers

import "models"

// Container will hold all dependencies for your application.
type Container struct {
    Customers []Customer
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer() (Container, error) {
    c := Container{}
    return c, nil
}