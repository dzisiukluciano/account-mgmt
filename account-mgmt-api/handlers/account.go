package handlers

import (
	"net/http"

	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/usecase/service"

	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/domain"
	"github.com/labstack/echo/v4"
)

type AccountHandler interface {
	GetAccountBalance(c echo.Context) error
	GetTransactions(c echo.Context) error
	GetTransaction(c echo.Context) error
	SaveTransaction(c echo.Context) error
}

type accountHandler struct {
	service service.AccountService
}

// NewAccountHandler creates a new AccountHandler
func NewAccountHandler(service service.AccountService) AccountHandler {
	return &accountHandler{service}
}

func (h *accountHandler) GetAccountBalance(c echo.Context) error {
	balance, apiError := h.service.GetAccountBalance()
	if apiError != nil {
		return c.JSON(apiError.Status, apiError)
	}
	return c.JSON(http.StatusOK, balance)
}

func (h *accountHandler) GetTransactions(c echo.Context) error {
	transactions, apiError := h.service.GetTransactionsList()
	if apiError != nil {
		return c.JSON(apiError.Status, apiError)
	}
	return c.JSON(http.StatusOK, transactions)
}

func (h *accountHandler) GetTransaction(c echo.Context) error {
	id := c.Param("id")
	transaction, apiError := h.service.GetTransaction(id)
	if apiError != nil {
		return c.JSON(apiError.Status, apiError)
	}
	return c.JSON(http.StatusOK, transaction)
}

func (h *accountHandler) SaveTransaction(c echo.Context) error {
	body := &domain.TransactionOp{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	transaction, apiError := h.service.SaveTransaction(body)
	if apiError != nil {
		return c.JSON(apiError.Status, apiError)
	}
	return c.JSON(http.StatusOK, transaction)
}
