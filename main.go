package main

import (
	"container/list"
	"fmt"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	dx := 2
	dy := 2

	maze := CreateMaze(dx, dy)

	//TEST
	maze.cells[0][0].top = true
	maze.cells[0][0].right = true
	maze.cells[1][0].bottom = true
	maze.cells[1][0].left = true
	maze.cells[1][1].top = true
	maze.cells[1][1].left = true
	maze.cells[0][1].left = true
	maze.cells[0][1].right = true

	image := maze.Visualize(dx, dy)

	file, _ := os.Create("test.png")
	png.Encode(file, image)

	return

	//Pick random start
	dxStart := rand.Intn(dx)
	dyStart := rand.Intn(dy)
	maze.cells[dxStart][dyStart].visited = true

	//Depth-First iterative
	stack := Stack{stack: list.New()}
	stack.push(Position{dxStart, dyStart})

	for stack.hasElements() {
		current := stack.pop()
		neighbour, err := current.RollUnvisitedNeighbour(maze)
		if err != nil {
			continue
		}
		maze.RemoveWalls(current, neighbour)
		maze.cells[neighbour.x][neighbour.y].visited = true
		stack.push(neighbour)
	}
	fmt.Print(maze)
}

type Stack struct {
	stack *list.List
}

func (s Stack) hasElements() bool {
	return s.stack.Len() > 0
}

func (s Stack) push(pos Position) {
	s.stack.PushFront(pos)
}

func (s Stack) pop() Position {
	return Position{0, 0}
}
