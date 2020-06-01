package registry

import (
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/handlers"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/usecase/repository"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/usecase/service"
)

type Registry interface {
	GetAccountHandler() handlers.AccountHandler
}

type registry struct{}

// NewRegistry creates a Registry
func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) GetAccountHandler() handlers.AccountHandler {
	repo := repository.NewAccountRepository()
	service := service.NewAccountService(repo)
	handler := handlers.NewAccountHandler(service)
	return handler
}
