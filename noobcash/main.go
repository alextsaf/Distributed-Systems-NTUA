package main

import "fmt"

func main() {
	_, thimaPagwnaKey := generateKeyPair(2048)
	portofoliPagwna := generateWallet()
	leftaPagwna := TransactionOutput{
		ID:                  69,
		parentTransactionID: 420,
		recipient:           portofoliPagwna.publicKey,
		amount:              1000,
	}

	leftaPagwna2 := TransactionOutput{
		ID:                  619,
		parentTransactionID: 4220,
		recipient:           portofoliPagwna.publicKey,
		amount:              69,
	}

	transOut := [2]TransactionOutput{
		{
			ID:                  1234,
			parentTransactionID: 321,
			recipient:           portofoliPagwna.publicKey,
			amount:              10000 - 1000,
		},
		{
			ID:                  12345,
			parentTransactionID: 321,
			recipient:           thimaPagwnaKey,
			amount:              1000,
		},
	}

	var transIn []TransactionInput
	transIn = append(transIn, TransactionInput{
		previousOutputID: 100321})

	fakeTransaction := Transaction{
		SenderAddress:      thimaPagwnaKey,
		ReceiverAddress:    portofoliPagwna.publicKey,
		Amount:             100,
		TransactionID:      42,
		TransactionInputs:  transIn,
		TransactionOutputs: transOut,
	}

	portofoliPagwna.walletUTXOs.Push(leftaPagwna)
	pcbudget := balance(portofoliPagwna)
	fmt.Println(pcbudget)

	portofoliPagwna.walletUTXOs.Push(leftaPagwna2)
	pcbudget2 := balance(portofoliPagwna)
	fmt.Println(pcbudget2)

	unhashedTrans := generateTransaction(10, portofoliPagwna, thimaPagwnaKey)

	fakeHashTrans := fakeTransaction.signTransaction(portofoliPagwna)

	hashedTransaction := unhashedTrans.signTransaction(portofoliPagwna)

	einaiTouPagwnaTaLefta := hashedTransaction.verifySignature()
	einaiTouPagwnaTaLefta2 := fakeHashTrans.verifySignature()

	fmt.Printf("\n")
	fmt.Println(einaiTouPagwnaTaLefta)
	fmt.Println(einaiTouPagwnaTaLefta2)
}
