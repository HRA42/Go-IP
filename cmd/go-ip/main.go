package main

import (
	"flag"
	"github.com/hra42/Go-IP/internal/network"
)

func main() {
	cidrFlag := flag.Bool(
		"cidr",
		false,
		"Gibt die Subnetz Maske zu einer Netzwerk Adresse in CIDR Notation aus.",
	)
	flag.Parse()

	switch true {
	case *cidrFlag:
		networkAddress := network.InputNetworkAddress()
		newNetwork := network.Network{
			IpAddress:  networkAddress.IP,
			SubnetMask: networkAddress.Mask,
		}
		newNetwork.PrintCIDR()
	default:
		newNetwork := network.Network{
			IpAddress:  network.InputIP(),
			SubnetMask: network.InputSubnetMask(),
		}
		newNetwork.Print()
	}

}
