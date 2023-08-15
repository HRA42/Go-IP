package network

import (
	"fmt"
	"log"
	"net"
)

type Network struct {
	IpAddress  net.IP
	SubnetMask net.IPMask
}

func (n Network) GetBroadcastAddress() net.IP {
	ip := n.IpAddress.To4()
	mask := n.SubnetMask

	broadcast := make(net.IP, len(ip))
	for i := range ip {
		broadcast[i] = ip[i] | ^mask[i]
	}

	return broadcast
}

func (n Network) GetNetworkAddress() net.IP {
	return n.IpAddress.Mask(n.SubnetMask)
}

func (n Network) GetLastIP() net.IP {
	broadcastAddress := n.GetBroadcastAddress().To4()
	lastUsableAddress := make(net.IP, len(broadcastAddress))
	copy(lastUsableAddress, broadcastAddress)
	lastUsableAddress[3]--
	return lastUsableAddress
}

func (n Network) GetFirstIP() net.IP {
	networkAddress := n.GetNetworkAddress().To4()
	firstUsableAddress := make(net.IP, len(networkAddress))
	copy(firstUsableAddress, networkAddress)
	firstUsableAddress[3]++
	return firstUsableAddress
}

func InputIP() net.IP {
	var ipAddress string
	fmt.Print("Bitte gib die IP Adresse ein: ")
	_, err := fmt.Scanln(&ipAddress)
	if err != nil {
		log.Fatal("Fehler bei der Eingabe:", err)

	}
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		log.Fatal("IP Address nicht valide!")
	}
	fmt.Println("Du hast folgende IP Adresse eingegeben:", ip)
	return ip
}

func InputSubnetMask() net.IPMask {
	var a, b, c, d uint8
	fmt.Print("Bitte gib die Subnetzmaske in der Form (a.b.c.d): ")
	_, err := fmt.Scanf("%d.%d.%d.%d\n", &a, &b, &c, &d)
	if err != nil {
		log.Fatal("Fehler bei der Eingabe:", err)
	}
	subnet := net.IPv4Mask(a, b, c, d)
	if IsValidSubnetMask(subnet) {
		log.Printf("Du hast folgende Subnetzmaske eingegeben: %d.%d.%d.%d\n",
			subnet[0], subnet[1], subnet[2], subnet[3])
	} else {
		log.Fatal("Die Subnetzmaske ist ungültig!")
	}

	return subnet
}

func IsValidSubnetMask(subnetMask net.IPMask) bool {
	mask := uint32(subnetMask[0])<<24 + uint32(subnetMask[1])<<16 + uint32(subnetMask[2])<<8 + uint32(subnetMask[3])

	// count leading one bit
	leadingOnes := 0
	for i := 31; i >= 0; i-- {
		if (mask>>i)&1 == 1 {
			leadingOnes++
		} else {
			break
		}
	}

	// the mask should be all ones, followed by all zeros
	expectedMask := ^uint32(0) << uint32(32-leadingOnes)
	return mask == expectedMask
}
