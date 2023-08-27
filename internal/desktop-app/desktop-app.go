package desktop_app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hra42/Go-IP/internal/network"
	"github.com/hra42/Go-IP/internal/numberConversion"
	"log"
	"net"
	"strconv"
)

func TestDesktop() {
	// set up the app
	a := app.New()
	w := a.NewWindow("Go-IP")

	// set the icon
	resourceIconPng, err := fyne.LoadResourceFromURLString("https://gpt-files.postrausch.tech/icon.png")
	if err != nil {
		log.Fatal(err)
	}
	w.SetIcon(resourceIconPng)

	// set the window size
	w.Resize(fyne.NewSize(250, 300))

	// main menu
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("Menu",
			fyne.NewMenuItem("Convert IP", func() {
				convertIP(w)
			}),
			fyne.NewMenuItem("Convert CIDR", func() {
				convertCIDR(w)
			}),
			fyne.NewMenuItem("Convert Decimal", func() {
				convertDecimal(w)
			}),
			fyne.NewMenuItem("Convert Binary", func() {
				convertBinary(w)
			}),
			fyne.NewMenuItem("Convert Hexadecimal", func() {
				convertHexadecimal(w)
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
	)

	w.SetMainMenu(mainMenu)

	w.ShowAndRun()
}

func convertIP(w fyne.Window) {
	// set up the content
	inputIP := widget.NewEntry()
	inputIP.SetPlaceHolder("Enter a ip address")
	inputMask := widget.NewEntry()
	inputMask.SetPlaceHolder("Enter a subnet mask")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()
	networkAddress := network.Network{
		IpAddress:  nil,
		SubnetMask: nil,
	}
	w.SetContent(container.NewVBox(
		inputIP,
		inputMask,
		placeholder,
		widget.NewButton("Submit", func() {
			networkAddress.IpAddress = net.ParseIP(inputIP.Text)
			networkAddress.SubnetMask = net.IPMask(net.ParseIP(inputMask.Text).To4())
			placeholder.SetText(
				"First IP: " + networkAddress.GetFirstIP().String() + "\n" +
					"Last IP: " + networkAddress.GetLastIP().String() + "\n" +
					"Network Address: " + networkAddress.GetNetworkAddress().String() + "\n" +
					"Broadcast Address: " + networkAddress.GetBroadcastAddress().String(),
			)
			placeholder.Show()
		}),
	))
}

func convertCIDR(w fyne.Window) {
	// set up the content
	inputIP := widget.NewEntry()
	inputIP.SetPlaceHolder("Enter a network address in CIDR notation")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()
	networkAddress := network.Network{
		IpAddress:  nil,
		SubnetMask: nil,
	}
	w.SetContent(container.NewVBox(
		inputIP,
		placeholder,
		widget.NewButton("Submit", func() {
			IP, CIDR, err := net.ParseCIDR(inputIP.Text)
			if err != nil {
				log.Fatal(err)
			}
			networkAddress.IpAddress = IP
			networkAddress.SubnetMask = CIDR.Mask
			placeholder.SetText(
				"First IP: " + networkAddress.GetFirstIP().String() + "\n" +
					"Last IP: " + networkAddress.GetLastIP().String() + "\n" +
					"Network Address: " + networkAddress.GetNetworkAddress().String() + "\n" +
					"Broadcast Address: " + networkAddress.GetBroadcastAddress().String(),
			)
			placeholder.Show()
		}),
	))
}

func convertDecimal(w fyne.Window) {
	// set up the content
	inputDecimal := widget.NewEntry()
	inputDecimal.SetPlaceHolder("Enter a decimal number")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()
	w.SetContent(container.NewVBox(
		inputDecimal,
		placeholder,
		widget.NewButton("Submit", func() {
			decimal, err := strconv.Atoi(inputDecimal.Text)
			if err != nil {
				log.Fatal(err)
			}
			placeholder.SetText(
				"Binary: " + numberConversion.DecimalToBinary(decimal) + "\n" +
					"Hexadecimal: " + numberConversion.DecimalToHexadecimal(decimal),
			)
			placeholder.Show()
		}),
	))
}

func convertBinary(w fyne.Window) {
	// set up the content
	inputBinary := widget.NewEntry()
	inputBinary.SetPlaceHolder("Enter a binary number")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()
	w.SetContent(container.NewVBox(
		inputBinary,
		placeholder,
		widget.NewButton("Submit", func() {
			placeholder.SetText(
				"Decimal: " + strconv.Itoa(numberConversion.BinaryToDecimal(inputBinary.Text)) + "\n" +
					"Hexadecimal: " + numberConversion.BinaryToHexadecimal(inputBinary.Text),
			)
			placeholder.Show()
		}),
	))
}

func convertHexadecimal(w fyne.Window) {
	// set up the content
	inputHexadecimal := widget.NewEntry()
	inputHexadecimal.SetPlaceHolder("Enter a hexadecimal number")
	placeholder := widget.NewLabel("This is a placeholder")
	placeholder.Hide()
	w.SetContent(container.NewVBox(
		inputHexadecimal,
		placeholder,
		widget.NewButton("Submit", func() {
			placeholder.SetText(
				"Decimal: " + strconv.FormatInt(numberConversion.HexadecimalToDecimal(inputHexadecimal.Text),
					10) + "\n" +
					"Binary: " + numberConversion.HexadecimalToBinary(inputHexadecimal.Text),
			)
			placeholder.Show()
		}),
	))
}
