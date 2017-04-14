package main

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	s := ""
	for i, v := range ip {
		dot := ""
		if i != len(ip)-1 {
			dot = "."
		}
		s = s + fmt.Sprintf("%v", v) + dot
	}
	return fmt.Sprintf(s)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
