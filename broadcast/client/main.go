package main

import (
	"log"
	"net"
	"sync"
	"time"
)

func FindActiveIPs() ([]string, error) {
	var addresses []string
	// list system network interfaces
	// https://golang.org/pkg/net/#Interfaces
	intfs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	// mapping between network interface name and index
	// https://golang.org/pkg/net/#Interface
	for _, intf := range intfs {
		// skip down interface & check next intf
		if intf.Flags&net.FlagUp == 0 {
			continue
		}
		// skip loopback & check next intf
		if intf.Flags&net.FlagLoopback != 0 {
			continue
		}
		// list of unicast interface addresses for specific interface
		// https://golang.org/pkg/net/#Interface.Addrs
		addrs, err := intf.Addrs()
		if err != nil {
			return nil, err
		}
		// network end point address
		// https://golang.org/pkg/net/#Addr
		for _, addr := range addrs {
			// if for windows may need to type switch

			// type assertion to access Addr interface
			// underlying IPNet IP method
			if addr == nil || addr.(*net.IPNet).IP.IsLoopback() {
				continue
			}
			// append active interfaces
			addresses = append(addresses, addr.String())
		}
	}
	return addresses, nil
}

func main() {
	localIps, err := FindActiveIPs()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	discoveredIps := make(chan net.IPNet, len(localIps))

	c, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for _, localIp := range localIps {
		ipv4Addr, ipv4Net, err := net.ParseCIDR(localIp)
		if err != nil {
			panic(err)
		}

		ipv4Addr = ipv4Addr.To4()

		if ipv4Addr == nil {
			continue
		}

		wg.Add(1)

		go func(ipv4Addr net.IP, ipv4Net *net.IPNet) {
			defer wg.Done()

			log.Printf("----------------")
			log.Printf("ip: %v", ipv4Addr)
			log.Printf("subnet: %v", ipv4Net)

			ones, bits := ipv4Net.Mask.Size()
			mask := net.CIDRMask(ones, bits)

			broadcast := net.IP(make([]byte, 4))
			for i := range ipv4Addr {
				broadcast[i] = ipv4Addr[i] | ^mask[i]
			}

			log.Printf("broadcast: %v", broadcast.String())

			dst, err := net.ResolveUDPAddr("udp", broadcast.String()+":8032")
			if err != nil {
				log.Fatal(err)
			}

			if _, err := c.WriteTo([]byte("hello from client"), dst); err != nil {
				log.Fatal(err)
			}
		}(ipv4Addr, ipv4Net)
	}

	wg.Wait()

	log.Printf("reading answers ...")

	b := make([]byte, 512)
	c.SetReadDeadline(time.Now().Add(time.Second * 1))
	for {
		n, peer, err := c.ReadFrom(b)
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				break
			} else {
				log.Fatal(err)
			}
		}
		log.Printf("%d bytes read from %+v: %s\n", n, peer, string(b[:n]))
	}

	close(discoveredIps)

	for discoveredIp := range discoveredIps {

		log.Printf("discovered ip: %+v", discoveredIp)
	}
}
