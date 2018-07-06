package chained

import (
	"strings"
	"time"
)

// Transaction represents an interaction between two
// nodes on the blockchain.
type Transaction struct {
	Recipient string
	Sender    string
	Message   string
	Data      []byte
	Timestamp time.Time
}

// NewTransaction will create and validate a
// transaction with the given parameters. An error
// will be returned if the transaction is invalid.
func NewTransaction(r, s, m string, ts time.Time) (Transaction, error) {
	trans := Transaction{r, s, m, []byte{}, ts}

	if !trans.IsValid() {
		return Transaction{}, ErrInvalidData
	}

	return trans, nil
}

// NewTransactionNow creates a transaction with
// NewTransaction, but sets the timestamp to
// time.Now().
func NewTransactionNow(r, s, m string) (Transaction, error) {
	return NewTransaction(r, s, m, time.Now())
}

// IsValid inspects the fields of the received
// transaction and determines if its contents are
// valid.
//
// A transaction must have no empty fields.
func (t Transaction) IsValid() bool {
	recipeintEmpty := strings.TrimSpace(t.Recipient) == ""
	senderEmpty := strings.TrimSpace(t.Sender) == ""
	messageEmpty := strings.TrimSpace(t.Message) == ""

	if recipeintEmpty || senderEmpty || messageEmpty {
		return false
	}

	return true
}

// JSONString converts the received transaction to
// a JSON string and returns it.
func (t Transaction) JSONString() string {
	s, _ := toJSONString(t)
	return s
}
