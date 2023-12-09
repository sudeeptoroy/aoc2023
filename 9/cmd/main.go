package main

import (
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

func predict(line []int, forward bool) int {
	var diff []int
	allZeros := true

	for i := 0; i < len(line)-1; i++ {
		v := line[i+1] - line[i]
		if v != 0 {
			allZeros = false
		}
		diff = append(diff, v)
	}
	fmt.Println("diff ", diff)
	if forward {
		if allZeros {
			return line[len(line)-1]
		} else {
			return line[len(line)-1] + predict(diff, true)
		}
	} else {
		if allZeros {
			return line[0]
		} else {
			return line[0] - predict(diff, false)
		}
	}
}

func main() {
	dat, e := os.ReadFile("./input.txt")
	check(e)

	ansFor := 0
	ansBack := 0

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		ls := strings.Fields(line)
		if len(ls) == 0 {
			continue
		}
		li := make([]int, len(ls))
		for i, j := range ls {
			v, e := strconv.Atoi(j)
			check(e)
			li[i] = v
		}
		ansFor += predict(li, true)
		ansBack += predict(li, false)
	}

	fmt.Println("ans = ", ansFor, ansBack)

	fmt.Println("vim-go")
}
