package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

func main() {
	dx := 3
	dy := 3

	maze := CreateMaze(dx, dy)

	//img := image.NewGray(image.Rect(0,0,dx*3,dy*3))

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
