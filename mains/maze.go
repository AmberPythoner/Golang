package main

import (
	"fmt"
	"os"
)

func readFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d%d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	x, y int
}

func (p point) add(r point) point {
	return point{r.x + p.x, r.y + p.y}
}

func (p point) at(grids [][]int) (int, bool) {
	if p.x < 0 || p.x >= len(grids) {
		return 0, false
	}
	if p.y < 0 || p.y >= len(grids[p.x]) {
		return 0, false
	}
	return grids[p.x][p.y], true
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)
			value, ok := next.at(maze)
			if !ok || value == 1 {
				continue
			}
			value, ok = next.at(steps)
			if !ok || value != 0 {
				continue
			}
			if next == start {
				continue
			}
			currentSteps, _ := cur.at(steps)
			steps[next.x][next.y] = currentSteps + 1

			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readFile("mains/maze.in")
	for _, step := range maze {
		for _, value := range step {
			fmt.Printf("%3d ", value)
		}
		fmt.Println()
	}
	fmt.Println("********************************")
	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	for _, step := range steps {
		for _, value := range step {
			fmt.Printf("%3d  ", value)
		}
		fmt.Println()
	}
}
