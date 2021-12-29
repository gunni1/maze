package main

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func MazeSize(t *testing.T) {
	maze := CreateMaze(2, 2)

	rX, rY := maze.Size()
	Equal(t, rX, 2)
	Equal(t, rY, 2)
}

func MazeSizeNotSquare(t *testing.T) {
	maze := CreateMaze(2, 3)

	rX, rY := maze.Size()
	Equal(t, rX, 2)
	Equal(t, rY, 3)
}

func TestRemoveWallsVert(t *testing.T) {
	maze := CreateMaze(2, 2)

	maze.RemoveWalls(Position{0, 0}, Position{0, 1})

	Equal(t, maze.cells[0][0].right, false)
	Equal(t, maze.cells[0][0].top, false)
	Equal(t, maze.cells[0][0].left, false)
	Equal(t, maze.cells[0][0].bottom, true)
	Equal(t, maze.cells[0][1].left, false)
	Equal(t, maze.cells[0][1].top, true)
	Equal(t, maze.cells[0][1].bottom, false)
	Equal(t, maze.cells[0][1].right, false)
}

func TestRemoveWallsHoriz(t *testing.T) {
	maze := CreateMaze(2, 2)

	maze.RemoveWalls(Position{1, 1}, Position{0, 1})

	Equal(t, maze.cells[0][1].right, true)
	Equal(t, maze.cells[0][1].top, false)
	Equal(t, maze.cells[0][1].left, false)
	Equal(t, maze.cells[0][1].bottom, false)
	Equal(t, maze.cells[1][1].left, true)
	Equal(t, maze.cells[1][1].top, false)
	Equal(t, maze.cells[1][1].bottom, false)
	Equal(t, maze.cells[1][1].right, false)
}

func NoDiagonal(t *testing.T) {
	maze := CreateMaze(2, 2)

	err := maze.RemoveWalls(Position{0, 0}, Position{1, 1})
	NotNil(t, err)
}

func NoJump(t *testing.T) {
	maze := CreateMaze(3, 3)

	err := maze.RemoveWalls(Position{0, 0}, Position{0, 2})
	NotNil(t, err)
}

func RollUnvisitedNeighbourTopLeft(t *testing.T) {
	maze := CreateMaze(2, 2)
	maze.cells[1][0].visited = true

	current := Position{0, 0}
	result, err := current.RollUnvisitedNeighbour(maze)

	Nil(t, err)
	Equal(t, result, Position{0, 1})
}

func RollUnvisitedNeighbourBottomRight(t *testing.T) {
	maze := CreateMaze(2, 2)
	maze.cells[0][1].visited = true

	current := Position{1, 1}
	result, err := current.RollUnvisitedNeighbour(maze)

	Nil(t, err)
	Equal(t, result, Position{1, 0})
}

func RollUnvisitedNeighbourNoUnvisitedRemain(t *testing.T) {
	maze := CreateMaze(2, 2)
	maze.cells[1][0].visited = true
	maze.cells[0][0].visited = true

	current := Position{0, 0}
	_, err := current.RollUnvisitedNeighbour(maze)
	NotNil(t, err)
}
