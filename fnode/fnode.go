//This package has functions that retrieve information from Callisto Node
package fnode

import (
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	//"github.com/eduartua/callisto/Galileo/blocks"
	"fmt"
)

var lastBlock  map[string]interface{}

func getLastBlock() {
	// Connect the client
	client, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("could not create ipc client: %v", err)
	}
	err = client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		fmt.Println("Error getting latest block:", err)
		return
	}

	for k, v := range lastBlock {
		fmt.Printf("%v: %v\n", k, v)
	}
}


