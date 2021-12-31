package main

import (
	"math/rand"
	"time"
)

//Depth-First iterative
func GeneratePathDeepFirst(maze Maze) {
	dx, dy := maze.Size()

	rand.Seed(time.Now().UnixNano())
	dxStart := rand.Intn(dx)
	dyStart := rand.Intn(dy)

	maze.cells[dxStart][dyStart].visited = true

	stack := Stack{Data: make([]Position, 0, dx*dy)}
	stack.push(Position{dxStart, dyStart})

	for stack.hasElements() {
		current := stack.pop()

		neighbour, err := current.RollUnvisitedNeighbour(maze)
		if err != nil {
			continue
		}
		//Current goes back to the stack
		stack.push(current)

		//Remove Walls and set chosen Cell as visited
		maze.RemoveWalls(current, neighbour)
		maze.cells[neighbour.x][neighbour.y].visited = true
		stack.push(neighbour)
	}

}

type Stack struct {
	Data []Position
}

func (s Stack) hasElements() bool {
	return len(s.Data) > 0
}

func (s *Stack) push(pos Position) {
	s.Data = append(s.Data, pos)
}

func (s *Stack) pop() Position {
	cell, cdr := s.Data[len(s.Data)-1], s.Data[:len(s.Data)-1]
	s.Data = cdr
	return cell
}
