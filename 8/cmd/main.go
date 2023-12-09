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

type lrInstruct struct {
	left  string
	right string
}

func main() {
	dat, e := os.ReadFile("./input.txt")
	check(e)

	lines := strings.Split(string(dat), "\n")
	instructions := lines[0] + lines[1]
	node_map := map[string]lrInstruct{}

	for _, l := range lines[2:] {
		if len(l) == 0 {
			continue
		}

		ele := strings.Fields(l)
		//fmt.Println(ele)
		node := ele[0]
		le := ele[2][1:4]
		re := ele[3][:3]
		lrI := lrInstruct{
			left:  le,
			right: re,
		}
		node_map[node] = lrI

	}

	//fmt.Println(instructions, "\n", node_map)

	var XXAs []string

	for k, _ := range node_map {
		if k[2] == byte('A') {
			XXAs = append(XXAs, k)
		}
	}

	fmt.Println(XXAs)
	targetNodeCount := make([]int, len(XXAs))

	targetNode := ""
	for ni, n := range XXAs {
		targetNode = n
		oneStep := 0
		for i := 0; i < len(instructions); i++ {
			m := instructions[i]
			//fmt.Println("instructions = ", string(m))
			//fmt.Println("targetNode = ", targetNode)

			if string(m) == "L" {
				targetNode = node_map[targetNode].left
			} else {
				targetNode = node_map[targetNode].right
			}
			oneStep++
			if targetNode[2] == byte('Z') {
				targetNodeCount[ni] = oneStep
				break
			}
			if i == len(instructions)-1 {
				i = -1
			}
		}
	}
	fmt.Println("targetNodeCount = ", targetNodeCount)
	fmt.Println("now calculate the lcm online ^")
	fmt.Println("vim-go")
}
