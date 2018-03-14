package bowling_game

type Game struct {
	frames [10]Frame
}

type Frame struct {
	rolls [3]int
}

func (frame Frame) Total() int {
	var sum int

	for _, roll := range frame.rolls {
		sum += roll
	}

	return sum
}

func (game Game) Total() int {
	var sum int

	for idx, frame := range game.frames {
		total := frame.Total()
		isStrike := frame.rolls[0] == 10
		bonus := 0

		switch isStrike {
		case true:
			if idx < 8 {
				nextRoll1 := game.frames[idx+1].rolls[0]
				nextRoll2 := game.frames[idx+1].rolls[1]

				if nextRoll1 == 10 {
					nextRoll2 = game.frames[idx+2].rolls[0]
				}

				bonus = nextRoll1 + nextRoll2
			} else if idx == 8 {
				nextRoll1 := game.frames[idx+1].rolls[0]
				nextRoll2 := game.frames[idx+1].rolls[1]

				bonus = nextRoll1 + nextRoll2
			} else if idx == 9 {
				nextRoll1 := game.frames[idx].rolls[1]
				nextRoll2 := game.frames[idx].rolls[2]

				bonus = nextRoll1 + nextRoll2
			}
		case false:
			if total > 9 && idx < 9 {
				bonus = game.frames[idx+1].rolls[0]
			}
		}

		sum += total + bonus
	}

	return sum
}
