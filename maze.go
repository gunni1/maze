package main

import (
	"errors"
	"image"
	"image/color"
	"math/rand"
)

type Position struct {
	x int
	y int
}

type Maze struct {
	cells [][]Cell
}

//Represent a cell of a maze. True for top,left,right,bottom means open in this direction
type Cell struct {
	top     bool
	left    bool
	right   bool
	bottom  bool
	visited bool
}

func (maze Maze) Size() (int, int) {
	dx := 0
	dy := 0
	for i, cols := range maze.cells {
		dy = len(cols)
		dx = i + 1
	}
	//TODO implement
	return dx, dy
}

//Return a a random unvisited neighbour Cell in the Maze.
//Return Error, if there are no unvisited Cells remain.
func (current Position) RollUnvisitedNeighbour(maze Maze) (Position, error) {
	candidates := make([]Position, 0)

	candidates = appendIfValidUnvisited(candidates, Position{current.x, current.y - 1}, maze)
	candidates = appendIfValidUnvisited(candidates, Position{current.x - 1, current.y}, maze)
	candidates = appendIfValidUnvisited(candidates, Position{current.x + 1, current.y}, maze)
	candidates = appendIfValidUnvisited(candidates, Position{current.x, current.y + 1}, maze)

	if len(candidates) == 0 {
		return Position{}, errors.New("no unvisited remain")
	} else {
		return candidates[rand.Intn(len(candidates))], nil
	}
}

func appendIfValidUnvisited(candidates []Position, candidate Position, maze Maze) []Position {
	xMax, yMax := maze.Size()
	inBounds := candidate.x < xMax && candidate.x >= 0 && candidate.y < yMax && candidate.y >= 0
	if inBounds && !maze.cells[candidate.x][candidate.y].visited {
		return append(candidates, candidate)
	} else {
		return candidates
	}
}

//Positions are not allowed to be out of the bounds of the maze
func (maze Maze) RemoveWalls(curr Position, neighb Position) error {
	if curr.x == neighb.x {
		//vertikal
		if curr.y > neighb.y {
			maze.cells[curr.x][curr.y].top = true
			maze.cells[neighb.x][neighb.y].bottom = true
		} else {
			maze.cells[curr.x][curr.y].bottom = true
			maze.cells[neighb.x][neighb.y].top = true
		}
	} else {
		if curr.x > neighb.x {
			maze.cells[curr.x][curr.y].left = true
			maze.cells[neighb.x][neighb.y].right = true
		} else {
			maze.cells[curr.x][curr.y].right = true
			maze.cells[neighb.x][neighb.y].left = true
		}
	}
	return nil
}

func CreateMaze(dx int, dy int) Maze {
	maze := make([][]Cell, dx)
	for i := range maze {
		maze[i] = make([]Cell, dy)
	}
	return Maze{cells: maze}
}

//Visualize a Maze with each Cell as its 3x3 px representation
func (maze Maze) Visualize() image.Image {
	dx, dy := maze.Size()
	img := image.NewGray(image.Rect(0, 0, dx*3, dy*3))

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			cell := maze.cells[i][j]
			//Render 1 particular Cell as 3x3 pixel representation
			//TOP Line
			imageX := i * 3
			imageY := j * 3

			img.Set(imageX, imageY, color.Black)
			img.Set(imageX+1, imageY, whiteIf(cell.top))
			img.Set(imageX+2, imageY, color.Black)
			//Middle Line
			img.Set(imageX, imageY+1, whiteIf(cell.left))
			img.Set(imageX+1, imageY+1, color.White)
			img.Set(imageX+2, imageY+1, whiteIf(cell.right))
			//Bottom Line
			img.Set(imageX, imageY+2, color.Black)
			img.Set(imageX+1, imageY+2, whiteIf(cell.bottom))
			img.Set(imageX+2, imageY+2, color.Black)
		}
	}
	return img
}

func whiteIf(direction bool) color.Color {
	if direction {
		return color.White
	} else {
		return color.Black
	}
}
