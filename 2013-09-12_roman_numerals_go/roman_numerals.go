package roman_numerals

import (

)

func Convert(number int) string {
	roman := ""

	if number >= 10 {
		tmpNumber := number/10 // remove last digit
		switch tmpNumber%10 {
		case 1:
			roman += "X"
		case 2:
			roman += "XX"
		case 3:
			roman += "XXX"
		case 4:
			roman += "XL"
		case 5:
			roman += "L"
		case 6:
			roman += "LX"
		case 7:
			roman += "LXX"
		case 8:
			roman += "LXXX"
		case 9:
			roman += "XC"
		}
	}

	switch number%10 {
	case 1:
		roman += "I"
	case 2:
		roman += "II"
	case 3:
		roman += "III"
	case 4:
		roman += "IV"
	case 5:
		roman += "V"
	case 6:
		roman += "VI"
	case 7:
		roman += "VII"
	case 8:
		roman += "VIII"
	case 9:
		roman += "IX"
	}

	return roman
}
