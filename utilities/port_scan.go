package utilities

import (
	"fmt"
	"net"
	"strconv"
)

var targetIP = "192.168.1.158"

func ScanPorts() {
	for port := 1; port <= 65535; port++ {
		address := targetIP + ":" + strconv.Itoa(port)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			// Port is open
			fmt.Printf("Port %d is open\n", port)
			conn.Close()
		}
	}
}
