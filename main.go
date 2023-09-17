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

var input string

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func calculate(firstNumber, secondNumber float64, operation string) (string, error) {
	var result float64
	switch operation {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "*":
		result = firstNumber * secondNumber
	case "/":
		if secondNumber == 0 {
			return "", fmt.Errorf("division by 0 is undefined")
		}
		result = firstNumber / secondNumber
	case "^":
		result = math.Pow(firstNumber, secondNumber)
	}

	result = roundFloat(result, 9)

	var res string
	if result == float64(int(result)) {
		res = fmt.Sprintf("%.0f", result)
	} else {
		res = strconv.FormatFloat(result, 'f', -1, 64)
	}
	return res, nil
}

func parseInput(input string) (float64, float64, string, error) {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, ",", ".")

	if len(input) == 0 {
		return 0, 0, "", fmt.Errorf("input is empty")
	}

	pattern := `([,.0-9]+|[\+\-\*\/\^])`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllString(input, -1)

	if len(matches) != 3 {
		return 0, 0, "", fmt.Errorf("invalid format")
	}

	firstNumber, err := strconv.ParseFloat(matches[0], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid format of first number")
	}
	secondNumber, err := strconv.ParseFloat(matches[2], 64)
	if err != nil {
		return 0, 0, "", fmt.Errorf("invalid format of second number")
	}

	return firstNumber, secondNumber, matches[1], nil
}

func main() {
	for {
		input = ""
		scanner := bufio.NewScanner(os.Stdin)

		// Prompt the user for input
		fmt.Print("Enter a calculation (e.g. 2+2): ")

		// Use scanner to read the entire line of text, including spaces
		if scanner.Scan() {
			input = scanner.Text()
		} else {
			fmt.Println("Error:", scanner.Err())
			continue
		}

		firstNumber, secondNumber, operation, err := parseInput(input)
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}

		result, err := calculate(firstNumber, secondNumber, operation)
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		fmt.Printf("Result: %s\n", result)
	}
}
