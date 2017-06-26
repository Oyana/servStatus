package main

import (
	"fmt"

	"net"
)

func main() {
	fmt.Println("PlopIplop")
	domainList := [...] string {
		"bulko.net",
		"koregraf.com",
		"golga.io" }

	for i := 0; i < len(domainList); i++ {
		addr, err := net.LookupHost(domainList[i])
		fmt.Println(domainList[i], addr, err)
	}
}
