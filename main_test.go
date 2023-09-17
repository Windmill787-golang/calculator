package main

import "testing"

func TestEmpty(t *testing.T) {
	result, err := validateInput("")
	if err == nil {
		t.Errorf("Validate input must return error if input is empty")
	}
	if result != "" {
		t.Errorf("Result must be empty string if input is empty")
	}
}

func TestAddition(t *testing.T) {
	result, err := calculate("4+2")
	if err != nil {
		t.Errorf("4+2 must not return error")
	}
	if result != "6" {
		t.Errorf("The result of 4+2 must be 6")
	}
}

func TestSubstraction(t *testing.T) {
	result, err := calculate("5-3")
	if err != nil {
		t.Errorf("5-3 must not return error")
	}
	if result != "2" {
		t.Errorf("The result of 5-3 must be 2")
	}
}

func TestMultiplication(t *testing.T) {
	result, err := calculate("9*8")
	if err != nil {
		t.Errorf("9*8 must not return error")
	}
	if result != "72" {
		t.Errorf("The result of 9*8 must be 72")
	}
}

func TestDivisiton(t *testing.T) {
	result, err := calculate("9/2")
	if err != nil {
		t.Errorf("9/2 must not return error")
	}
	if result != "4.5" {
		t.Errorf("The result of 9/2 must be 4.5")
	}

	result, err = calculate("8/0")
	if err == nil {
		t.Errorf("Division by zero must return error")
	}
	if result != "" {
		t.Errorf("The result of 9/0 must be empty string")
	}
}

func TestPower(t *testing.T) {
	result, err := calculate("2^6")
	if err != nil {
		t.Errorf("2^6 must not return error")
	}
	if result != "64" {
		t.Errorf("The result of 2^6 must be 64")
	}
}
