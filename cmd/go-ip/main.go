package main

import (
	"github.com/hra42/Go-IP/internal/network"
	"log"
)

func main() {
	newNetwork := network.Network{
		IpAddress:  network.InputIP(),
		SubnetMask: network.InputSubnetMask(),
	}

	log.Printf("Die Netzwerk Adresse ist: %v\n", newNetwork.GetNetworkAddress())
	log.Printf("Die erste nutzbare Adresse des Netzwerks ist: %v\n", newNetwork.GetFirstIP())
	log.Printf("Die letzte nutzbare Adresse des Netzwerks ist: %v\n", newNetwork.GetLastIP())
	log.Printf("Die Broadcast Adresse des Netzwerks ist: %v\n", newNetwork.GetBroadcastAddress())
}
