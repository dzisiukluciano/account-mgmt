package main

import (
	"fmt"
	"net/http"

	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/registry"

	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/constants"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/utils"
	"github.com/labstack/echo/v4"
)

var (
	host    = utils.GetEnv("HOST", constants.DefaultHost)
	port    = utils.GetEnv("PORT", constants.DefaultPort)
	apiPath = utils.GetEnv("BASE_PATH", constants.BasePath)
)

func main() {
	e := echo.New()
	e.GET(constants.HCPath, func(c echo.Context) error {
		return c.String(http.StatusOK, "up")
	})

	container := registry.NewRegistry()
	accHdlr := container.GetAccountHandler()

	base := e.Group(apiPath)

	accRoutes := base.Group(constants.AccountPath)
	accRoutes.GET("/balance", accHdlr.GetAccountBalance)

	txRoutes := accRoutes.Group(constants.TransactionPath)
	txRoutes.GET("/:id", accHdlr.GetTransaction)
	txRoutes.POST("", accHdlr.SaveTransaction)

	txsRoutes := accRoutes.Group(constants.TransactionsPath)
	txsRoutes.GET("", accHdlr.GetTransactions)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", host, port)))
}
