package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/constants"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/domain"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/utils"
)

// AccountRepository defines an interface for account operations
type AccountRepository interface {
	GetAccountBalance() *domain.AccountBalance
	GetTransactionsList() []*domain.Transaction
	GetTransactionByID(string) *domain.Transaction
	SaveTransaction(*domain.TransactionOp) (*domain.Transaction, error)
}

type accountRepository struct {
	accountBalance    *domain.AccountBalance
	transactions      []*domain.Transaction
	transactionsIndex map[string]*domain.Transaction
	lock              sync.RWMutex
}

// NewAccountRepository creates a new AccountRepository
func NewAccountRepository() AccountRepository {
	return &accountRepository{
		accountBalance:    &domain.AccountBalance{},
		transactions:      make([]*domain.Transaction, 0),
		transactionsIndex: make(map[string]*domain.Transaction),
	}
}

func (r *accountRepository) GetAccountBalance() *domain.AccountBalance {
	return r.accountBalance
}

func (r *accountRepository) GetTransactionsList() []*domain.Transaction {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.transactions
}

func (r *accountRepository) GetTransactionByID(id string) *domain.Transaction {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.transactionsIndex[id]
}

// Inserts a transaction and emulates database constraints
func (r *accountRepository) SaveTransaction(txData *domain.TransactionOp) (*domain.Transaction, error) {
	if txData == nil {
		return nil, errors.New("TransactionOp cannot be nil")
	}
	if txData.Amount < 0 {
		return nil, errors.New("Amount cannot be negative")
	}
	if r.transactionsIndex == nil {
		return nil, errors.New("TransactionRepository not initialized")
	}
	r.lock.Lock()
	defer r.lock.Unlock()

	balance := r.accountBalance.Balance
	if txData.Type == constants.DEBIT.String() && txData.Amount > balance {
		return nil, errors.New("Non-sufficient funds")
	}
	newTransaction := &domain.Transaction{
		ID:            utils.GetUUID(),
		Type:          txData.Type,
		Amount:        txData.Amount,
		EffectiveDate: time.Now().Format(time.RFC3339),
	}
	r.transactions = append(r.transactions, newTransaction)
	r.transactionsIndex[newTransaction.ID] = newTransaction
	r.accountBalance.Balance = balance + newTransaction.Delta()

	return newTransaction, nil
}
