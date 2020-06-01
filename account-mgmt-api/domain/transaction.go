package domain

import (
	"github.com/dzisiukluciano/account-mgmt/account-mgmt-api/constants"
)

// Transaction model used to store and return transactions
type Transaction struct {
	ID            string  `json:"id"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	EffectiveDate string  `json:"effectiveDate"`
}

// TransactionOp model used to create transactions
type TransactionOp struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

// Delta returns the amount to add to balance
func (tx *Transaction) Delta() float64 {
	value := tx.Amount
	if tx.Type == constants.DEBIT.String() {
		return -value
	}
	return value
}
