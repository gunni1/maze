package main

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

func (maze Maze) RemoveWalls(curr Position, neighb Position) {
	//richtung finden
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

//Visualize a Cell as a grid image 3x3 representation
func (cell Cell) Visualize() {

}
