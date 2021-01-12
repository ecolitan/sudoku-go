package main

import "fmt"

var grid [][]int

func main() {
	//easy
	/*grid = [][]int{
		{0,0,0,8,0,5,4,2,7},
		{0,0,7,2,0,3,9,0,5},
		{0,2,4,7,0,1,0,0,0},
		{7,4,0,0,8,6,1,0,0},
		{0,0,8,9,1,0,7,0,4},
		{0,1,0,0,3,0,0,9,8},
		{0,0,0,0,0,0,0,0,1},
		{9,0,1,3,7,0,0,0,6},
		{2,0,0,0,5,0,3,4,0},
	}*/

	//hard
	grid = [][]int{
		{0,8,0,4,7,9,0,0,2},
		{0,0,7,0,0,2,0,0,0},
		{0,0,0,5,0,0,0,6,4},
		{2,4,3,0,0,7,0,0,0},
		{8,0,0,9,1,0,0,0,0},
		{0,0,1,0,0,0,0,0,0},
		{6,0,0,3,0,5,0,0,0},
		{0,0,8,0,0,0,0,0,9},
		{3,7,0,0,9,1,2,0,0},
	}

	//many solutions
	/*grid = [][]int{
		{0,8,0,4,7,9,0,0,2},
		{0,0,7,0,0,2,0,0,0},
		{0,0,0,5,0,0,0,6,4},
		{2,4,3,0,0,7,0,0,0},
		{8,0,0,9,1,0,0,0,0},
		{0,0,1,0,0,0,0,0,0},
		{6,0,0,3,0,5,0,0,0},
		{0,0,8,0,0,0,0,0,9},
		{3,7,0,0,9,1,0,0,0},
	}*/

	PrintGrid()
	Solve()
}

func Solve() {
	var y,x,n int
	for y = 0; y < 9; y++ {
		for x = 0; x < 9; x++ {
			if grid[y][x] == 0 {
				for n = 1; n < 10; n++ {
					if Possible(y,x,n) == true {
						grid[y][x] = n
						Solve()
						grid[y][x] = 0
					}
				}
				return
			}
		}
	}
	PrintGrid()
	fmt.Print("\nPress enter to search for more solutions")
	fmt.Scanln()

	return
}

func Possible(y int, x int, n int) bool {
	var i,j,x0,y0 int

	for i = 0; i < 9; i++ {
		if grid[y][i] == n {
			return false
		}

	}
	for i = 0; i < 9; i++ {
		if grid[i][x] == n {
			return false
		}
	}

	y0 = (y/3)*3
	x0 = (x/3)*3
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if grid[y0+i][x0+j] == n {
				return false
			}
		}
	}

	return true
}

func PrintGrid() {
	fmt.Println()
	for _,v := range grid {
		fmt.Println(v)
	}
}