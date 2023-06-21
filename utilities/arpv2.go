package utilities

// import (
// 	"fmt"
// 	"net"
// 	"strings"
// 	"time"
// )

// func GetARP() {
// 	rokuMAC := "xx:xx:xx:xx:xx:xx" // Replace with the MAC address of your Roku device
// 	timeout := 2 * time.Second     // Adjust the timeout value based on your network

// 	rokuIP, err := findRokuIP(rokuMAC, timeout)
// 	if err != nil {
// 		fmt.Println("Failed to find Roku IP address:", err)
// 		return
// 	}

// 	fmt.Println("Roku IP address:", rokuIP)
// }

// func findRokuIP(rokuMAC string, timeout time.Duration) (string, error) {
// 	// Retrieve local network interfaces
// 	interfaces, err := net.Interfaces()
// 	if err != nil {
// 		return "", err
// 	}

// 	// Iterate over interfaces
// 	for _, iface := range interfaces {
// 		// Skip loopback and inactive interfaces
// 		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
// 			continue
// 		}

// 		// Retrieve interface addresses
// 		addrs, err := iface.Addrs()
// 		if err != nil {
// 			return "", err
// 		}

// 		// Iterate over addresses
// 		for _, addr := range addrs {
// 			ipNet, ok := addr.(*net.IPNet)
// 			if !ok || ipNet.IP.To4() == nil {
// 				continue
// 			}

// 			// Perform ARP scan on the local network
// 			ipRange := ipNet.IP.Mask(ipNet.Mask)
// 			ipRange[3] = 0 // Adjust based on network configuration

// 			parsedMAC, err := net.ParseMAC(rokuMAC)
// 			if err != nil {
// 				return "", err
// 			}

// 			ip, err := arpScan(ipRange.String(), iface.Name, parsedMAC, timeout)
// 			if err != nil {
// 				return "", err
// 			}

// 			if ip != nil {
// 				return ip.String(), nil
// 			}
// 		}
// 	}

// 	return "", fmt.Errorf("Roku device not found")
// }

// func arpScan(network string, ifaceName string, targetMAC net.HardwareAddr, timeout time.Duration) (net.IP, error) {
// 	iface, err := net.InterfaceByName(ifaceName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Open a raw socket to send ARP packets
// 	socket, err := net.ListenPacket("arp", iface.Name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer socket.Close()

// 	// Send ARP request packets and wait for responses
// 	startTime := time.Now()
// 	targetIP := targetMAC.IP()
// 	request, err := newARPRequest(iface.HardwareAddr, iface.IP, targetMAC, targetIP)
// 	if err != nil {
// 		return nil, err
// 	}
// 	socket.WriteTo(request, &net.IPAddr{IP: targetIP})
// 	socket.SetDeadline(startTime.Add(timeout))

// 	// Receive ARP response packets
// 	buffer := make([]byte, 42)
// 	for {
// 		_, addr, err := socket.ReadFrom(buffer)
// 		if err != nil {
// 			if strings.Contains(err.Error(), "i/o timeout") {
// 				return nil, fmt.Errorf("ARP scan timed out")
// 			}
// 			return nil, err
// 		}

// 		packet, err := parseARPPacket(buffer)
// 		if err != nil {
// 			continue
// 		}

// 		if packet.Operation == arpOpReply && bytesEqual(packet.TargetMAC, iface.HardwareAddr) &&
// 			bytesEqual(packet.TargetIP, iface.IP) && bytesEqual(packet.SenderMAC, targetMAC) {
// 			return packet.SenderIP, nil
// 		}
// 	}
// }

// func newARPRequest(senderMAC net.HardwareAddr, senderIP net.IP, targetMAC net.HardwareAddr, targetIP net.IP) ([]byte, error) {
// 	hwType := []byte{0x00, 0x01}    // Ethernet
// 	protoType := []byte{0x08, 0x00} // IPv4
// 	hwSize := []byte{0x06}          // MAC address size
// 	protoSize := []byte{0x04}       // IP address size
// 	opcode := []byte{0x00, 0x01}    // ARP request
// 	senderMACBytes := senderMAC
// 	targetMACBytes := targetMAC
// 	packet := append(hwType, protoType...)
// 	packet = append(packet, hwSize...)
// 	packet = append(packet, protoSize...)
// 	packet = append(packet, opcode...)
// 	packet = append(packet, senderMACBytes...)
// 	packet = append(packet, senderIP.To4()...)
// 	packet = append(packet, targetMACBytes...)
// 	packet = append(packet, targetIP.To4()...)
// 	return packet, nil
// }

// func parseARPPacket(packet []byte) (*arpPacket, error) {
// 	if len(packet) < 42 || packet[7] != 0x02 || packet[20] != 0x00 || packet[21] != 0x01 {
// 		return nil, fmt.Errorf("Invalid ARP packet")
// 	}
// 	return &arpPacket{
// 		HardwareType: packet[0:2],
// 		ProtocolType: packet[2:4],
// 		HardwareSize: packet[4:5],
// 		ProtocolSize: packet[5:6],
// 		Operation:    packet[6:8],
// 		SenderMAC:    net.HardwareAddr(packet[8:14]),
// 		SenderIP:     net.IP(packet[14:18]),
// 		TargetMAC:    net.HardwareAddr(packet[18:24]),
// 		TargetIP:     net.IP(packet[24:28]),
// 	}, nil
// }

// func bytesEqual(a, b []byte) bool {
// 	return len(a) == len(b) && bytesEqualPartial(a, b, len(a))
// }

// func bytesEqualPartial(a, b []byte, length int) bool {
// 	for i := 0; i < length; i++ {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// type arpPacket struct {
// 	HardwareType []byte
// 	ProtocolType []byte
// 	HardwareSize []byte
// 	ProtocolSize []byte
// 	Operation    []byte
// 	SenderMAC    net.HardwareAddr
// 	SenderIP     net.IP
// 	TargetMAC    net.HardwareAddr
// 	TargetIP     net.IP
// }
