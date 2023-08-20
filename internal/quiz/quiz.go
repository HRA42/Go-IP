package quiz

import (
	"fmt"
	"github.com/hra42/Go-IP/internal/numberConversion"
	"strconv"
)

func createQuiz() (quizMap map[string]string) {
	quizMap = map[string]string{
		"Rechne die Zahl 10 in das Binär System um!":       numberConversion.DecimalToBinary(10),
		"Rechne die Zahl 10 in das Hexadezimal System um!": numberConversion.DecimalToHexadecimal(10),
		"Rechne die Zahl 11110000 in das Dezimal System um!": strconv.FormatInt(
			int64(numberConversion.BinaryToDecimal("11110000")), 10),
		"Rechne die Zahl 0011110000 in das Hexadezimal System um!": numberConversion.BinaryToHexadecimal(
			"11110000"),
		"Rechne die Zahl 0x10F45 in das Dezimal System um!": strconv.FormatInt(
			numberConversion.HexadecimalToDecimal("10F45"), 10),
		"Rechne die Zahl 0x0010F45 in das Binär System um!": numberConversion.HexadecimalToBinary("0010F45"),
	}
	return
}

func checkAnswer(solution string, answer string) bool {
	if solution == answer {
		fmt.Println("Richtig!")
		return true
	} else {
		fmt.Println("Falsch!")
		fmt.Printf("Die richtige Antwort wäre: %v\n", solution)
		return false
	}
}

func readAnswer() string {
	var answer string
	_, err := fmt.Scan(&answer)
	if err != nil {
		fmt.Println(err)
	}
	return answer
}

// Start the main function for quiz
func Start() {
	quiz := createQuiz()
	var score int
	for i := range quiz {
		fmt.Println(i)
		answer := readAnswer()
		solution := quiz[i]
		if checkAnswer(solution, answer) {
			score++
		}
	}
	fmt.Printf("Du hast %v von %v Fragen richtig beantwortet!\n", score, len(quiz))
	if float32(score/len(quiz)) >= 0.5 {
		fmt.Println("Glückwunsch du hast das Quiz bestanden!!")
	} else {
		fmt.Println("Schade, leider bist du durchgefallen!")
	}
}
