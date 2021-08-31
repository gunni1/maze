package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

func main() {
	dx := 3
	dy := 3

	maze := createMaze(dx, dy)
	//img := image.NewGray(image.Rect(0,0,dx*3,dy*3))

	//Pick random start
	dxStart := rand.Intn(dx)
	dyStart := rand.Intn(dy)
	maze.cells[dxStart][dyStart].visited = true

	//Depth-First iterative
	stack := Stack{stack: list.New()}
	stack.push(Position{dxStart, dyStart})

	//Solange stack nicht leer --
	current := stack.pop()
	neighbour := current.rollUnvisitedNeighbour(maze)
	maze.removeWalls(current, neighbour)
	//Zwischen current und gewählter die Wände entfernen
	//gewählte als besucht markieren, push in stack

	fmt.Print(maze)
}

//Return a random unvisited neighbour cell of the maze
func (current Position) rollUnvisitedNeighbour(maze Maze) Position {
	return Position{0, 0}
}

func (maze Maze) removeWalls(curr Position, neighb Position) {

}

func createMaze(dx int, dy int) Maze {
	maze := make([][]Cell, dy)
	for i := range maze {
		maze[i] = make([]Cell, dx)
	}
	initial := Cell{top: true, left: true, right: true, bottom: true, visited: false}
	for i, row := range maze {
		for j := range row {
			maze[i][j] = initial
		}
	}

	return Maze{cells: maze}
}

//Visualize a Cell as a grid image 3x3 representation
func (cell Cell) visualize() {

}

type Stack struct {
	stack *list.List
}

func (s Stack) push(pos Position) {
	s.stack.PushFront(pos)
}

func (s Stack) pop() Position {
	return Position{0, 0}
}

type Position struct {
	x int
	y int
}

type Maze struct {
	cells [][]Cell
}

type Cell struct {
	top     bool
	left    bool
	right   bool
	bottom  bool
	visited bool
}
