package main

import (
	"flag"
	desktop_app "github.com/hra42/Go-IP/internal/desktop-app"
	"github.com/hra42/Go-IP/internal/network"
	"github.com/hra42/Go-IP/internal/numberConversion"
	"github.com/hra42/Go-IP/internal/quiz"
)

func main() {
	IPFlag := flag.Bool(
		"ip",
		false,
		"Gibt die Netzwerk Adresse, Broadcast Adresse, erste und letzte IP Adresse aus.",
	)
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
	case *IPFlag:
		newNetwork := network.Network{
			IpAddress:  network.InputIP(),
			SubnetMask: network.InputSubnetMask(),
		}
		newNetwork.Print()
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
		desktop_app.TestDesktop()
	}

}
