package roman_numerals

import (
	"testing"
)

func Test_NumeralForNumbersOneTwoTreeFourFiveSixSevenEightNineTen(t *testing.T) {
	tests := []struct{
		in  int
		out string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
	}

	for _, tt := range tests {
		solution := Numeral(tt.in)

		if solution != tt.out {
			t.Error(
				"For", tt.in,
				"expected", tt.out,
				"got", solution,
			)
		}
	}
}
