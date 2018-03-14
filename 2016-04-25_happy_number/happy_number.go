package happy_number

import (
	"math"
	"strconv"
)

//
// RECURSIVE
//

func happyNumberRecursive(num int, seq map[int]bool) bool {
	var retVal bool

	if num == 1 {
		retVal = true
	} else if seq[num] {
		retVal = false
	} else {
		total := calculateNext(num)
		seq[num] = true
		retVal = happyNumberRecursive(total, seq)
	}

	return retVal
}

//
// UNROLLED
//

type HappyNumberUnrolled struct {
	seq map[int]bool
}

func (self HappyNumberUnrolled) IsHappy(num int) bool {
	var retVal bool
	self.seq = map[int]bool{}
	maxSeq := 100
	next := num

	for i := 0; i < maxSeq; i++ {
		if next == 1 {
			retVal = true
			break
		} else if self.seq[next] {
			retVal = false
			break
		} else {
			self.seq[next] = true
			next = calculateNext(next)
		}
	}

	return retVal
}

//
// helper
//

func calculateNext(num int) int {
	total := 0
	numDigits := len(strconv.Itoa(num))
	for i := 0; i < numDigits; i++ {
		position := math.Pow(float64(10), float64(i))
		digit := int(num/int(position)) % 10
		subTotal := math.Pow(float64(digit), 2)
		total += int(subTotal)
	}
	return total
}

//
// wrapper
//

type AlgorithmType int

const (
	RECURSIVE AlgorithmType = 1 + iota
	UNROLLED
)

func HappyNumber(num int, algType AlgorithmType) bool {
	var retVal bool

	switch algType {
	case RECURSIVE:
		retVal = happyNumberRecursive(num, map[int]bool{})
	case UNROLLED:
		otto := HappyNumberUnrolled{}
		retVal = otto.IsHappy(num)
	}

	return retVal
}
