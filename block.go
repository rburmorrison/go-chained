package chained

import (
	"strings"
	"time"
)

// Block represents a group of transactions secured
// with hashing.
type Block struct {
	Nonce        int
	Transactions []Transaction
	PreviousHash string
	Timestamp    time.Time
}

// NewBlock will create and validate a block with the
// given parameters. An error will be returned if the
// block is invalid.
func NewBlock(n int, trans []Transaction, ph string, ts time.Time) (Block, error) {
	b := Block{n, trans, ph, ts}

	if !b.IsValid() {
		return Block{}, ErrInvalidData
	}

	return b, nil
}

// NewEmptyBlock will create a new block with default
// fields. This includes the timestamp. The timestamp
// will likely need to be set to something
// reasonable.
func NewEmptyBlock() Block {
	ph := strings.Repeat("0", HashStringSize)
	b, _ := NewBlock(0, []Transaction{}, ph, time.Time{})
	return b
}

// AddTransaction will add a transaction to the block
// and will return an error if the transaction is
// invalid.
func (b *Block) AddTransaction(t Transaction) error {
	if !t.IsValid() {
		return ErrInvalidData
	}

	// Add transaction
	b.Transactions = append(b.Transactions, t)

	return nil
}

// Hash will return the hash of the received block's
// JSON string equivalent.
func (b Block) Hash() string {
	return hash512String([]byte(b.JSONString()))
}

// IsValid inspects the fields of the received
// block and determines if its contents are valid.
//
// A block must have all valid transactions, have
// the proper previous hash length, and must contain
// the correct amount of leading zeros.
func (b Block) IsValid() bool {
	// All transactions must be valid
	for _, t := range b.Transactions {
		if !t.IsValid() {
			return false
		}
	}

	// Previous hash must have length equal to
	// HashStringSize constant.
	if len(b.PreviousHash) != HashStringSize {
		return false
	}

	// Number of leading zeros in hash must be at least
	// as many as Target.
	if !strings.HasPrefix(b.Hash(), strings.Repeat("0", Target)) {
		return false
	}

	return true
}

// JSONString converts the received block to a JSON
// string and returns it.
func (b Block) JSONString() string {
	s, _ := toJSONString(b)
	return s
}
