package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// For testing purposes (remove later)
// type Node struct {
// 	id          int
// 	isBootstrap bool
// 	ip          string
// 	port        string
// 	wallet      Wallet
// }

// type Wallet struct {
// 	publicKey string
// }

func bootStrapMain() {
	// Bootstrap phase
	bootstrapIP := "192.168.0.1"
	bootstrapPort := "50000"

	// node := Node{
	// 	blockchain:  Blockchain{},
	// 	ip:          os.Args[1],
	// 	port:        os.Args[2],
	// 	nodes:       nil,
	// 	id:          0,
	// 	wallet:      Wallet{},
	// 	isBootstrap: os.Args[3] == "bootstrap",
	// }

	node := Node{
		ip:          os.Args[1],
		port:        os.Args[2],
		id:          0,
		wallet:      generateWallet(),
		isBootstrap: os.Args[3] == "bootstrap",
	}

	if node.isBootstrap {
		// Listen for incoming nodes
		id := 1

		http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
			// ip := r.URL.Query().Get("ip")
			// port := r.URL.Query().Get("port")
			// publicKey := r.URL.Query().Get("publicKey")

			w.Write([]byte(string(id)))
			id++
		})

		listenPort := fmt.Sprintf(":%s", node.port)
		go http.ListenAndServe(listenPort, nil)

		fmt.Println("Hello")

		for id < 5 {

		}

	} else {
		// Request to be part of ring
		url := fmt.Sprintf("http://%s:%s/join?ip=%s&port=%s&publicKey=%s",
			bootstrapIP, bootstrapPort, node.ip, node.port, node.wallet.publicKey)

		res, _ := http.Get(url)

		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(body[0])
	}

}
