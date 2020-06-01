package service

import (
	"strings"
	"testing"

	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/constants"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/domain"
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/usecase/repository"
)

func TestGetAccountBalance(t *testing.T) {
	mockRepo := repository.NewAccountRepository()
	mockService := NewAccountService(mockRepo)
	balance, apiError := mockService.GetAccountBalance()
	if apiError != nil {
		t.Error("GetAccountBalance should not fail")
	}
	if balance.Balance != 0 {
		t.Error("Balance should start at 0")
	}
}

func TestSaveTransaction(t *testing.T) {
	mockRepo := repository.NewAccountRepository()
	mockService := NewAccountService(mockRepo)
	tx1 := &domain.TransactionOp{
		Type:   constants.DEBIT.String(),
		Amount: -1,
	}
	_, errOnNeg := mockService.SaveTransaction(tx1)
	if errOnNeg == nil {
		t.Error("Error expected on negative amount")
	}
	tx2 := &domain.TransactionOp{
		Type:   constants.DEBIT.String(),
		Amount: 100,
	}
	_, errOnNSF := mockService.SaveTransaction(tx2)
	if errOnNSF == nil || !strings.Contains(errOnNSF.Message, "funds") {
		t.Error("Error expected on non-sufficient funds")
	}
	tx3 := &domain.TransactionOp{
		Type:   "invalid",
		Amount: 100,
	}
	_, errOnInvalidType := mockService.SaveTransaction(tx3)
	if errOnInvalidType == nil || !strings.Contains(errOnInvalidType.Message, "type") {
		t.Error("Error expected with invalid tx type")
	}
	income := 1000.0
	outcome := 100.0
	tx4 := &domain.TransactionOp{
		Type:   constants.CREDIT.String(),
		Amount: income,
	}
	succTx4, errorOnSuccessCredit := mockService.SaveTransaction(tx4)
	if errorOnSuccessCredit != nil {
		t.Error("Credit operation should be successful")
	}
	if succTx4 == nil || succTx4.Amount != income {
		t.Error("Credit operation should return valid transaction")
	}
	balance, _ := mockService.GetAccountBalance()
	if balance.Balance != income {
		t.Error("Expect balance to be equals income")
	}
	tx5 := &domain.TransactionOp{
		Type:   constants.DEBIT.String(),
		Amount: outcome,
	}
	succTx5, errorOnSuccessDebit := mockService.SaveTransaction(tx5)
	if errorOnSuccessDebit != nil {
		t.Error("Debit operation should be successful")
	}
	if succTx5 == nil || succTx5.Amount != outcome {
		t.Error("Debit operation should return valid transaction")
	}
	updatedBalance, _ := mockService.GetAccountBalance()
	if updatedBalance.Balance != income-outcome {
		t.Errorf("Expect balance to be equals %f", income-outcome)
	}
}
