package utilities

import (
	"fmt"
	"net"
)

type ARPEntry struct {
	IP           net.IP
	HardwareAddr net.HardwareAddr
}

func ARP() {
	arpTable, err := getARPTable()
	if err != nil {
		fmt.Println("Failed to retrieve ARP table:", err)
		return
	}

	fmt.Println("ARP table:")
	for _, entry := range arpTable {
		fmt.Printf("IP Address: %s, MAC Address: %s\n", entry.IP, entry.HardwareAddr)
	}
}

func getARPTable() (arpTable []ARPEntry, err error) {

	whoop, err := net.InterfaceByName("en0")
	if err != nil {
		return nil, err
	}

	ifaces := []*net.Interface{whoop}

	for _, iface := range ifaces {
		fmt.Println("interface", iface.Name)
		// Skip loopback and inactive interfaces
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		fmt.Println("addrs", addrs)

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)

			switch v := addr.(type) {
			case *net.IPAddr:
				fmt.Printf("%v : %s (%s)\n", iface.Name, v, v.IP.DefaultMask())

			case *net.IPNet:
				fmt.Printf("%v : %s [%v/%v]\n", iface.Name, v, v.IP, v.Mask)
			}

			if ok {
				// Retrieve ARP table entry for the interface's IP address
				arpEntry := ipNet.IP
				if err != nil {
					return nil, err
				}

				arpTable = append(arpTable, ARPEntry{
					IP:           arpEntry,
					HardwareAddr: iface.HardwareAddr,
				})
			}
		}
	}

	return arpTable, nil
}
