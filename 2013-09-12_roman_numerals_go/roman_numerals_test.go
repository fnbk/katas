package roman_numerals

import (
	"testing"
)

func verify(t *testing.T, in int, out string, solution string) {
	if solution != out {
		t.Error(
			"For", in,
			"expected", out,
			"got", solution,
		)
	}
}

func TestConvertSingleDigitsToRoman(t *testing.T) {
	tests := []struct{
		in int
		out string
	}{
		{in: 1, out: "I"},
		{in: 2, out: "II"},
		{in: 3, out: "III"},
		{in: 4, out: "IV"},
		{in: 5, out: "V"},
		{in: 6, out: "VI"},
		{in: 7, out: "VII"},
		{in: 8, out: "VIII"},
		{in: 9, out: "IX"},
	}

	for _, tt := range tests {
		solution := Convert(tt.in)

		verify(t, tt.in, tt.out, solution)
	}
}

func TestConvertTwoDigitsToRoman(t *testing.T) {
	tests := []struct{
		in int
		out string
	}{
		{in: 10, out: "X"},
		{in: 11, out: "XI"},
		{in: 12, out: "XII"},
		{in: 13, out: "XIII"},
		{in: 14, out: "XIV"},
		{in: 15, out: "XV"},
		{in: 16, out: "XVI"},
		{in: 17, out: "XVII"},
		{in: 18, out: "XVIII"},
		{in: 19, out: "XIX"},
	}

	for _, tt := range tests {
		solution := Convert(tt.in)

		verify(t, tt.in, tt.out, solution)
	}
}

func TestConvertSecondDigitToRoman(t *testing.T) {
	tests := []struct{
		in int
		out string
	}{
		{in: 10, out: "X"},
		{in: 20, out: "XX"},
		{in: 30, out: "XXX"},
		{in: 40, out: "XL"},
		{in: 50, out: "L"},
		{in: 60, out: "LX"},
		{in: 70, out: "LXX"},
		{in: 80, out: "LXXX"},
		{in: 90, out: "XC"},
		{in: 99, out: "XCIX"},
	}

	for _, tt := range tests {
		solution := Convert(tt.in)

		verify(t, tt.in, tt.out, solution)
	}
}
