package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type gridMapQ struct {
	hl, r, c, dr, dc, n int
}

type PriorityQueue []*gridMapQ

func (piq PriorityQueue) Len() int {
	return len(piq)
}

func (piq PriorityQueue) Less(i, j int) bool {
	return piq[i].hl < piq[j].hl
}

func (piq PriorityQueue) Swap(i, j int) {
	piq[i], piq[j] = piq[j], piq[i]

}
func (piq *PriorityQueue) Push(x interface{}) {
	item := x.(*gridMapQ)
	*piq = append(*piq, item)
}

func (piq *PriorityQueue) Pop() interface{} {
	old := *piq
	n := len(old)
	item := old[n-1]
	*piq = old[0 : n-1]
	return item
}

type seenTracker struct {
	r, c, dr, dc, n int
}

func main() {
	dat, e := os.ReadFile("./input.txt")
	check(e)
	lines := strings.Split(string(dat), "\n")

	var seen = make(map[seenTracker]int, 0)
	piq := PriorityQueue{}
	p := gridMapQ{0, 0, 0, 0, 0, 0}
	heap.Push(&piq, &p)

	for piq.Len() > 0 {
		g := heap.Pop(&piq).(*gridMapQ)
		//fmt.Println(g, len(lines), len(lines[0]))
		if g.r == len(lines)-2 && g.c == len(lines[0])-1 && g.n >= 4 {
			fmt.Println(g.r, g.c, "cost: ", g.hl)
			break
		}

		s := seenTracker{g.r, g.c, g.dr, g.dc, g.n}
		if _, ok := seen[s]; ok == true {
			continue
		}
		seen[s] = 1

		if g.n < 10 && !(g.dr == 0 && g.dc == 0) {
			nr := g.r + g.dr
			nc := g.c + g.dc
			if 0 <= nr && nr < len(lines)-1 && 0 <= nc && nc < len(lines[0]) {
				cost, _ := strconv.Atoi(string((lines[nr][nc])))
				q := gridMapQ{g.hl + cost, nr, nc, g.dr, g.dc, g.n + 1}
				heap.Push(&piq, &q)
			}
		}

		nextSteps := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		if g.n >= 4 || (g.dr == 0 && g.dc == 0) {
			for _, v := range nextSteps {
				if !(v[0] == g.dr && v[1] == g.dc) && !(v[0] == -g.dr && v[1] == -g.dc) {
					nr := g.r + v[0]
					nc := g.c + v[1]
					if 0 <= nr && nr < len(lines)-1 && 0 <= nc && nc < len(lines[0]) {
						cost, _ := strconv.Atoi(string((lines[nr][nc])))
						q := gridMapQ{g.hl + cost, nr, nc, v[0], v[1], 1}
						heap.Push(&piq, &q)
					}
				}
			}
		}

	}

	fmt.Println("vim-go")
}
