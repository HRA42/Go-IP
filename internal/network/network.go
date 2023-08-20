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
	if broadcastAddress == nil {
		return nil
	}
	lastUsableAddress := make(net.IP, len(broadcastAddress))
	copy(lastUsableAddress, broadcastAddress)
	// If it is /31 CIDR
	// In CIDR, for /31, broadcast and network address are valid hosts
	cidr, _ := n.SubnetMask.Size()
	if cidr == 31 {
		return lastUsableAddress
	}
	// If it is /32 CIDR
	// In CIDR, for /32, the host itself is its own network and broadcast address
	if cidr == 32 {
		return n.IpAddress
	}
	// For CIDR < /31
	if len(lastUsableAddress) > 0 {
		lastUsableAddress[len(lastUsableAddress)-1]--
	}
	return lastUsableAddress
}

func (n Network) GetFirstIP() net.IP {
	networkAddress := n.GetNetworkAddress().To4() // n.GetNetworkAddress() is assumed to be defined elsewhere
	if networkAddress == nil {
		return nil
	}
	firstUsableAddress := make(net.IP, len(networkAddress))
	copy(firstUsableAddress, networkAddress)
	cidr, _ := n.SubnetMask.Size()
	// If it is /31 CIDR notation
	if cidr == 31 {
		firstUsableAddress[len(firstUsableAddress)-1]++
		return firstUsableAddress
	}
	// If it is /32 CIDR notation
	if cidr == 32 {
		return n.IpAddress
	}
	// For CIDR < /31
	if len(firstUsableAddress) > 0 {
		firstUsableAddress[len(firstUsableAddress)-1]++
	}
	return firstUsableAddress
}

func (n Network) GetCIDRFromSubnetMask() string {
	cidr, _ := n.SubnetMask.Size()
	return fmt.Sprintf("%v/%v", n.GetNetworkAddress(), cidr)
}

func InputNetworkAddress() net.IPNet {
	var NetworkAddress string
	fmt.Print("Bitte gib die Netzwerk Adresse ein: ")
	_, err := fmt.Scanln(&NetworkAddress)
	if err != nil {
		log.Fatal("Fehler bei der Eingabe:", err)
	}
	_, netAddress, err := net.ParseCIDR(NetworkAddress)
	if netAddress == nil {
		log.Fatal("IP Address nicht valide!")
	}
	fmt.Println("Du hast folgende Netzwerk Adresse eingegeben:", netAddress)
	return *netAddress
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

func (n Network) PrintCIDR() {
	log.Printf("Die Subnetzmaske ist: %d.%d.%d.%d\n",
		n.SubnetMask[0],
		n.SubnetMask[1],
		n.SubnetMask[2],
		n.SubnetMask[3],
	)
	log.Printf("Die Netzwerk Adresse ist: %v\n",
		n.GetNetworkAddress(),
	)
	log.Printf("Die erste nutzbare Adresse des Netzwerks ist: %v\n",
		n.GetFirstIP(),
	)
	log.Printf("Die letzte nutzbare Adresse des Netzwerks ist: %v\n",
		n.GetLastIP(),
	)
	log.Printf("Die Broadcast Adresse des Netzwerks ist: %v\n",
		n.GetBroadcastAddress(),
	)
}

func (n Network) Print() {
	log.Printf("Die CIDR Notation des Netzwerks ist: %v",
		n.GetCIDRFromSubnetMask(),
	)
	log.Printf("Die Netzwerk Adresse ist: %v\n",
		n.GetNetworkAddress(),
	)
	log.Printf("Die erste nutzbare Adresse des Netzwerks ist: %v\n",
		n.GetFirstIP(),
	)
	log.Printf("Die letzte nutzbare Adresse des Netzwerks ist: %v\n",
		n.GetLastIP(),
	)
	log.Printf("Die Broadcast Adresse des Netzwerks ist: %v\n",
		n.GetBroadcastAddress(),
	)
}
