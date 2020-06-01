package service

import (
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/constants"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/domain"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/errors"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/usecase/repository"
)

// AccountService defines an interface for Account business logic
type AccountService interface {
	GetAccountBalance() (*domain.AccountBalance, *errors.APIError)
	GetTransactionsList() ([]*domain.Transaction, *errors.APIError)
	GetTransaction(string) (*domain.Transaction, *errors.APIError)
	SaveTransaction(*domain.TransactionOp) (*domain.Transaction, *errors.APIError)
}

type accountService struct {
	repo repository.AccountRepository
}

// NewAccountService creates an AccountService
func NewAccountService(r repository.AccountRepository) AccountService {
	return &accountService{r}
}

func (s *accountService) GetAccountBalance() (*domain.AccountBalance, *errors.APIError) {
	accountBalance := s.repo.GetAccountBalance()
	return accountBalance, nil
}

func (s *accountService) GetTransactionsList() ([]*domain.Transaction, *errors.APIError) {
	transactions := s.repo.GetTransactionsList()
	return transactions, nil
}

func (s *accountService) GetTransaction(id string) (*domain.Transaction, *errors.APIError) {
	transaction := s.repo.GetTransactionByID(id)
	if transaction == nil {
		return nil, errors.NotFound()
	}
	return transaction, nil
}

func (s *accountService) SaveTransaction(txData *domain.TransactionOp) (*domain.Transaction, *errors.APIError) {
	if txData.Type != constants.DEBIT.String() && txData.Type != constants.CREDIT.String() {
		return nil, errors.BadRequest("Invalid transaction type")
	}
	if txData.Amount < 0 {
		return nil, errors.BadRequest("Invalid amount")
	}
	transaction, err := s.repo.SaveTransaction(txData)
	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}
	return transaction, nil
}
