package main

import (
	"reflect"
	"testing"
)

func TestStarten(t *testing.T) {
	reizfolge := Reizfolge{
		Reiz{Anzahl: 5, Index: 1, Buchstabe: "A"},
		Reiz{Anzahl: 5, Index: 2, Buchstabe: "B"},
		Reiz{Anzahl: 5, Index: 3, Buchstabe: "A"},
		Reiz{Anzahl: 5, Index: 4, Buchstabe: "C"},
		Reiz{Anzahl: 5, Index: 5, Buchstabe: "D"},
	}
	expected := Reiz{Anzahl: 5, Index: 1, Buchstabe: "A"}

	uc := UseCases{reizfolge}
	actual := uc.Starten(5, 3)

	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected:\n%v got:\n%v\n", expected, actual)
	}
}
