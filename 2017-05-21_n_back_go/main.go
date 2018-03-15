package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

// main Manueller Test für den PO (Überprüfungstest)
func main() {
	ui := Ui{}

	cfg := Cfg{
		N:           3,
		Reizdauer:   2000,
		AnzahlReize: 10,
		Probant:     "Peter",
	}

	onStart := func() {
		r := Reiz{
			Buchstabe: "A",
			Index:     1,
			Anzahl:    10,
		}
		ui.Reiz(r)
	}

	antwortCount := 0
	trefferCount := 0
	onAntwort := func(a Antwort) {
		if a == Wiederholung {
			trefferCount++
		}
		r := Reiz{}
		switch antwortCount {
		case 0:
			r = Reiz{
				Buchstabe: "B",
				Index:     2,
				Anzahl:    10,
			}
		case 1:
			r = Reiz{
				Buchstabe: "C",
				Index:     3,
				Anzahl:    10,
			}
		}
		if reflect.DeepEqual(r, Reiz{}) {
			ui.Ergebnis(Ergebnis{trefferCount})
		} else {
			antwortCount++
			ui.Reiz(r)
		}
	}

	ui.OnStart = onStart
	ui.OnAntwort = onAntwort

	ui.Config(cfg)
}

type Cfg struct {
	N           int
	Reizdauer   int
	AnzahlReize int
	Probant     string
}

type Reiz struct {
	Buchstabe string
	Index     int
	Anzahl    int
}

type Ergebnis struct {
	Treffer int
}

type Antwort int

const (
	Wiederholung Antwort = 1 << iota
	Neu
)

type StartCallback func()
type AntwortCallback func(Antwort)

type Ui struct {
	OnStart   StartCallback
	OnAntwort AntwortCallback
}

func (self *Ui) Config(c Cfg) {
	out := fmt.Sprintf("Testparameter:\n  N: %d\n  Reizdauer: %d msec\n  Anzahl Reize: %d\n  Probant: %s\n", c.N, c.Reizdauer, c.AnzahlReize, c.Probant)
	fmt.Println(out)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // wartet auf \n in stdin

	self.OnStart()
}

func (self *Ui) Reiz(r Reiz) {
	out := fmt.Sprintf("Reiz %d/%d: %s.....\r", r.Index, r.Anzahl, r.Buchstabe)
	fmt.Printf(out)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()        // wartet auf \n in stdin
	txt := scanner.Text() // liest alle Zeichen aus stdin bis \n ein

	antwort := Antwort(0)
	if txt == "w" {
		antwort = Wiederholung
	} else if txt == " " {
		antwort = Neu
	}

	self.OnAntwort(antwort)
}

func (self *Ui) Ergebnis(e Ergebnis) {
	out := fmt.Sprintf("Ergebnis %d\n", e.Treffer)
	fmt.Println(out)
}
