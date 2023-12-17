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

type coordinate struct {
	r, c, dr, dc int
}

func pop(q []coordinate) (coordinate, []coordinate) {
	c := q[0]
	if len(q) <= 1 {
		if len(q) == 0 {
			c = coordinate{}
		}
		retQ := make([]coordinate, 0)
		return c, retQ
	}
	return c, q[1:]
}

func push(q []coordinate, c coordinate) []coordinate {
	return append(q, c)
}

func getNextNode(cn coordinate, g []string) []coordinate {

	n1 := cn
	n2 := cn
	retTwo := false
	//r := cn.r + cn.dr
	//c := cn.c + cn.dc

	n1.r = cn.r + cn.dr
	n1.c = cn.c + cn.dc
	n2.r = n1.r
	n2.c = n1.c

	if isNodeValid(n1, len(g), len(g[0])) == false {
		return nil
	}
	//fmt.Println(len(g), len(g[0]))

	ch := g[n1.r][n1.c]
	//fmt.Printf("for: %d, %d %s %s", n1.r, n1.c, string(ch), " --> ")

	switch ch {
	case '/':
		n1.dr = -cn.dc
		n1.dc = -cn.dr
		break
	case '\\':
		n1.dr = cn.dc
		n1.dc = cn.dr
		break
	case '|':
		n1.dr = 1
		n1.dc = 0
		n2.dr = -1
		n2.dc = 0
		retTwo = true
		break
	case '-':
		n1.dr = 0
		n1.dc = 1
		n2.dr = 0
		n2.dc = -1
		retTwo = true
		break
	default:
		break
	}

	/*
		fmt.Printf("next nodes: ")
		n11 := n1
		n11.r = n1.r + n1.dr
		n11.c = n1.c + n1.dc
		if isNodeValid(n11, len(g), len(g[0])) == true {
			fmt.Printf(" %s (%d, %d) ", string(g[n11.r][n11.c]), n11.r, n11.c)
		}
	*/
	ret := []coordinate{n1}
	if retTwo == true {
		/*
			n3 := n2
			n3.r = n2.r + n2.dr
			n3.c = n2.c + n2.dc
			if isNodeValid(n3, len(g), len(g[0])) == true {
				fmt.Printf(" and %s (%d, %d)\n", string(g[n3.r][n3.c]), n3.r, n3.c)
			}
		*/
		ret = append(ret, n2)
	}

	return ret

}

func isNodeValid(n coordinate, r, c int) bool {
	if n.r >= r-1 || n.r < 0 || n.c >= c || n.c < 0 {
		return false
	}
	return true
}

type xyStore struct {
	r, c int
}

func findCount(lines []string, c coordinate) int {

	q := []coordinate{}
	visited := make(map[coordinate]int, 100000)
	col := len(lines[0])
	row := len(lines)
	seenBlocks := make(map[xyStore]int, 100000)

	q = push(q, c)
	n := coordinate{}
	for len(q) > 0 {
		n, q = pop(q)
		nextNodes := getNextNode(n, lines)
		for _, t := range nextNodes {
			if isNodeValid(t, row, col) == false {
				continue
			}
			if _, ok := visited[t]; ok == true {
				visited[t] += 1
				continue
			}
			xy := xyStore{t.r, t.c}
			seenBlocks[xy] = 1
			visited[t] = 1
			q = push(q, t)
		}
	}

	//fmt.Println(seenBlocks, len(seenBlocks))

	count := 0
	for _, _ = range seenBlocks {
		count++
	}
	return count
}

func main() {
	dat, e := os.ReadFile("./input.txt")
	check(e)

	lines := strings.Split(string(dat), "\n")

	c := coordinate{0, -1, 0, 1}
	ans := findCount(lines, c)

	max_val := 0
	for r := 0; r < (len(lines) - 1); r++ {
		c1 := coordinate{r, -1, 0, 1}
		v1 := findCount(lines, c1)
		c2 := coordinate{r, len(lines[0]), 0, -1}
		v2 := findCount(lines, c2)
		max_val = max(max_val, v1)
		max_val = max(max_val, v2)
	}

	for col := 0; col < (len(lines[0])); col++ {
		c1 := coordinate{-1, col, 1, 0}
		v1 := findCount(lines, c1)
		c2 := coordinate{len(lines[0]), col, -1, 0}
		v2 := findCount(lines, c2)
		max_val = max(max_val, v1)
		max_val = max(max_val, v2)
	}

	fmt.Println(ans)
	fmt.Println("max = ", max_val)

	fmt.Println("vim-go")
}
