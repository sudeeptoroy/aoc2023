package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTranspose(g []string) []string {

	col := len(g[0])
	row := len(g)
	if len(g[row-1]) == 0 {
		row--
	}

	newG := make([][]string, col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			newG[j] = append(newG[j], string(g[i][j]))
		}
	}

	retG := make([]string, 0)

	for i := 0; i < col; i++ {
		s := ""
		for j := 0; j < len(newG[i]); j++ {
			s += newG[i][j]
		}
		retG = append(retG, s)
	}

	//fmt.Println(retG)

	//fmt.Println(g, newG)
	return retG

}

func findMirror(g []string) (int, bool) {
	//col := len(g[0])
	row := len(g)
	if len(g[row-1]) == 0 {
		row--
	}
	//fmt.Println("grids are:", g, col, row)

	// for i := 0; i < row; i++ {
	// 	fmt.Println(g[i])
	// }

	mirrorFound := true
	left := 0
	right := 0
	//prevIndex := -1
	//finalRet := true
	mirrorLine := make([][]int, row)

	for r, line := range g {
		if len(line) == 0 {
			continue
		}
		for i := 0; i < len(line); i++ {
			mirrorFound = true

			reachedEnd := false
			for j := 0; j < len(line); j++ {
				left = i - j
				right = i + j + 1
				if left < 0 || right >= len(line) {
					break
				}
				//fmt.Println("checking i=", i, left, right, string(byte(line[left])), " ", string(byte(line[right])))
				if line[left] != line[right] {
					mirrorFound = false
				}
				if (left == 0 && right != (len(line)-1)) || ((left > 0) && (right == len(line)-1)) {
					reachedEnd = true
				}
			}
			if mirrorFound == true && reachedEnd == true {
				mirrorLine[r] = append(mirrorLine[r], i)
			}
		}
		//fmt.Println("for row, index = ", line, mirrorLine[r], mirrorFound)
	}
	count_map := map[int]int{}
	count := 0
	for _, oneLine := range mirrorLine {
		for _, k := range oneLine {
			if _, ok := count_map[k]; ok == true {
				continue
			}
			for _, l := range mirrorLine {
				for _, k1 := range l {
					if k == k1 {
						count++
						break
					}
				}
			}
			count_map[k] = count
			count = 0
		}
	}

	match_col := -1
	mismatchAllowed := 1
	for k, v := range count_map {
		if v == (row - mismatchAllowed) {
			match_col = k
		}
	}

	fmt.Println("returning:", match_col, count_map)
	return match_col, match_col != -1

}

func main() {
	dat, e := os.ReadFile("./input.txt")
	check(e)

	lines := strings.Split(string(dat), "\n\n")

	ans := 0
	for i, grid := range lines {
		oneGrid := strings.Split(grid, "\n")
		fmt.Println("grid no: ", i)

		v, result := findMirror(oneGrid)
		if result == true {
			ans += (v + 1)
		} else {
			// transpose
			newG := getTranspose(oneGrid)
			//fmt.Println("transpose:\n", newG)
			v, result := findMirror(newG)
			if result == false {
				fmt.Println("nothing here<----")
				panic("oops not possible")
			} else {
				ans += (v + 1) * 100
			}
		}
	}

	fmt.Println("ans: ", ans)

	fmt.Println("vim-go")
}
