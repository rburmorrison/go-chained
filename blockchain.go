package chained

import (
	"strconv"
	"strings"
	"time"
)

// Blockchain represents a structure that links
// blocks together through hashing. It also contains
// a transaction pool of transactions to be put onto
// the next block as well as an identifier.
type Blockchain struct {
	Blocks          []Block
	TransactionPool []Transaction
	ChainID         string
	Peers           []Peer
	creationTime    time.Time
}

// NewBlockchain will create and return a new
// blockchain with a genesis block and an initial
// transaction. The genesis block will work to have
// a valid block, and may take some time, depending
// on the target difficulty, to create.
func NewBlockchain() Blockchain {
	// Create genesis transaction and block
	ph := strings.Repeat("0", HashStringSize)
	t := Transaction{"root", "root", "GENESIS", []byte{}, time.Time{}}
	block := Block{13466, []Transaction{t}, ph, time.Time{}}

	// Find right nonce
	for !block.IsValid() {
		block.Nonce++
	}

	var b Blockchain
	b.Blocks = []Block{block}
	b.TransactionPool = []Transaction{}
	b.ChainID = hash512String([]byte(time.Now().String() + strconv.Itoa(block.Nonce)))[:IDLength]
	b.Peers = []Peer{}
	b.creationTime = time.Now()

	return b
}

// NewBlockchainWithIdentifier will create a new
// blockchain in the usual way, but allows for the
// assignment of a custom identifier. The identifier
// must have a length of IDLength, or it will fail
// its validation check.
func NewBlockchainWithIdentifier(id string) Blockchain {
	b := NewBlockchain()
	b.ChainID = id
	return b
}

// VerifiedTransactions will iterate through the
// blocks on the blockchain and extract all of the
// transactions out of them.
func (b Blockchain) VerifiedTransactions() []Transaction {
	var trans []Transaction

	for _, b := range b.Blocks {
		trans = append(trans, b.Transactions...)
	}

	return trans
}

// ResolveChain accepts a blockchain and will compare
// it to the received blockchain to determine if it
// is superior or not. A blockchain is superior if
// it has more blocks and is fully valid. If the
// passed blockchain is determined to be superior,
// all blocks on the received blockchain will be
// replaced by the superior one. ResoveChain will
// return true if the chain was replaced, and false
// otherwise.
func (b *Blockchain) ResolveChain(b2 Blockchain) bool {
	if len(b2.Blocks) > len(b.Blocks) && b2.IsValid() {
		b.Blocks = b2.Blocks
		return true
	}

	return false
}

// ResolveTransactions will accept a list of
// transactions and add any new ones to the received
// blockchain's transaction pool. No transactions
// will be added if any of the transactions in the
// list of new transactions are invalid. True will
// be returned if any new transactions were added to
// the pool.
func (b *Blockchain) ResolveTransactions(ts []Transaction) bool {
	var tAdded bool

	// All transactions must be valid
	for _, t := range ts {
		if !t.IsValid() {
			return false
		}

		// Determine if received blockchain contains this
		// transaction.
		var contains bool
		tHash := hash512String([]byte(t.JSONString()))

		for _, t2 := range b.TransactionPool {
			t2Hash := hash512String([]byte(t2.JSONString()))
			if tHash == t2Hash {
				// Equivalent transaction found
				contains = true
				break
			}
		}

		if !contains {
			b.TransactionPool = append(b.TransactionPool, t)
			tAdded = true
		}
	}

	return tAdded
}

// AddBlock will attempt to add a block to the
// blockchain. If it is invalid, the block will not
// be added, and an error will be returned.
func (b *Blockchain) AddBlock(block Block) error {
	if !block.IsValid() {
		return ErrInvalidData
	}

	b.Blocks = append(b.Blocks, block)
	return nil
}

// AddTransaction will attempt to add a transaction
// to the blockchain. If it is invalid, the
// transaction will not be added, and an error will
// be returned.
func (b *Blockchain) AddTransaction(t Transaction) error {
	if !t.IsValid() {
		return ErrInvalidData
	}

	b.TransactionPool = append(b.TransactionPool, t)
	return nil
}

// AddPeer will attempt to add a peer to the
// blockchain. If it is invalid, the peer will not
// be added, and an error will be returned.
func (b *Blockchain) AddPeer(peer Peer) error {
	if !peer.IsValid() {
		return ErrInvalidData
	}

	b.Peers = append(b.Peers, peer)
	return nil
}

// MineNewBlock will work towards creating a new
// valid block containing tranactions from the
// transaction pool by incrementing the nonce value
// until the block passes validation. A reward
// transaction will be added to the block to indicate
// who the block was mined by, and the timestamp of
// the transaction will indicate when the block was
// mined. The valid block will be returned.
func (b Blockchain) MineNewBlock() Block {
	block := NewEmptyBlock()
	block.Timestamp = time.Now()
	block.Transactions = b.TransactionPool
	block.PreviousHash = b.LastBlock().Hash()

	// Create and add reward transaction
	rw, _ := NewTransactionNow(b.ChainID, "root", "MINED")
	block.Transactions = append(block.Transactions, rw)

	for !block.IsValid() {
		block.Nonce++
		block.Timestamp = time.Now()

		// Continuously update the list of transactions in
		// the case that any new transactions have been added
		// within the time it took to mine the block.
		block.Transactions = b.TransactionPool
		block.Transactions = append(block.Transactions, rw)
	}

	return block
}

// MineNewBlockAndApply will mine a new block with
// MineNewBlock and will immediately add it to the
// blockchain when found. The transaction pool will
// be flushed just before the new block is added.
func (b *Blockchain) MineNewBlockAndApply() {
	block := b.MineNewBlock()

	b.TransactionPool = []Transaction{}
	b.Blocks = append(b.Blocks, block)
}

// LastBlock will return the previously mined block
// on the received chain.
func (b Blockchain) LastBlock() Block {
	return b.Blocks[len(b.Blocks)-1]
}

// CreationTime Returns the time that the blockchain
// was created.
func (b Blockchain) CreationTime() time.Time {
	return b.creationTime
}

// IsValid inspects the fields of the received
// blockchain and determines if its contents are
// valid.
//
// A blockchain must have all valid blocks, have all
// valid transactions, and must have a ChainID of
// the correct length.
func (b Blockchain) IsValid() bool {
	// All blocks must be valid
	for i, block := range b.Blocks {
		if !block.IsValid() {
			return false
		}

		// Ensure previous hash matches
		if i == 0 {
			ph := strings.Repeat("0", HashStringSize)
			if block.PreviousHash != ph {
				return false
			}
		} else {
			if block.PreviousHash != b.Blocks[i-1].Hash() {
				return false
			}
		}
	}

	// All transactions in the transaction pool must be
	// valid
	for _, t := range b.TransactionPool {
		if !t.IsValid() {
			return false
		}
	}

	// ChainID must be of the proper length
	if len(b.ChainID) != IDLength {
		return false
	}

	return true
}

// JSONString converts the received blockchain to a
// JSON string and returns it.
func (b Blockchain) JSONString() string {
	s, _ := toJSONString(b)
	return s
}
