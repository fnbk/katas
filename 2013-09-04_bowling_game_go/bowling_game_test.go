package bowling_game

import (
	"testing"
)

func TestTotalScoreIsZeroIfEveryTurnIsEmpty(t *testing.T) {
	game := Game{
		[10]Frame{
			Frame{[3]int{0, 0, 0}},
			Frame{[3]int{0, 0, 0}},
			Frame{[3]int{0, 0, 0}},
			Frame{[3]int{0, 0, 0}},
			Frame{[3]int{0, 0, 0}},
			Frame{[3]int{0, 0, 0}},
			Frame{[3]int{0, 0, 0}},
			Frame{[3]int{0, 0, 0}},
			Frame{[3]int{0, 0, 0}},
			Frame{[3]int{0, 0, 0}}}}

	if game.Total() != 0 {
		t.Errorf("Total score is not zero if every turn is empty: %v", game.Total())
	}
}

type GameWithoutSparesOrStrikes struct {
	in  Game
	out int
}

func TestTotalScoreIsTheSumOfAllTurnsWhenThereAreNoSparesOrStrikes(t *testing.T) {
	gameWithoutSparesOrStrikes := [...]GameWithoutSparesOrStrikes{
		GameWithoutSparesOrStrikes{
			in: Game{[10]Frame{
				Frame{[3]int{9, 0, 0}},
				Frame{[3]int{9, 0, 0}},
				Frame{[3]int{9, 0, 0}},
				Frame{[3]int{9, 0, 0}},
				Frame{[3]int{9, 0, 0}},
				Frame{[3]int{9, 0, 0}},
				Frame{[3]int{9, 0, 0}},
				Frame{[3]int{9, 0, 0}},
				Frame{[3]int{9, 0, 0}},
				Frame{[3]int{9, 0, 0}},
			}}, // total: 90
			out: 90},
		GameWithoutSparesOrStrikes{
			in: Game{[10]Frame{
				Frame{[3]int{1, 8, 0}}, // 9
				Frame{[3]int{2, 7, 0}}, // 9
				Frame{[3]int{3, 6, 0}}, // 9
				Frame{[3]int{4, 5, 0}}, // 9
				Frame{[3]int{5, 4, 0}}, // 9
				Frame{[3]int{6, 3, 0}}, // 9
				Frame{[3]int{7, 2, 0}}, // 9
				Frame{[3]int{8, 1, 0}}, // 9
				Frame{[3]int{9, 0, 0}}, // 9
				Frame{[3]int{0, 0, 0}}, // 0
			}}, // total: 81
			out: 81},
	}

	for _, elem := range gameWithoutSparesOrStrikes {
		game := elem.in
		total := elem.out
		if game.Total() != total {
			t.Errorf("Total score is not the sum of all turns, expected %v, got %v", total, game.Total())
		}
	}
}

type GameWithSpares struct {
	in  Game
	out int
}

func TestTotalScoreCountsSpares(t *testing.T) {
	gameWithSpares := [...]GameWithSpares{
		GameWithSpares{
			in: Game{[10]Frame{
				Frame{[3]int{9, 1, 0}}, // spare
				Frame{[3]int{2, 0, 0}}, // next throw: 2
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}}},
			}, // total: (9+1+2) + 2 = 14
			out: 14},
		GameWithSpares{
			in: Game{[10]Frame{
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{0, 0, 0}},
				Frame{[3]int{9, 1, 0}}}, // spare
			}, // total: 10
			out: 10},
		GameWithSpares{
			in: Game{[10]Frame{
				Frame{[3]int{5, 5, 0}},
				Frame{[3]int{5, 5, 0}},
				Frame{[3]int{5, 5, 0}},
				Frame{[3]int{5, 5, 0}},
				Frame{[3]int{5, 5, 0}},
				Frame{[3]int{5, 5, 0}},
				Frame{[3]int{5, 5, 0}},
				Frame{[3]int{5, 5, 0}},
				Frame{[3]int{5, 5, 0}},
				Frame{[3]int{5, 5, 5}}},
			}, // total: 150
			out: 150},
	}

	for _, elem := range gameWithSpares {
		game := elem.in
		total := elem.out
		if game.Total() != total {
			t.Errorf("Total score does not take into account spares, expected %v, got %v", total, game.Total())
		}
	}
}

type GameWithStrikes struct {
	in  Game
	out int
}

func TestTotalScoreCountsStrikes(t *testing.T) {
	gameWithStrikes := [...]GameWithStrikes{
		GameWithStrikes{
			in: Game{
				[10]Frame{
					Frame{[3]int{10, 0, 0}}, // strike: 10 + 1 (next roll) + 1 (next roll): 12
					Frame{[3]int{1, 1, 0}},  // 2
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
				}}, // total: 14
			out: 14},
		GameWithStrikes{
			in: Game{
				[10]Frame{
					Frame{[3]int{10, 0, 0}}, // strike: 10 + 10 (next roll) + 1 (next roll): 21
					Frame{[3]int{10, 0, 0}}, // strike: 10 + 1  (next roll) + 0 (next roll): 11
					Frame{[3]int{1, 0, 0}},  // 1
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
				}}, // total: 33
			out: 33},
		GameWithStrikes{
			in: Game{
				[10]Frame{
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{0, 0, 0}},
					Frame{[3]int{10, 0, 0}},  // strike: 10 + 10 (next roll) + 10 (next roll): 30
					Frame{[3]int{10, 10, 0}}, // strike: 10 + 10 (next_roll) + 10 + 0 (next roll): 30
				}}, // total: 60
			out: 60},
		GameWithStrikes{
			in: Game{
				[10]Frame{
					Frame{[3]int{10, 0, 0}},
					Frame{[3]int{10, 0, 0}},
					Frame{[3]int{10, 0, 0}},
					Frame{[3]int{10, 0, 0}},
					Frame{[3]int{10, 0, 0}},
					Frame{[3]int{10, 0, 0}},
					Frame{[3]int{10, 0, 0}},
					Frame{[3]int{10, 0, 0}},
					Frame{[3]int{10, 0, 0}},
					Frame{[3]int{10, 10, 0}},
				}}, // total: 300
			out: 300},
	}

	for _, elem := range gameWithStrikes {
		game := elem.in
		total := elem.out
		if game.Total() != total {
			t.Errorf("Total score does not take into account strikes, expected %v, got %v", total, game.Total())
		}
	}
}
