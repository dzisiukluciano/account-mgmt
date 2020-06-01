package repository

import (
	"strings"
	"testing"

	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/constants"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/domain"
)

func TestSaveTransaction(t *testing.T) {
	mockRepo := NewAccountRepository()
	_, errOnNil := mockRepo.SaveTransaction(nil)
	if errOnNil == nil {
		t.Error("Error expected on nil param")
	}
	tx1 := &domain.TransactionOp{
		Type:   constants.DEBIT.String(),
		Amount: -1,
	}
	_, errOnNeg := mockRepo.SaveTransaction(tx1)
	if errOnNeg == nil {
		t.Error("Error expected on negative amount")
	}
	tx2 := &domain.TransactionOp{
		Type:   constants.DEBIT.String(),
		Amount: 100,
	}
	_, errOnNSF := mockRepo.SaveTransaction(tx2)
	if errOnNSF == nil || !strings.Contains(errOnNSF.Error(), "funds") {
		t.Error("Error expected on non-sufficient funds")
	}
}

func TestGetAccountBalance(t *testing.T) {
	mockRepo := NewAccountRepository()
	balance := mockRepo.GetAccountBalance()
	if balance.Balance != 0 {
		t.Error("Balance should start at 0")
	}
}
