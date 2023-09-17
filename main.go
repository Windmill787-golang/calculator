package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	for {
		var input string
		scanner := bufio.NewScanner(os.Stdin)

		// Prompt the user for input
		fmt.Print("Enter a calculation (e.g. 2+2): ")

		// Use scanner to read the entire line of text, including spaces
		if scanner.Scan() {
			input = scanner.Text()
			// fmt.Printf("You entered: '%s'\n", input)
		} else {
			fmt.Println("Error:", scanner.Err())
			return
		}

		input = strings.ReplaceAll(input, " ", "")
		// fmt.Printf("Trimmed: '%s'\n", input)

		pattern := `([,.0-9]+|[\+\-\*\/\^])`
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllString(input, -1)

		if len(matches) != 3 {
			fmt.Println("Wrong format")
			continue
		}
		// fmt.Println(matches)
		// fmt.Println(len(matches))
		firstNumber, err := strconv.ParseFloat(matches[0], 64)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}
		secondNumber, err := strconv.ParseFloat(matches[2], 64)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}

		// fmt.Println(firstNumber, secondNumber)

		var result float64
		switch matches[1] {
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber - secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			if secondNumber == 0 {
				fmt.Println("Cannot divide 0")
				continue
			}
			result = firstNumber / secondNumber
		case "^":
			result = math.Pow(firstNumber, secondNumber)
		}

		result = roundFloat(result, 9)

		if result == float64(int(result)) {
			fmt.Printf("Result: %.0f\n", result)
		} else {
			fmt.Printf("Result: %s\n", strconv.FormatFloat(result, 'f', -1, 64))
		}

		// for index, value := range matches {
		// 	fmt.Println(index, value)
		// }
	}
}
