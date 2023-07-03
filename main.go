package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// CustomersCustomerIdAccountsAccountIdBalancesGet - Reterive current account balance
	e.GET("/customers/:customerId/accounts/:accountId/balances", c.CustomersCustomerIdAccountsAccountIdBalancesGet)

	// CustomersCustomerIdAccountsAccountIdDepositsGet - Retreive all deposits for a customer
	e.GET("/customers/:customerId/accounts/:accountId/deposits", c.CustomersCustomerIdAccountsAccountIdDepositsGet)

	// CustomersCustomerIdAccountsAccountIdDepositsPost - Record deposit of a recycled item
	e.POST("/customers/:customerId/accounts/{accountId]/deposits", c.CustomersCustomerIdAccountsAccountIdDepositsPost)

	// CustomersCustomerIdAccountsAccountIdTransfersPost - Transfer account balance to Paypal or other service
	e.POST("/customers/:customerId/accounts/:accountId/transfers", c.CustomersCustomerIdAccountsAccountIdTransfersPost)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}