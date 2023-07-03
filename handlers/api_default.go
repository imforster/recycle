package handlers

import (
    "github.com/imforster/recycle/models"
    "github.com/labstack/echo/v4"
    "net/http"
)

// CustomersCustomerIdAccountsAccountIdBalancesGet - Reterive current account balance
func (c *Container) CustomersCustomerIdAccountsAccountIdBalancesGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// CustomersCustomerIdAccountsAccountIdDepositsGet - Retreive all deposits for a customer
func (c *Container) CustomersCustomerIdAccountsAccountIdDepositsGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// CustomersCustomerIdAccountsAccountIdDepositsPost - Record deposit of a recycled item
func (c *Container) CustomersCustomerIdAccountsAccountIdDepositsPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// CustomersCustomerIdAccountsAccountIdTransfersPost - Transfer account balance to Paypal or other service
func (c *Container) CustomersCustomerIdAccountsAccountIdTransfersPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
