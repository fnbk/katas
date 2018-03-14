package mars_rover

import (

)

type Rover struct {
	x, y int
	d rune
	grid Grid
}

type Grid struct {
	x, y int
	obstacles []Obstacle
}

type Obstacle struct {
	x, y int
}

func (rover *Rover) Go(commands string) {
	for _, command := range commands {
		switch command {
		case 'f':
			rover.forward()
			for _, obstacle := range rover.grid.obstacles {
				if rover.obstacleDetected(obstacle) {
					rover.backward()
					return
				}
			}
		case 'b':
			rover.backward()
			for _, obstacle := range rover.grid.obstacles {
				if rover.obstacleDetected(obstacle) {
					rover.forward()
					return
				}
			}
		case 'l':
			rover.left()
			for _, obstacle := range rover.grid.obstacles {
				if rover.obstacleDetected(obstacle) {
					rover.right()
					return
				}
			}
		case 'r':
			rover.right()
			for _, obstacle := range rover.grid.obstacles {
				if rover.obstacleDetected(obstacle) {
					rover.left()
					return
				}
			}
		}
	}
}

func (rover *Rover) forward() {
	switch rover.d {
	case 'N':
		if rover.y == 10 {
			rover.y = -10
		} else {
			rover.y += 1
		}
	case 'E':
		if rover.x == 10 {
			rover.x = -10
		} else {
			rover.x += 1
		}
	case 'W':
		if rover.x == -10 {
			rover.x = 10
		} else {
			rover.x -= 1
		}
	case 'S':
		if rover.y == -10 {
			rover.y = 10
		} else {
			rover.y -= 1
		}
	}
}

func (rover *Rover) backward() {
	switch rover.d {
	case 'N':
		if rover.y == -10 {
			rover.y = 10
		} else {
			rover.y -= 1
		}
	case 'E':
		if rover.x == -10 {
			rover.x = 10
		} else {
			rover.x -= 1
		}
	case 'W':
		if rover.x == 10 {
			rover.x = -10
		} else {
			rover.x += 1
		}
	case 'S':
		if rover.y == 10 {
			rover.y = -10
		} else {
			rover.y += 1
		}
	}
}

func (rover *Rover) left() {
	switch rover.d {
	case 'N':
		rover.d = 'W'
	case 'E':
		rover.d = 'N'
	case 'W':
		rover.d = 'S'
	case 'S':
		rover.d = 'E'
	}
}

func (rover *Rover) right() {
	switch rover.d {
	case 'N':
		rover.d = 'E'
	case 'E':
		rover.d = 'S'
	case 'W':
		rover.d = 'N'
	case 'S':
		rover.d = 'W'
	}
}

func (rover *Rover) obstacleDetected(obstacle Obstacle) bool {
	if rover.x == obstacle.x && rover.y == obstacle.y {
		return true
	}
	return false
}
