package main

import (
	"flag"
	"github.com/hra42/Go-IP/internal/network"
	"github.com/hra42/Go-IP/internal/numberConversion"
	"github.com/hra42/Go-IP/internal/quiz"
)

func main() {
	cidrFlag := flag.Bool(
		"cidr",
		false,
		"Gibt die Subnetz Maske zu einer Netzwerk Adresse in CIDR Notation aus.",
	)
	DecimalConversionFlag := flag.Bool(
		"decimal",
		false,
		"Wandle eine Zahl vom Dezimal System in das Binär und Hexadezimal System um",
	)
	BinaryConversionFlag := flag.Bool(
		"binary",
		false,
		"Wandle eine Zahl vom Binär System in das Dezimal und Hexadezimal System um",
	)
	HexadecimalConversionFlag := flag.Bool(
		"hexadecimal",
		false,
		"Wandle eine Zahl vom Hexadezimal System in das Dezimal und Binär System um",
	)
	QuizFlag := flag.Bool(
		"quiz",
		false,
		"Startet ein Quiz",
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
	case *DecimalConversionFlag:
		numberConversion.PrintNumber(numberConversion.InputNumber())
	case *BinaryConversionFlag:
		numberConversion.PrintBinary(numberConversion.InputNoneDecimal())
	case *HexadecimalConversionFlag:
		numberConversion.PrintHexadecimal(numberConversion.InputNoneDecimal())
	case *QuizFlag:
		quiz.Start()
	default:
		newNetwork := network.Network{
			IpAddress:  network.InputIP(),
			SubnetMask: network.InputSubnetMask(),
		}
		newNetwork.Print()
	}

}
