package main

import (
	"fmt"
	"os"
)

// point 坐标点
type point struct {
	i, j int
}

func (p point) at(grid [][]int) (int, bool) {
	//在行上越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	// 在列上越界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

// dirs 探索方向：上左下右
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func readMaze(filename string) [][]int {
	var row, col int
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Fscanf(file, "%d %d", &row, &col)
	// fmt.Printf("row:%d,col:%d\n", row, col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func walk(maze [][]int, start, end point) [][]int {
	fmt.Printf("start:%v,end:%v\n", start, end)
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Quene := []point{start}
	for len(Quene) > 0 {
		current := Quene[0]
		Quene = Quene[1:]

		//探索到终点退出
		if current == end {
			break
		}

		//从上左下右依次探索
		for _, dir := range dirs {
			next := point{dir.i + current.i, dir.j + current.j}

			//可以执行下去的三个条件
			//next at maze is 0
			//next at steps is 0
			//and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			// 把探索到的点的填入steps
			curStep, _ := current.at(steps)
			steps[next.i][next.j] = curStep + 1
			// 把探索到的点的坐标添加入Quene
			Quene = append(Quene, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze/maze.txt")
	fmt.Println("maze")
	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%2d ", col)
		}
		fmt.Println()
	}
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Println("steps")
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
