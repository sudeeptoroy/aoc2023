package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type subset struct {
	blue, red, green int
}

type game struct {
	gameid  int
	subsets []subset
}

func main() {
	dat, e := os.ReadFile("input.txt")
	if e != nil {
		panic(e)
	}
	games := []game{}
	lines := strings.Split(string(dat), "\n")
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		g := game{}
		g.gameid = i + 1
		parts := strings.Split(line, ": ")
		part2 := strings.Split(parts[1], ";")
		for _, ga := range part2 {
			oneGame := strings.Split(ga, ",")
			s := subset{}
			for _, color := range oneGame {
				c := strings.Fields(color)
				if c[1] == "blue" {
					s.blue, _ = strconv.Atoi(c[0])
				} else if c[1] == "red" {
					s.red, _ = strconv.Atoi(c[0])
				} else if c[1] == "green" {
					s.green, _ = strconv.Atoi(c[0])
				} else {
					fmt.Println(c)
				}
			}
			g.subsets = append(g.subsets, s)
		}
		games = append(games, g)
	}

	/*
		green := 13
		blue := 14
		red := 12
		flag := false
	*/
	total := 0
	mgreen := 0
	mblue := 0
	mred := 0

	for _, g := range games {
		//flag = false
		mgreen = 0
		mblue = 0
		mred = 0

		for _, s := range g.subsets {
			mgreen = max(mgreen, s.green)
			mblue = max(mblue, s.blue)
			mred = max(mred, s.red)
			/*
					if s.red < 0 || s.red > red {
						flag = true
						break
					}
					if s.blue < 0 || s.blue > blue {
						flag = true
						break
					}
					if s.green < 0 || s.green > green {
						flag = true
						break
					}
				}
				if flag == true {
					continue
				}
			*/
		}
		total += mgreen * mred * mblue
	}
	fmt.Println(total)
	fmt.Println("vim-go")
}
