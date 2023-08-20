package quiz

import (
	"fmt"
	"github.com/hra42/Go-IP/internal/numberConversion"
	"math/rand"
	"strconv"
	"time"
)

func createQuiz() (quizMap map[string]string) {
	quizMap = map[string]string{
		"Rechne die Zahl 10 in das Binär System um!":               numberConversion.DecimalToBinary(10),
		"Rechne die Zahl 10 in das Hexadezimal System um!":         numberConversion.DecimalToHexadecimal(10),
		"Rechne die Zahl 11110000 in das Dezimal System um!":       strconv.FormatInt(int64(numberConversion.BinaryToDecimal("11110000")), 10),
		"Rechne die Zahl 0011110000 in das Hexadezimal System um!": numberConversion.BinaryToHexadecimal("11110000"),
		"Rechne die Zahl 0x10F45 in das Dezimal System um!":        strconv.FormatInt(numberConversion.HexadecimalToDecimal("10F45"), 10),
		"Rechne die Zahl 0x0010F45 in das Binär System um!":        numberConversion.HexadecimalToBinary("0010F45"),
		"Rechne die Zahl 25 in das Binär System um!":               numberConversion.DecimalToBinary(25),
		"Rechne die Zahl 45 in das Binär System um!":               numberConversion.DecimalToBinary(45),
		"Rechne die Zahl 15 in das Hexadezimal System um!":         numberConversion.DecimalToHexadecimal(15),
		"Rechne die Zahl 75 in das Hexadezimal System um!":         numberConversion.DecimalToHexadecimal(75),
		"Rechne die Zahl 10101010 in das Dezimal System um!":       strconv.FormatInt(int64(numberConversion.BinaryToDecimal("10101010")), 10),
		"Rechne die Zahl 11001100 in das Dezimal System um!":       strconv.FormatInt(int64(numberConversion.BinaryToDecimal("11001100")), 10),
		"Rechne die Zahl 10101111 in das Hexadezimal System um!":   numberConversion.BinaryToHexadecimal("10101111"),
		"Rechne die Zahl 11010101 in das Hexadezimal System um!":   numberConversion.BinaryToHexadecimal("11010101"),
		"Rechne die Zahl 0xA23B in das Dezimal System um!":         strconv.FormatInt(numberConversion.HexadecimalToDecimal("A23B"), 10),
		"Rechne die Zahl 0x5C7F in das Dezimal System um!":         strconv.FormatInt(numberConversion.HexadecimalToDecimal("5C7F"), 10),
		"Rechne die Zahl 0x12A5 in das Binär System um!":           numberConversion.HexadecimalToBinary("12A5"),
		"Rechne die Zahl 0x3BC9 in das Binär System um!":           numberConversion.HexadecimalToBinary("3BC9"),
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

func shuffle(q map[string]string) map[string]string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	keys := make([]string, 0, len(q))
	for k := range q {
		keys = append(keys, k)
	}
	r.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	shuffledMap := make(map[string]string)
	for _, key := range keys {
		shuffledMap[key] = q[key]
	}

	return shuffledMap
}

// Start the main function for quiz
func Start() {
	quiz := createQuiz()
	quiz = shuffle(quiz)
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
