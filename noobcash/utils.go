package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func generateRandomNumber() uint64 {
	var randomNum uint64

	seed := make([]byte, 8)
	_, _ = rand.Read(seed)
	randomNum = binary.BigEndian.Uint64(seed[:])

	return randomNum
}

func generateRandomNonce() [32]byte {
	var nonce [32]byte
	_, _ = rand.Read(nonce[:])
	// fmt.Println(nonce)
	return nonce
}

func getRequest(requestURL string) string {

	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("ioutil.ReadAll() error %s", err)
	}

	return string(body)
}

func postRequest(requestURL string, body io.ReadCloser) {
	http.Post(requestURL, "application/json", body)
}
