package main

// import (
// 	"fmt"
// 	"crypto/rsa" //public key
// 	"time" //time.Now().Unix() -> int64, unix epoch
// 	"crypto/sha256"
// )

type Block struct {
	Index        uint
	Timestamp    int64
	Transactions [transactionCapacity]Transaction
	Nonce        [32]byte
	PreviousHash [32]byte
}

// Same exact issue as with the transaction. You need to keep a copy of the original block to validate it
// This is the simplest way i.m.o
type HashedBlock struct {
	blockData   Block
	currentHash [32]byte
}

func generateGenesisBlock() HashedBlock {

}

// From python template
func generateBlockHash() [32]byte {
	// Calculate self.hash
}

// From python template
func addTransaction(transaction Transaction, block Block) bool {
	// Add a transaction to the block
}
