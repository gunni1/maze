package main

import (
	"image"
	"image/color"
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

//Return a random unvisited neighbour cell of the maze
func (current Position) RollUnvisitedNeighbour(maze Maze) (Position, error) {
	return Position{0, 0}, nil
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
	//TODO out of bounds fehler erzeugen
}

func CreateMaze(dx int, dy int) Maze {
	maze := make([][]Cell, dy)
	for i := range maze {
		maze[i] = make([]Cell, dx)
	}
	return Maze{cells: maze}
}

func (maze Maze) CreateGridImage() {

}

//Visualize a Maze with each Cell as its 3x3 px representation
func (maze Maze) Visualize(dx int, dy int) image.Image {
	//TODO: find dimension programmatically
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
