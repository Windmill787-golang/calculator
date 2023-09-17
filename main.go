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

func calculate(input string) (string, error) {
	// fmt.Printf("Trimmed: '%s'\n", input)

	pattern := `([,.0-9]+|[\+\-\*\/\^])`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllString(input, -1)

	if len(matches) != 3 {
		return "", fmt.Errorf("wrong format")
	}
	// fmt.Println(matches)
	// fmt.Println(len(matches))
	firstNumber, err := strconv.ParseFloat(matches[0], 64)
	if err != nil {
		return "", err
	}
	secondNumber, err := strconv.ParseFloat(matches[2], 64)
	if err != nil {
		return "", err
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

func validateInput(input string) (string, error) {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) == 0 {
		return "", fmt.Errorf("input is empty")
	}

	return input, nil
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
			// fmt.Printf("You entered: '%s'\n", input)
		} else {
			fmt.Println("Error:", scanner.Err())
			return
		}

		input, err := validateInput(input)
		if err != nil {
			fmt.Println("Input error:", err.Error())
			continue
		}

		result, err := calculate(input)
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		fmt.Printf("Result: %s\n", result)
	}
}
