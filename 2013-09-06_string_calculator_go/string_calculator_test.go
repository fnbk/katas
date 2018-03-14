package string_calculator

import (
	"testing"
)

func Test_AddReturnsZeroForEmptyString(t *testing.T) {
	result := Add("")
	if result != 0 {
		t.Errorf("Add(\"\") did not return 0. Actual: %v", result)
	}
}

func Test_AddReturnsTheSumOfOneNumber(t *testing.T) {
	result := Add("1")
	if result != 1 {
		t.Errorf("Add('1') did not return 1. Actual: %v", result)
	}
}

func Test_AddReturnsTheSumOfTwoNumbers(t *testing.T) {
	result := Add("1,2")
	if result != 3 {
		t.Errorf("Add('1,2') did not return 3. Actual: %v", result)
	}
}

func Test_AddReturnsTheSumOfUnknownNumbers(t *testing.T) {
	result := Add("1,2,3")
	if result != 6 {
		t.Errorf("Add(%v) did not return %v. Actual: %v", "1,2,3", 6, result)
	}
}

func Test_AddReturnsTheSumWithNewLineAsDelimiter(t *testing.T) {
	result := Add("1\n2,3")
	if result != 6 {
		t.Errorf("Add(%v) did not return %v. Actual: %v", "1,2,3", 6, result)
	}
}

func Test_AddReturnsTheSumWithFirstLineSpecifyingDefaultDelimiter(t *testing.T) {
	result := Add("//;\n1;2;3")
	if result != 6 {
		t.Errorf("Add(%v) did not return %v. Actual: %v", "1;2;3", 6, result)
	}
}
