package happy_number

import (
	"testing"
)

func TestHappyNumbersUpToHundred(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{1, true},
		{7, true},
		{10, true},
		{13, true},
		{19, true},
		{23, true},
		{28, true},
		{31, true},
		{32, true},
		{44, true},
		{49, true},
		{68, true},
		{70, true},
		{79, true},
		{82, true},
		{86, true},
		{91, true},
		{94, true},
		{97, true},
		{100, true},
	}

	algorithms := []AlgorithmType{
		RECURSIVE,
		UNROLLED,
	}

	for _, algo := range algorithms {
		for _, tt := range tests {
			mood := "happy"
			if tt.out == false {
				mood = "sad"
			}
			isHappy := HappyNumber(tt.in, algo)
			if isHappy != tt.out {
				t.Errorf("Expected number %v to be %s", tt.in, mood)
			}
		}
	}
}

func TestSadNumbersSelected(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{0, false},
		{2, false},
		{3, false},
		{4, false},
		{5, false},
		{6, false},
		{8, false},
		{9, false},
	}

	algorithms := []AlgorithmType{
		RECURSIVE,
		UNROLLED,
	}

	for _, algo := range algorithms {
		for _, tt := range tests {
			mood := "happy"
			if tt.out == false {
				mood = "sad"
			}
			isHappy := HappyNumber(tt.in, algo)
			if isHappy != tt.out {
				t.Errorf("Expected number %v to be %s", tt.in, mood)
			}
		}
	}
}
