package main

import (
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestBerechneZeilen(t *testing.T) {
	tests := []struct {
		inputStr string
		count    int
	}{
		{"", 0},
		{"\n", 0},
		{"\n\n", 0},
		{"a", 1},
		{"a\nb\n", 2},
		{"\t", 0},
		{" ", 0},
		{"\ta", 1},
		{"//", 0},
		{"//a", 0},
		{"a//", 1},
		{`"//"`, 1},
		{`//"a"`, 0},
		// {`/*"a"*/`, 0},
		// {`/**/"a"`, 1},
		// {`a/**/`, 1},
		// {`a/*
		// 		*/`, 1},
		// {`a/*
		// 		*/
		// 		`, 1},
		// {`a/*
		// 		*/
		// 		b`, 2},
		// {`/*"*/"b"`, 1},
	}

	t.Log("Gegeben ist die Notwändigkeit die Anzahl der Zeilen einer Zeichenkette zu testen")
	{
		for i, tt := range tests {
			t.Logf("\tTest: %d\tWenn die Zeichenkette %q hineingereicht wird", i, tt.inputStr)

			berecheneteZeilen := berecheneZeilen(tt.inputStr)

			if berecheneteZeilen == tt.count {
				t.Logf("\t%s\tSollte %d zurückgegeben werden", succeed, tt.count)
			} else {
				t.Errorf("\t%s\tSollte %d zurückgegeben werden Rückgabe: %d", failed, tt.count, berecheneteZeilen)
			}
		}
	}
}

func TestEnthältCode(t *testing.T) {
	tests := []struct {
		zeile string
		code  bool
	}{
		{"//", false},
		{"a//", true},
		{"", false},
		{"\t", false},
		{"\ta", true},
		{"\t a ", true},
		{" ", false},
		{" \t", false},
		{"\t\t\ts\t", true},
	}

	t.Log("Gegeben ist die Notwändigkeit zu prüfen, ob eine Zeile Code enthält")
	{
		for i, tt := range tests {
			t.Logf("\tTest: %d\tWenn die Zeile %q hineingereicht wird", i, tt.zeile)

			ergebnis := enthaeltCode(tt.zeile)

			if ergebnis == tt.code {
				t.Logf("\t%s\tSollte %t zurückgegeben werden", succeed, tt.code)
			} else {
				t.Errorf("\t%s\tSollte %t zurückgegeben werden. Rückgabe: %t", failed, tt.code, ergebnis)
			}
		}
	}
}
