package server

import (
	"fmt"
	"log"

	"golang.zx2c4.com/wireguard/wgctrl"
)

func main() {
	// initialize the wireguard control client
	client, err := wgctrl.New()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	dev, err := client.Device("wg0")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Device:", dev.Name)
	fmt.Println("Public key:", dev.PublicKey)

	for _, peer := range dev.Peers {
		fmt.Println("Peer:", peer.PublicKey)
	}
}
