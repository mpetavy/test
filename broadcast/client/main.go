package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func getBroadcast() {
	mask := net.CIDRMask(24, 32)
	ip := net.IP([]byte{192, 168, 1, 3})

	broadcast := net.IP(make([]byte, 4))
	for i := range ip {
		broadcast[i] = ip[i] | ^mask[i]
	}

	fmt.Println(broadcast)
}


func main() {
	c, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	dst, err := net.ResolveUDPAddr("udp", "255.255.255.255:8032")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := c.WriteTo([]byte("hello from client"), dst); err != nil {
		log.Fatal(err)
	}
	b := make([]byte, 512)
	c.SetReadDeadline(time.Now().Add(time.Second * 1))
	for {
		n, peer, err := c.ReadFrom(b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n, "bytes read from", peer)
		log.Printf("read: %s\n", string(b[:n]))
	}
}
