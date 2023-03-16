package main

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

const transactionCapacity = 5

type Block struct {
	Index        uint
	Timestamp    int64
	Transactions []Transaction
	Nonce        [32]byte
	PreviousHash [32]byte
}

type HashedBlock struct {
	blockData   Block
	currentHash [32]byte
}

func generateGenesisBlock(numberOfNodes uint, bootstrapNode Node) HashedBlock {

	N := numberOfNodes

	genesisTransactionID := generateRandomNumber()

	bootstrapTransactionOutput := TransactionOutput{
		ID:                  generateRandomNumber(),
		parentTransactionID: genesisTransactionID,
		recipient:           bootstrapNode.wallet.publicKey,
		amount:              uint(100 * N),
	}

	genesisTransactionOutput := TransactionOutput{
		ID:                  generateRandomNumber(),
		parentTransactionID: genesisTransactionID,
		recipient:           rsa.PublicKey{N: big.NewInt(0), E: 1},
		amount:              0,
	}

	genesisTransactionOutputs := [2]TransactionOutput{
		bootstrapTransactionOutput,
		genesisTransactionOutput,
	}

	var nonce, previousHash [32]byte

	for i := range nonce {
		nonce[i] = 0
		previousHash[i] = 0
	}
	previousHash[31] = 1

	var genesisTransactions []Transaction

	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: genesisTransactions,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}

	genesisTransactions = append(genesisBlock.Transactions, Transaction{
		SenderAddress:      rsa.PublicKey{},
		ReceiverAddress:    bootstrapNode.wallet.publicKey, /* boostrapNode */
		Amount:             uint(100 * N),
		TransactionID:      genesisTransactionID,
		TransactionInputs:  []TransactionInput{},
		TransactionOutputs: genesisTransactionOutputs,
	})

	genesisBlockHash := generateBlockHash(genesisBlock)

	genesisHashedBlock := HashedBlock{
		blockData:   genesisBlock,
		currentHash: genesisBlockHash,
	}

	return genesisHashedBlock
}

func generateBlockHash(blockToBeHashed Block) [32]byte {

	blockJSON, err := json.Marshal(blockToBeHashed)
	if err != nil {
		fmt.Printf("Error encoding Block to JSON :( %v\n", err)
	}

	blockHash := sha256.Sum256(blockJSON) //returns a byte slice ([32]byte)
	return blockHash
}

/*
As soon as we reach the max capacity for a block, everyone starts mining to find the nonce and
hash it. When someone finds the correct nonce, they broadcast the HashedBlock to the ring.
Everyone then validates the block (see below) and adds it to the blockchain.

The addition to the blockchain may have conflicts (different lenghts of chains). Conflicts are
dealt with by chosing the biggest chain length (without gaps).
*/

func mineBlock(blockToBeMined Block, miner Node) HashedBlock {
	fmt.Printf("Miner ID: %d started mining\n", miner.id)
	/*
	   We need to find a nonce number that starts with {difficulty} number of 0s
	   Create a block with all the information that we have so far and hash it, with different
	   nonce values. If the hash complies with our difficulty level, accept this has as the
	   currentHash of the block.
	*/
	for {
		// Need to check periodically if new block has been added to the blockchain and cancel our
		// mining operation. (Skeftomai go routine alla exoume mono 1 pyrhna sta 10 nodes ara idk :( )
		nonce := generateRandomNonce()
		blockToBeMined.Nonce = nonce
		blockHash := generateBlockHash(blockToBeMined)

		difficultyFlag := true
		// Check difficulty levels
		for i := range nonce[0:miner.difficulty] {
			if nonce[i] != 0 {
				difficultyFlag = false
				break
			}
		}

		// If it abides to our difficulty rule, then return the hashed block
		if difficultyFlag {
			return HashedBlock{
				blockData:   blockToBeMined,
				currentHash: blockHash,
			}
		}
	}
}

func (node *Node) validateBlock(block HashedBlock) bool {
	blockData := block.blockData
	blockIndex := blockData.Index

	currentBlockchainLength := uint(len(node.blockchain))

	// Automatically validate genesis block, its accepted as valid every time
	if currentBlockchainLength == 0 && blockIndex == 0 {

	} else if currentBlockchainLength > blockIndex {
		return false
	} else if currentBlockchainLength < blockIndex {
		// Kalw thn resolve conflict giati yparxei endexomeno na exei mpei kapoio allo block
		// prin ginei to validate aftou
		return false //vazw proswrina return false mexri na ly8ei
	}

	// If the block is correct one, check its validity by comparing the hashes
	blockHash := generateBlockHash(blockData)

	if blockHash != block.currentHash {
		fmt.Printf("Block %d current hash verification failed ", blockIndex)
	}
	previousBlock := node.blockchain[currentBlockchainLength-1]

	if blockData.PreviousHash != previousBlock.currentHash {
		fmt.Printf("Block %d current hash verification failed ", blockData.Index)
		return false
	}

	// The block is accepted - we need to add it to the blockchain

	// TODO: add block to blockchain

	return true
}
