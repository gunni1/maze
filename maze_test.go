package main

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestMazeSize(t *testing.T) {
	maze := CreateMaze(2, 2)

	rX, rY := maze.Size()
	Equal(t, rX, 2)
	Equal(t, rY, 2)
}

func TestMazeSizeNotSquare(t *testing.T) {
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

func TestRollUnvisitedNeighbourTopLeft(t *testing.T) {
	maze := CreateMaze(2, 2)
	maze.cells[1][0].visited = true

	current := Position{0, 0}
	result, err := current.RollUnvisitedNeighbour(maze)

	Nil(t, err)
	Equal(t, result, Position{0, 1})
}

func TestRollUnvisitedNeighbourBottomRight(t *testing.T) {
	maze := CreateMaze(2, 2)
	maze.cells[0][1].visited = true

	current := Position{1, 1}
	result, err := current.RollUnvisitedNeighbour(maze)

	Nil(t, err)
	Equal(t, result, Position{1, 0})
}

func TestRollUnvisitedNeighbourNoUnvisitedRemain(t *testing.T) {
	maze := CreateMaze(2, 2)
	maze.cells[1][0].visited = true
	maze.cells[0][1].visited = true

	current := Position{0, 0}
	_, err := current.RollUnvisitedNeighbour(maze)
	NotNil(t, err)
}
