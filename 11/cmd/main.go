package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	y, x int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findDistance(a, b position, row, col []int, p int) int {
	dis := 0
	//fmt.Println("between pos = ", a, b)
	dis += abs((a.y - b.y))
	for _, r := range row {
		if r > min(a.y, b.y) && r < max(a.y, b.y) {
			dis += p
		}
	}

	dis += abs((a.x - b.x))
	for _, c := range col {
		if c > min(a.x, b.x) && c < max(a.x, b.x) {
			dis += p
		}
	}
	//fmt.Println("dis = ", dis)
	return dis
}

func main() {
	dat, e := os.ReadFile("./input.txt")
	check(e)

	lines := strings.Split(string(dat), "\n")
	X, Y := len(lines[0]), len(lines)-1
	fmt.Println(X, Y)

	// col
	var emptyCols []int
	for i := 0; i < X; i++ {
		isEmpty := true
		for j := 0; j < Y; j++ {
			if lines[j][i] == byte('#') {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyCols = append(emptyCols, i)
		}
	}

	// rows
	var emptyRows []int
	for i := 0; i < X; i++ {
		isEmpty := true
		for j := 0; j < Y; j++ {
			if lines[i][j] == byte('#') {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, i)
		}
	}

	fmt.Println(emptyCols, emptyRows)
	/*
		// create map
		m := make([][]byte, Y+len(emptyRows))
		for i := 0; i < X; i++ {
			m[i] = make([]byte, X+len(emptyCols))
		}
	*/

	galaxy := []position{}

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if lines[i][j] == byte('#') {
				galaxy = append(galaxy, position{y: i, x: j})
			}
		}
	}

	fmt.Println(galaxy)

	// min distaces

	var d int = 0
	for i, _ := range galaxy {
		l := len(galaxy)
		for j := i + 1; j < l; j++ {
			d += findDistance(galaxy[i], galaxy[j], emptyRows, emptyCols, (1000000 - 1))
		}
	}

	fmt.Println(d)

	fmt.Println("vim-go")
}
