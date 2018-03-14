package roman_numerals


import (

)

func Numeral(digit int) string {
	numeral := "-1"

	switch digit {
	case 1:
		numeral = "I"
	case 2:
		numeral = "II"
	case 3:
		numeral = "III"
	case 4:
		numeral = "IV"
	case 5:
		numeral = "V"
	case 6:
		numeral = "VI"
	case 7:
		numeral = "VII"
	case 8:
		numeral = "VIII"
	case 9:
		numeral = "IX"
	case 10:
		numeral = "X"
	}

	return numeral
}





