package main

import (
	"net"
)

import (
	"log"
)

func main() {
	c, err := net.ListenPacket("udp", ":8032")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for {
		b := make([]byte, 512)
		n, peer, err := c.ReadFrom(b)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("read: %s\n", string(b[:n]))
		if _, err := c.WriteTo([]byte("Servus from server"), peer); err != nil {
			log.Fatal(err)
		}
	}
}
