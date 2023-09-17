package main

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	firstNumber, secondNumber, operation, err := parseInput("")
	if err == nil {
		t.Errorf("Validate input must return error if input is empty")
	}
	if firstNumber != 0 {
		t.Errorf("First number must be 0 if input is empty")
	}
	if secondNumber != 0 {
		t.Errorf("Second number must be 0 if input is empty")
	}
	if operation != "" {
		t.Errorf("Operation must be empty string if input is empty")
	}

	firstNumber, secondNumber, operation, err = parseInput("text + text")
	if err == nil {
		t.Errorf("Validate input must return error if input has wrong format")
	}
	if firstNumber != 0 {
		t.Errorf("First number must be 0 if input has wrong format")
	}
	if secondNumber != 0 {
		t.Errorf("Second number must be 0 if input has wrong format")
	}
	if operation != "" {
		t.Errorf("Operation must be empty string if input has wrong format")
	}

	firstNumber, secondNumber, operation, err = parseInput("2 + 2")
	if err != nil {
		t.Errorf("Validate input must return nil error if input is 2 + 2")
	}
	if firstNumber != 2 {
		t.Errorf("First number must be 0 if input is 2 + 2")
	}
	if secondNumber != 2 {
		t.Errorf("Second number must be 0 if input is 2 + 2")
	}
	if operation != "+" {
		t.Errorf("Operation must be empty string if input is 2 + 2")
	}
}

func TestAddition(t *testing.T) {
	result, err := calculate(4, 2, "+")
	if err != nil {
		t.Errorf("4+2 must not return error")
	}
	if result != "6" {
		t.Errorf("The result of 4+2 must be 6")
	}
}

func TestSubstraction(t *testing.T) {
	result, err := calculate(5, 3, "-")
	if err != nil {
		t.Errorf("5-3 must not return error")
	}
	if result != "2" {
		t.Errorf("The result of 5-3 must be 2")
	}
}

func TestMultiplication(t *testing.T) {
	result, err := calculate(9, 8, "*")
	if err != nil {
		t.Errorf("9*8 must not return error")
	}
	if result != "72" {
		t.Errorf("The result of 9*8 must be 72")
	}
}

func TestDivisiton(t *testing.T) {
	result, err := calculate(9, 2, "/")
	if err != nil {
		t.Errorf("9/2 must not return error")
	}
	if result != "4.5" {
		t.Errorf("The result of 9/2 must be 4.5")
	}

	result, err = calculate(8, 0, "/")
	if err == nil {
		t.Errorf("Division by zero must return error")
	}
	if result != "" {
		t.Errorf("The result of 9/0 must be empty string")
	}
}

func TestPower(t *testing.T) {
	result, err := calculate(2, 6, "^")
	if err != nil {
		t.Errorf("2^6 must not return error")
	}
	if result != "64" {
		t.Errorf("The result of 2^6 must be 64")
	}
}
