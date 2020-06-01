package constants

// TransactionType enum (credit, debit)
type TransactionType int

// TransactionType const
const (
	CREDIT TransactionType = iota
	DEBIT
)

var types = [...]string{"credit", "debit"}

// String() function will return the string representation
// of a TransactionType
func (t TransactionType) String() string {
	return types[t]
}
