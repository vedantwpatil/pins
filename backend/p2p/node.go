package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {
	// Create default libp2p host
	host, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	fmt.Println("Node ID:", host.ID())
	fmt.Println("Listening addresses:", host.Addrs())
}
