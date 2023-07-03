package handlers

import (
	"net/http"
	"fmt"
	"github.com/imforster/recycle/models"
	"github.com/labstack/echo/v4"
	"strconv"
)

// CustomersCustomerIdAccountsAccountIdBalancesGet - Reterive current account balance
func (c *Container) CustomersCustomerIdAccountsAccountIdBalancesGet(ctx echo.Context) error {
	fmt.Println("Retrieving balances")
	customerId, err := strconv.Atoi(ctx.Param("customerId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	accountId, err2 := strconv.Atoi(ctx.Param("accountId"))
	if err2 != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	bal := c.Customers[customerId].Accounts[accountId].Balance
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Balance is: " + strconv.Itoa(bal),
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
	customerId, err := strconv.Atoi(ctx.Param("customerId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	accountId, err2 := strconv.Atoi(ctx.Param("accountId"))
	if err2 != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	d := new(models.DepositItem)
	if err := ctx.Bind(d); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	fmt.Println(d)
	if (d.ItemSize == "small") {
		c.Customers[customerId].Accounts[accountId].Balance = c.Customers[customerId].Accounts[accountId].Balance + 5
	}
	if (d.ItemSize == "large") {
		c.Customers[customerId].Accounts[accountId].Balance = c.Customers[customerId].Accounts[accountId].Balance + 10
	}
	c.Customers[customerId].Accounts[accountId].Deposits.append(d)
	bal := c.Customers[customerId].Accounts[accountId].Balance
	fmt.Println(c)
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Balance is: " + strconv.Itoa(bal),
	})
}

// CustomersCustomerIdAccountsAccountIdTransfersPost - Transfer account balance to Paypal or other service
func (c *Container) CustomersCustomerIdAccountsAccountIdTransfersPost(ctx echo.Context) error {
	customerId, err := strconv.Atoi(ctx.Param("customerId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	accountId, err2 := strconv.Atoi(ctx.Param("accountId"))
	if err2 != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	c.Customers[customerId].Accounts[accountId].Balance = 10
	fmt.Println("Balance is: " + strconv.Itoa(c.Customers[customerId].Accounts[accountId].Balance))
	c.Customers[customerId].Accounts[accountId].Balance = 0
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Balance is: " + strconv.Itoa(c.Customers[customerId].Accounts[accountId].Balance),
	})
}
