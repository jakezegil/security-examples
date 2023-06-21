package utilities

import (
	"fmt"
	"net"
	"time"
)

var ipRange = "192.168.1.158/24"

func ScanIPs(cidr *string) {
	if cidr != nil {
		ipRange = *cidr
	}

	// Parse the IP range
	ip, ipNet, _ := net.ParseCIDR(ipRange)

	// Iterate through IP addresses in the range
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		go func(ip net.IP) {
			// Send ICMP echo request (ping)
			err := ping(ip.String(), time.Second*1)
			if err == nil {
				fmt.Printf("IP address %s is open\n", ip)
			}

		}(ip)
	}

	time.Sleep(time.Second * 5) // Wait for the pings to complete
}

// Helper function to increment an IP address
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// Function to send ICMP echo request
func ping(ip string, timeout time.Duration) error {
	conn, err := net.DialTimeout("ip4:icmp", ip, timeout)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}
