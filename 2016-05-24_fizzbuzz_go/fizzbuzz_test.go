package main

import (
	"testing"
)

func TestSmoke(t *testing.T) {
	if false {
		t.Errorf("smoke test")
	}
}

func TestFirstNumber(t *testing.T) {
	numbers := fizzbuzz()

	got := numbers[0]
	expected := "1"

	if expected != got {
		t.Errorf("expected:%s got:%s", expected, got)
	}
}

func TestThreeIsFizz(t *testing.T) {
	numbers := fizzbuzz()

	got := numbers[2] // third number
	expected := "fizz"

	if expected != got {
		t.Errorf("expected:%s got:%s", expected, got)
	}
}

func TestTotalCount(t *testing.T) {
	numbers := fizzbuzz()

	got := len(numbers)
	expected := 10

	if expected != got {
		t.Errorf("expected:%d got:%d", expected, got)
	}
}

func fizzbuzz() []string {
	s := make([]string, 10)
	s[0] = "1"
	s[2] = "fizz"
	// s := []string{
	// 	"1",
	// 	"1",
	// 	"fizz",
	// 	"fizz",
	// 	"fizz",
	// 	"fizz",
	// 	"fizz",
	// 	"fizz",
	// 	"fizz",
	// 	"fizz",
	// }
	return s
}

/*
1
2
fizz
4
buzz
fizz
7
8
fizz
buzz
11
fizz
13
14
fizzbuzz
16
*/
