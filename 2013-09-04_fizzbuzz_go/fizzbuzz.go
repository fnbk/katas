package main

import "strconv"

func Generate() [100]string {
	var numbers [100]string
	for idx, _ := range numbers {
		if idx%3 !=0 && idx%5 != 5 {
			numbers[idx] = strconv.Itoa(idx)
		}
		if idx%3 == 0 {
			numbers[idx] += "fizz"
		}
		if idx%5 == 0 {
			numbers[idx] += "buzz"
		}
	}
	numbers[0] = "0"
	return numbers
}
