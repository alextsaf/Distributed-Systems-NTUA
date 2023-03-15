package main

import (
	"fmt"			//stdio.h xa0x0a0xa0xa0a0
	"crypto/rsa" 	//public key
	"crypto/rand"  	//rand generator (to lene reader genika, kalh fash)
	"os" 			//ela pou den ksereis
)

/*
	Somewhat of an over-engineering move. You can just add the key pair
	in the Node struct, but this is more clear and organized
*/

type Wallet struct {
	privateKey 	rsa.PrivateKey //rsa struct that holds all the information
	publicKey  	rsa.PublicKey // rsa.PrivateKey.PublicKey (struct city :3)
	walletUTXOs	UTXOStack
}

// Function generating a new Keypair of public and private key for this wallet
func generateKeyPair(RSABytes int) (rsa.PrivateKey, rsa.PublicKey) {
	privateKeyAddress, err := rsa.GenerateKey(rand.Reader, RSABytes);
	if err != nil {
		fmt.Printf("Error when creating Wallet RSA KeyPair\n");
		os.Exit(1);
	}

	privateKey := *privateKeyAddress;
	publicKey := privateKey.PublicKey;

	return privateKey,publicKey
}

func generateWallet() Wallet {
	privateKey, publicKey := generateKeyPair(2048);
	return Wallet{privateKey: privateKey, publicKey: publicKey, walletUTXOs: make(UTXOStack,0)};
}

func balance(wallet Wallet) uint {
	return wallet.walletUTXOs.SumAmounts();
}

func sendFunds() {

}

func sendNudes() {

	// :3
}