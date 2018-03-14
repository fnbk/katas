package game_of_life

import (
	"testing"
)

type testpair struct {
	in    Cell
	expectedState string
}

func verify(t *testing.T, cell Cell, expectedState string, actualState string) {
	if actualState != expectedState {
		t.Error(
			"For", cell,
			"expected", expectedState,
			"got", actualState,
		)
	}
}

func Test_CellDiesWithFewerThanTwoLiveNeighbors(t *testing.T) {
	tests := []testpair{
		{
			in: Cell{state: "alive", neighbors: 1},
			expectedState: "dead",
		},
		{
			in: Cell{state: "alive", neighbors: 0},
			expectedState: "dead",
		},
    }

	for _, tt := range tests {
		cell := tt.in
		cell.Tick()

		verify(t, cell, tt.expectedState, cell.state)
	}
}

func Test_CellLivesOnWithTwoOrThreeLiveNeighbors(t *testing.T) {
	tests := []testpair{
		{
			in :Cell{state: "alive", neighbors: 2},
			expectedState: "alive",
		},
		{
			in :Cell{state: "alive", neighbors: 3},
			expectedState: "alive",
		},
	}

	for _, tt := range tests {
		cell := tt.in
		cell.Tick()

		verify(t, cell, tt.expectedState, cell.state)
    }
}

func Test_CellDiesWithMoreThanTreeNeighbors(t *testing.T) {
	tests := []testpair{
		{
			in :Cell{state: "alive", neighbors: 4},
			expectedState: "dead",
		},
	}

	for _, tt := range tests {
		cell := tt.in
		cell.Tick()

		verify(t, cell, tt.expectedState, cell.state)
	}
}

func Test_CellBecomesAliveWithExactlyThreeLiveNeighbors(t *testing.T) {
	tests := []testpair{
		{
			in :Cell{state: "dead", neighbors: 3},
			expectedState: "alive",
		},
	}

	for _, tt := range tests {
		cell := tt.in
		cell.Tick()

		verify(t, cell, tt.expectedState, cell.state)
	}
}
