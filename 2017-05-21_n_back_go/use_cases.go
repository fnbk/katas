package main

type Reizfolge []Reiz

type UseCases struct {
	reizfolge Reizfolge
}

func (self *UseCases) Starten(anzahl, n int) Reiz {
	reiz := self.reizfolge.Peek()
	self.reizfolge.Pop()
}
