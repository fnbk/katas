package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerateStartWithZero(t *testing.T) {
	numbers := Generate()
	if numbers[0] != "0" {
		t.Errorf("first number should be 0, got: %v", numbers[0])
	}
}

func TestEveryThirdElementContainsFizz(t *testing.T) {
	numbers := Generate()
	for idx, _ := range numbers {
		if idx == 0 {
			continue
		}
		if idx%3 == 0 {
			if !strings.Contains(numbers[idx], "fizz") {
				t.Errorf("numbers[%v] does not include 'fizz': %v", idx, numbers[idx])
			}
		}
	}
}

func TestEveryFifthElementContainsBuzz(t *testing.T) {
	numbers := Generate()
	for idx, _ := range numbers {
		if idx == 0 {
			continue
		}
		if idx%5 == 0 {
			if !strings.Contains(numbers[idx], "buzz") {
				t.Errorf("numbers[%v] does not include 'buzz': %v", idx, numbers[idx])
			}
		}
	}
}

func TestEveryNonThirdAndFifthElementIsANumber(t *testing.T) {
	numbers := Generate()
	for idx, _ := range numbers {
		if idx == 0 {
			continue
		}
		if idx%3 != 0 && idx%5 != 0 {
			number, err := strconv.Atoi(numbers[idx])
			if err != nil {
				t.Errorf("%v: %v", err, numbers[idx])
			}
			if number != idx {
				t.Errorf("numbers[%v] is not a number: %v", idx, numbers[idx])
			}
		}
	}
}
