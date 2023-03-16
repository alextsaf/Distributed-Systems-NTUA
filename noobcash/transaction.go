package main

import (
	"crypto"
	"crypto/rand"   //random generator and Reader
	"crypto/rsa"    //public key
	"crypto/sha256" //sha256 hashing

	// uint random generator
	"encoding/json" // json encoding/decoding Marshal
	"fmt"
)

/*
Μην αλλάξετε το capitalization των λεξεων, πρεπει να ξεκινάνε με
κεφαλαίο για να γίνει σωστά το json encoding, αλλιώς δεν γίνονται
extracted τα variables
*/
type Transaction struct {
	SenderAddress      rsa.PublicKey
	ReceiverAddress    rsa.PublicKey
	Amount             uint
	TransactionID      uint64
	TransactionInputs  []TransactionInput
	TransactionOutputs [2]TransactionOutput
}

type SignedTransaction struct {
	transactionData Transaction
	signature       []byte
}

type TransactionOutput struct {
	ID                  uint64
	parentTransactionID uint64        //the id of the transaction this output was created in
	recipient           rsa.PublicKey //also known as the new owner of these coins.
	amount              uint          //the amount of coins they own
}

type TransactionInput struct {
	previousOutputID uint64 //Reference to TransactionOutputs -> transactionId
}

func generateTransaction(amount uint, sendersWallet Wallet, receiverAddress rsa.PublicKey) Transaction {
	var UTXOsTotal uint
	UTXOsTotal = 0

	var activeUTXOs []TransactionOutput
	var poppedUTXO TransactionOutput
	var isEmpty bool

	// Go over the list until I find a sum that is >= the amount
	for UTXOsTotal < amount {
		poppedUTXO, isEmpty = sendersWallet.walletUTXOs.Pop()

		if isEmpty {
			fmt.Println("Empty UTXO Queue - not enough available funds")
			// Re-add transcations
			for _, poppedUTXO := range activeUTXOs {
				sendersWallet.walletUTXOs.Push(poppedUTXO)
			}

			return Transaction{}
		}
		UTXOsTotal += poppedUTXO.amount
		activeUTXOs = append(activeUTXOs, poppedUTXO)
	}

	fmt.Println(UTXOsTotal)

	var transInputs []TransactionInput

	// On god mhn ksexasete pote to _, efaga 25 lepta debugging giati nomiza oti h Pop() epestrefe int
	// enw htan to index gia to range. Mia xara htan h C, hkseran ti ekanan 50 xronia twra
	for _, poppedUTXO := range activeUTXOs {
		transInputs = append(transInputs, TransactionInput{poppedUTXO.ID})
	}

	// Create the transaction outputs
	var transactionId, receiverOutputId, senderOutputId uint64
	transactionId = generateRandomNumber()
	receiverOutputId = generateRandomNumber()
	senderOutputId = generateRandomNumber()

	receiverTransactionOutput := TransactionOutput{
		ID:                  receiverOutputId,
		parentTransactionID: transactionId,
		recipient:           receiverAddress,
		amount:              amount,
	}

	senderTransactionOutput := TransactionOutput{
		ID:                  senderOutputId,
		parentTransactionID: transactionId,
		recipient:           sendersWallet.publicKey,
		amount:              UTXOsTotal - amount,
	}

	transOutput := [2]TransactionOutput{
		receiverTransactionOutput,
		senderTransactionOutput}

	// Create the unsigned Transaction struct and return it
	transaction := Transaction{
		SenderAddress:      sendersWallet.publicKey,
		ReceiverAddress:    receiverAddress,
		Amount:             amount,
		TransactionID:      transactionId,
		TransactionInputs:  transInputs,
		TransactionOutputs: transOutput,
	}

	return transaction
}

func calculateHash(transactionData Transaction) [32]byte {
	transactionJSON, err := json.Marshal(transactionData)
	if err != nil {
		fmt.Printf("Error encoding transaction to JSON :( %v\n", err)
	}

	transactionHash := sha256.Sum256(transactionJSON) //returns a byte slice ([]byte)
	return transactionHash
}

func (trans Transaction) signTransaction(signeeWallet Wallet) SignedTransaction {
	transactionHash := calculateHash(trans)

	// create a signature with the wallets privateKey
	signature, err := rsa.SignPKCS1v15(rand.Reader, &signeeWallet.privateKey, crypto.SHA256, transactionHash[:])
	if err != nil {
		fmt.Printf("Error signing transaction :( %s\n", err)
	}

	signedTrans := SignedTransaction{
		transactionData: trans,
		signature:       signature,
	}

	return signedTrans
}

func (signedTrans SignedTransaction) verifySignature() bool {
	transactionHash := calculateHash(signedTrans.transactionData)

	verificationReturn := rsa.VerifyPKCS1v15(&signedTrans.transactionData.SenderAddress,
		crypto.SHA256, transactionHash[:], signedTrans.signature)

	if verificationReturn != nil {
		fmt.Printf("Transaction %d verification failed with error %s", signedTrans.transactionData.TransactionID, verificationReturn)
		return false
	} else {
		return true
	}
}
