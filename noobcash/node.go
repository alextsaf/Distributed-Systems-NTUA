package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
)

type NodeData struct {
	id        int
	ip        string
	port      string
	publicKey rsa.PublicKey
}

type Node struct {
	blockchain   []HashedBlock
	ip           string
	port         string
	nodeDataList []NodeData
	id           int
	wallet       Wallet
	isBootstrap  bool
	difficulty   uint
	capacity     uint

	globalUTXO map[rsa.PublicKey]UTXOStack
}

func (node *Node) broadcastNodeData(nodeData []NodeData) {

	nodeDataJSON, _ := json.Marshal(nodeData)

	for _, node := range node.nodeDataList {
		url := fmt.Sprintf("http://%s:%s/broadcastNodeData", node.ip, node.port)

		http.Post(url, "application/json", bytes.NewBuffer(nodeDataJSON))
	}
}

// func registerNodeToRing(node Node) {
// 	if !node.isBootstrap {
// 		fmt.Printf("Not authorized")
// 		return
// 	} else {
// 		Node.
// 			http.Get("http://nodeIP:nodePort/registerNodeToRing?id=id")

// 		// transaction := createTransaction(100 -> node)

// 		// Sign?
// 		broadcastTransaction(transaction)
// 	}
// }

// nodes := []Node

// func initializeNetworkConnections() {
// 	if "node" == "bootstrap" {
// 		id := 1

// 		mux := http.NewServeMux()
// 		mux.HandleFunc("/requestForRegister", func(w http.ResponseWriter, r *http.Request) {
// 			newNode := Node{
// 				blockchain: nil,
// 				ip:         r.URL.Query().Get("ip"),
// 				id:         cnt,
// 				port:       r.URL.Query().Get("port"),
// 				publicKey:  r.URL.Query().Get("publicKey"),
// 			}

// 			nodes.append(newNode)

// 			w.Write(id)
// 			id++
// 		})

// 		server := http.Server{
// 			Addr:    fmt.Sprintf(":%d", serverPort),
// 			Handler: mux,
// 		}
// 		if err := server.ListenAndServe(); err != nil {
// 			if !errors.Is(err, http.ErrServerClosed) {
// 				fmt.Printf("error running http server: %s\n", err)
// 			}
// 		}

// 		for id < maxID {
// 			// Continuously check if all nodes have entered
// 		}

// 	} else {
// 		http.Get("http://bootstrapNodeIP:bootstrapNodePort/register?publicKey=%d&ip=%d&port=%d")
// 	}
// }

func (node *Node) broadcastTransaction(transaction Transaction) {
	transactionJSON, _ := json.Marshal(transaction)

	for _, nodeData := range node.nodeDataList {
		if nodeData.ip != node.ip {
			url := fmt.Sprintf("http://%s:%s/broadcastTransaction", nodeData.ip, nodeData.port)

			http.Post(url, "application/json", bytes.NewBuffer(transactionJSON))
		}
	}
}

// use of signature and NBCs balance
func (node *Node) validateTransaction(transaction SignedTransaction) bool {
	if !transaction.verifySignature() {
		fmt.Printf("Error verifying Transaction %d signature\n", transaction.transactionData.TransactionID)
		return false
	}

	transactionInputs := transaction.transactionData.TransactionInputs

	sender := transaction.transactionData.SenderAddress
	senderUTXO := node.globalUTXO[sender]

	receiver := transaction.transactionData.ReceiverAddress
	receiverUTXO := node.globalUTXO[receiver]

	var utxoAmount uint = 0

	for _, input := range transactionInputs {

		transactionID := input.previousOutputID
		amount, found := senderUTXO.Contains(transactionID)

		if !found {
			fmt.Printf("Error verifying Transaction %d UTXOs\n", transaction.transactionData.TransactionID)
			return false
		}

		utxoAmount += amount
	}

	if uint(utxoAmount) < transaction.transactionData.Amount {
		fmt.Printf("Not sufficient amount in Transaction %d wallet\n", transaction.transactionData.TransactionID)
		return false
	}

	// Delete only the ones found in the TransactionInput list :)
	// senderUTXO.flush()

	// We need to create a system to include the UTXOs
	senderUTXO.Push(transaction.transactionData.TransactionOutputs[1])
	receiverUTXO.Push(transaction.transactionData.TransactionOutputs[0])

	return true
}

// if enough transactions  mine
func addTransactionToBlock() {

}

func (node Node) broadcastBlock(block Block) {
	blockJSON, _ := json.Marshal(block)

	// NOTE: Don't broadcast to self
	for _, nodeData := range node.nodeDataList {
		nodeIP := nodeData.ip
		nodePort := nodeData.port

		url := fmt.Sprintf("http://%s:%s/broadcastBlock", nodeIP, nodePort)

		http.Post(url, "application/json", bytes.NewBuffer(blockJSON))
	}
}

func (block HashedBlock) validateBlock() bool {
	return false
}

// Check for the longer chain across all nodes
func (node *Node) validateChain() bool {
	for _, hashedBlock := range node.blockchain {

		if !hashedBlock.validateBlock() {
			return false
		}
	}

	return true
}

// Resolve correct chain
func resolveConflicts() {

}
