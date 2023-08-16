package main

import (
	"flag"
	"log"

	"github.com/hra42/Go-IP/internal/network"
)

func main() {
	cidrFlag := flag.Bool(
		"cidr",
		false,
		"Gibt die Subnetz Maske zu einer Netzwerk Adresse in CIDR Notation aus.",
	)
	flag.Parse()
	if *cidrFlag {
		networkAddress := network.InputNetworkAddress()
		newNetwork := network.Network{
			IpAddress:  networkAddress.IP,
			SubnetMask: networkAddress.Mask,
		}
		log.Printf("Die Subnetzmaske ist: %d.%d.%d.%d\n",
			newNetwork.SubnetMask[0],
			newNetwork.SubnetMask[1],
			newNetwork.SubnetMask[2],
			newNetwork.SubnetMask[3],
		)
		log.Printf("Die Netzwerk Adresse ist: %v\n",
			newNetwork.GetNetworkAddress(),
		)
		log.Printf("Die erste nutzbare Adresse des Netzwerks ist: %v\n",
			newNetwork.GetFirstIP(),
		)
		log.Printf("Die letzte nutzbare Adresse des Netzwerks ist: %v\n",
			newNetwork.GetLastIP(),
		)
		log.Printf("Die Broadcast Adresse des Netzwerks ist: %v\n",
			newNetwork.GetBroadcastAddress(),
		)
	} else {
		newNetwork := network.Network{
			IpAddress:  network.InputIP(),
			SubnetMask: network.InputSubnetMask(),
		}
		log.Printf("Die CIDR Notation des Netzwerks ist: %v",
			newNetwork.GetCIDRFromSubnetMask(),
		)
		log.Printf("Die Netzwerk Adresse ist: %v\n",
			newNetwork.GetNetworkAddress(),
		)
		log.Printf("Die erste nutzbare Adresse des Netzwerks ist: %v\n",
			newNetwork.GetFirstIP(),
		)
		log.Printf("Die letzte nutzbare Adresse des Netzwerks ist: %v\n",
			newNetwork.GetLastIP(),
		)
		log.Printf("Die Broadcast Adresse des Netzwerks ist: %v\n",
			newNetwork.GetBroadcastAddress(),
		)
	}

}
