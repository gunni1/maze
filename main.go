package main

import (
	"fmt"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	dx := 3
	dy := 3

	maze := CreateMaze(dx, dy)

	//TEST
	/*
		maze.cells[0][0].top = true
		maze.cells[0][0].right = true
		maze.cells[1][0].bottom = true
		maze.cells[1][0].left = true
		maze.cells[1][1].top = true
		maze.cells[1][1].left = true
		maze.cells[0][1].left = true
		maze.cells[0][1].right = true
	*/

	//Pick random start
	dxStart := rand.Intn(dx)
	dyStart := rand.Intn(dy)
	fmt.Printf("Start: %d, %d \n", dxStart, dyStart)
	maze.cells[dxStart][dyStart].visited = true

	//Depth-First iterative
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

	image := maze.Visualize(dx, dy)

	file, _ := os.Create("test.png")
	png.Encode(file, image)
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
