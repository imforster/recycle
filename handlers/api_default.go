package handlers

import (
	"net/http"
	"github.com/imforster/recycle/models"
	"github.com/labstack/echo/v4"
	"strconv"
	log "github.com/sirupsen/logrus"
)

// CustomersCustomerIdAccountsAccountIdBalancesGet - Reterive current account balance
func (c *Container) CustomersCustomerIdAccountsAccountIdBalancesGet(ctx echo.Context) error {
	log.Info("Retrieving balances")
	customerId, err := strconv.Atoi(ctx.Param("customerId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	accountId, err2 := strconv.Atoi(ctx.Param("accountId"))
	if err2 != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	bal := c.Customers[customerId].Accounts[accountId].Balance
	return ctx.JSON(http.StatusOK, models.ResponseMessage{
		Message: "Balance is: " + strconv.Itoa(bal),
	})
}

// CustomersCustomerIdAccountsAccountIdDepositsGet - Retreive all deposits for a customer
func (c *Container) CustomersCustomerIdAccountsAccountIdDepositsGet(ctx echo.Context) error {
	customerId, err := strconv.Atoi(ctx.Param("customerId"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	accountId, err2 := strconv.Atoi(ctx.Param("accountId"))
	if err2 != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request!")
	}
	return ctx.JSON(http.StatusOK, c.Customers[customerId].Accounts[accountId].Deposits)
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
	log.Info(d)
	if (d.ItemSize == "small") {
		c.Customers[customerId].Accounts[accountId].Balance = c.Customers[customerId].Accounts[accountId].Balance + 5
	}
	if (d.ItemSize == "large") {
		c.Customers[customerId].Accounts[accountId].Balance = c.Customers[customerId].Accounts[accountId].Balance + 10
	}
	c.Customers[customerId].Accounts[accountId].Deposits =
		append(c.Customers[customerId].Accounts[accountId].Deposits, *d)
	bal := c.Customers[customerId].Accounts[accountId].Balance
	log.Info(c)
	return ctx.JSON(http.StatusOK, models.ResponseMessage{
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

	log.Info("Transfer Balance: " + strconv.Itoa(c.Customers[customerId].Accounts[accountId].Balance) + " to PayPal")
	c.Customers[customerId].Accounts[accountId].Balance = 0
	return ctx.JSON(http.StatusOK, models.ResponseMessage{
		Message: "Balance is: " + strconv.Itoa(c.Customers[customerId].Accounts[accountId].Balance),
	})
}
