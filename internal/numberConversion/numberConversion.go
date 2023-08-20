package numberConversion

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func InputNumber() int {
	var number int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Println(err)
	}
	return number
}

func InputNoneDecimal() string {
	var hexadecimal string
	fmt.Print("Enter a hexadecimal number: ")
	_, err := fmt.Scan(&hexadecimal)
	if err != nil {
		log.Println(err)
	}
	return hexadecimal
}

func DecimalToBinary(number int) string {
	return fmt.Sprintf("%b", number)
}

func BinaryToDecimal(binary string) (decimal int) {
	length := len(binary)
	for i := length - 1; i >= 0; i-- {
		// binary[i] is the i'th rune in binary; string(binary[i]) converts it to string
		bit, err := strconv.Atoi(string(binary[i]))
		if err != nil {
			log.Fatal(err)
		}
		decimal += bit * int(math.Pow(2, float64(length-i-1)))
	}
	return decimal
}

func DecimalToHexadecimal(number int) string {
	return fmt.Sprintf("%X", number)
}

func HexadecimalToDecimal(hexadecimal string) int64 {
	decimal, err := strconv.ParseInt(hexadecimal, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return decimal
}

func BinaryToHexadecimal(binary string) string {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%X", decimal)
}

func HexadecimalToBinary(hexadecimal string) string {
	decimal, err := strconv.ParseInt(hexadecimal, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%b", decimal)
}

func PrintNumber(number int) {
	fmt.Println("Decimal:", number)
	fmt.Println("Binary:", DecimalToBinary(number))
	fmt.Println("Hexadecimal:", DecimalToHexadecimal(number))
}

func PrintBinary(number string) {
	fmt.Println("Decimal:", BinaryToDecimal(number))
	fmt.Println("Binary:", number)
	fmt.Println("Hexadecimal:", BinaryToHexadecimal(number))
}

func PrintHexadecimal(number string) {
	fmt.Println("Decimal:", HexadecimalToDecimal(number))
	fmt.Println("Binary:", HexadecimalToBinary(number))
	fmt.Println("Hexadecimal:", number)
}
