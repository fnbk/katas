package string_calculator

import (
	"strings"
	"strconv"
	"unicode/utf8"
)

func Add(numbers string) int {
	var result int
	var delimiters []rune
	var hasDelimiterLine bool
	delimiter1, _ := utf8.DecodeRuneInString(",")
	delimiter2, _ := utf8.DecodeRuneInString("\n")
	delimiters = []rune{delimiter1, delimiter2}

	if len(numbers) != 0 && string(numbers[0]) == "/" && string(numbers[1]) == "/" {
		hasDelimiterLine = true
		endOfFirstLine := strings.Index(numbers, "\n")
		//firstLine      := numbers[1:endOfFirstLine-1]
		numbersWithoutFirstLine := numbers[endOfFirstLine+1:]
		numbers = numbersWithoutFirstLine
		delimiter3, _ := utf8.DecodeRuneInString(";")
		delimiters = []rune{delimiter3}
	}

	arrayOfNumbers := strings.FieldsFunc(numbers, func(r rune) bool {
			var isDelimiter bool = true

			if !hasDelimiterLine {
				isDelimiter = (r == ',' || r == '\n')
			} else {
				for _, delimiter := range delimiters {
					isDelimiter = (isDelimiter && (r == delimiter))
				}
			}
			return isDelimiter
		})

	var number int
	for _, stringNumber := range arrayOfNumbers {
		number, _ = strconv.Atoi(stringNumber)
		result += number
	}

	return result
}
