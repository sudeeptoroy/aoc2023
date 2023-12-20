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

func main() {
	dat, e := os.ReadFile("./input.txt")
	check(e)
	lines := strings.Split(string(dat), "\n")

	type coo struct {
		r, c int64
	}

	/*
		dirMap := map[string]coo{
			"U": {-1, 0},
			"D": {1, 0},
			"L": {0, -1},
			"R": {0, 1},
		}
	*/
	dirMap2 := map[byte]coo{
		'3': {-1, 0},
		'1': {1, 0},
		'2': {0, -1},
		'0': {0, 1},
	}
	var b int64 = 0
	coordinatesMap := []coo{{0, 0}}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		l := strings.Fields(line)
		number := l[2][2:7]
		dir := byte(l[2][7])
		//fmt.Println(number, string(dir))
		//d := l[0]
		//n, _ := strconv.Atoi(string(l[1]))
		n, _ := strconv.ParseInt(number, 16, 64)
		//fmt.Println(n)

		//dx := dirMap[d]
		dx := dirMap2[dir]
		b += n
		sc := coordinatesMap[len(coordinatesMap)-1]
		nc := coo{sc.r + dx.r*n, sc.c + dx.c*n}
		coordinatesMap = append(coordinatesMap, nc)
	}

	//fmt.Println(coordinatesMap)
	//fmt.Println(len(coordinatesMap))
	//fmt.Println(b)

	var A int64 = 0
	for i, _ := range coordinatesMap {
		//
		var pc int64 = 0
		if i == 0 {
			pc = coordinatesMap[0].c
		} else {
			pc = coordinatesMap[i-1].c
		}
		var nc int64 = 0
		if i == len(coordinatesMap)-1 {
			nc = coordinatesMap[0].c
		} else {
			nc = coordinatesMap[i+1].c
		}
		v := coordinatesMap[i].r * (pc - nc)
		if v < 0 {
			//v = -v
		}
		A += v
	}

	A = (A / 2)
	if A < 0 {
		A = -A
	}

	i := A - b/2 + 1

	fmt.Println(i + b)

	fmt.Println("vim-go")
}
