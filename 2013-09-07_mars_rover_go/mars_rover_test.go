package mars_rover

import (
	"testing"
)

func CompareAndVerifyRover(t *testing.T, rover1, rover2 Rover) {
	if (rover1.x != rover2.x) {
		t.Errorf("Expected starting point x to be %v, but was %v", rover2.x, rover1.x)
	}
	if (rover1.y != rover2.y) {
		t.Errorf("Expected starting point y to be %v, but was %v", rover2.y, rover1.y)
	}
	if (rover1.d != rover2.d) {
		t.Errorf("Expected direction to be %v, but was %v", rover2.d, rover1.d)
	}
}

func Test_RoverHasInitialStartingPointAndDirection(t *testing.T) {
	rover := Rover{x:0, y:0, d:'N'}

	CompareAndVerifyRover(t, rover, Rover{x:0, y:0, d:'N'})
}

func Test_RoverReceivesACharacterArrayOfCommands(t *testing.T) {
	rover := Rover{x:0, y:0, d:'N'}

	rover.Go("")

	CompareAndVerifyRover(t, rover, Rover{x:0, y:0, d:'N'})
}

func Test_RoverMoveForewardAndBackward(t *testing.T) {
	var tests = []struct{
		in Rover
		command string
		out Rover
	}{
		{Rover{x:0, y:0, d:'N'}, "f", Rover{x:0, y:1, d:'N'}},
		{Rover{x:0, y:0, d:'N'}, "b", Rover{x:0, y:-1, d:'N'}},
	}

	for _, tt := range tests {
		tt.in.Go(tt.command)
		CompareAndVerifyRover(t, tt.in, tt.out)
	}
}

func Test_RoverMoveLeftAndRight(t *testing.T) {
	var tests = []struct{
		in Rover
		command string
		out Rover
	}{
		{Rover{x:0, y:0, d:'N'}, "l", Rover{x:0, y:0, d:'W'}},
		{Rover{x:0, y:0, d:'N'}, "r", Rover{x:0, y:0, d:'E'}},
	}

	for _, tt := range tests {
		tt.in.Go(tt.command)
		CompareAndVerifyRover(t, tt.in, tt.out)
	}
}

func Test_RoverLivesOnASphare(t *testing.T) {
	var tests = []struct{
		in Rover
		command string
		out Rover
	}{
		{Rover{x:0, y:10, d:'N'}, "f", Rover{x:0, y:-10, d:'N'}},
		{Rover{x:10, y:0, d:'E'}, "f", Rover{x:-10, y:0, d:'E'}},
		{Rover{x:-10, y:0, d:'W'}, "f", Rover{x:10, y:0, d:'W'}},
		{Rover{x:0, y:-10, d:'S'}, "f", Rover{x:0, y:10, d:'S'}},
		{Rover{x:0, y:-10, d:'N'}, "b", Rover{x:0, y:10, d:'N'}},
		{Rover{x:-10, y:0, d:'E'}, "b", Rover{x:10, y:0, d:'E'}},
		{Rover{x:10, y:0, d:'W'}, "b", Rover{x:-10, y:0, d:'W'}},
		{Rover{x:0, y:10, d:'S'}, "b", Rover{x:0, y:-10, d:'S'}},
	}

	for _, tt := range tests {
		tt.in.Go(tt.command)
		CompareAndVerifyRover(t, tt.in, tt.out)
	}
}

func Test_RoverStopsAtObstacle(t *testing.T) {
	var grid Grid = Grid{x:10, y:10, obstacles: []Obstacle{{x:0, y:1}, {x:0, y:-1}}}
	var tests = []struct{
		in Rover
		command string
		out Rover
	}{
		{Rover{x:0, y:0, d:'N', grid: grid}, "f", Rover{x:0, y:0, d:'N'}},
		{Rover{x:0, y:0, d:'N', grid: grid}, "rflflf", Rover{x:1, y:1, d:'W'}},
	}

	for _, tt := range tests {
		tt.in.Go(tt.command)
		CompareAndVerifyRover(t, tt.in, tt.out)
	}
}
