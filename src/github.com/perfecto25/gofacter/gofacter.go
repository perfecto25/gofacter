package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// will need to refactor

	// ------ HOSTNAME

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hostmane: ", hostname)

	// ------- IP ADDRESS

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}

}
