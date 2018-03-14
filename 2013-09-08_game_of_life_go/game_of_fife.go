package game_of_life

import (

)

type Cell struct {
	state     string
	neighbors int
}

func (cell *Cell) Tick() {
	switch cell.neighbors {
	case 2, 3:
		cell.state = "alive"
	default:
		cell.state = "dead"
	}
}
