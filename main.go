package main

import (
	"image/png"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	dx, err1 := strconv.Atoi(args[0])
	dy, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		panic("Please provide exact two numeric arguments as dimension")
	}

	maze := CreateMaze(dx, dy)

	GeneratePathDeepFirst(maze)

	image := maze.Visualize()

	file, _ := os.Create("out.png")
	png.Encode(file, image)
}
