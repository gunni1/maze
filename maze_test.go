package main

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestRemoveWallsVert(t *testing.T) {
	maze := CreateMaze(2, 2)

	maze.RemoveWalls(Position{0, 0}, Position{1, 0})

	Equal(t, maze.cells[0][0].right, false)
	Equal(t, maze.cells[0][0].top, false)
	Equal(t, maze.cells[0][0].left, false)
	Equal(t, maze.cells[0][0].bottom, true)
	Equal(t, maze.cells[1][0].left, false)
	Equal(t, maze.cells[1][0].top, true)
	Equal(t, maze.cells[1][0].bottom, false)
	Equal(t, maze.cells[1][0].right, false)
}

func TestRemoveWallsHoriz(t *testing.T) {
	maze := CreateMaze(2, 2)

	maze.RemoveWalls(Position{1, 0}, Position{0, 0})

	Equal(t, maze.cells[0][0].right, true)
	Equal(t, maze.cells[0][0].top, false)
	Equal(t, maze.cells[0][0].left, false)
	Equal(t, maze.cells[0][0].bottom, false)
	Equal(t, maze.cells[0][1].left, true)
	Equal(t, maze.cells[0][1].top, false)
	Equal(t, maze.cells[0][1].bottom, false)
	Equal(t, maze.cells[0][1].right, false)
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
