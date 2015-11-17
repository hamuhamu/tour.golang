package main

import "fmt"

type IPAddr [4]byte

func (addrs IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", addrs[0], addrs[1], addrs[2], addrs[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
