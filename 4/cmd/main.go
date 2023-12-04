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

func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	cards := strings.Split(string(dat), "\n")

	//for _, c := range cards {
	//	fmt.Println(c)
	//}

	fmt.Println("first result:	", scratchcards1(cards))

	fmt.Println("vim-go")
}

func scratchcards1(cards []string) int {
	total := 0
	sub_total := 0
	for i := 0; i < len(cards)-1; i++ {
		cno := strings.Split(cards[i], ": ")
		//fmt.Println(cno)
		nu := strings.Split(cno[1], " | ")
		fmt.Println(nu[0])
		fmt.Println(nu[1])
		sub_total = 0
		for _, w := range strings.Fields(nu[0]) {
			for _, n := range strings.Fields(nu[1]) {
				//fmt.Println(w, n)
				if w == n {
					fmt.Printf("<-- match ")
					fmt.Println("i: ", i, w, n)
					if sub_total == 0 {
						sub_total = 1
					} else {
						sub_total = sub_total * 2
					}
					fmt.Printf("subtotal = %d", sub_total)
				} else {
					fmt.Printf("\n")
				}
			}
		}
		total += sub_total
		fmt.Printf(" %d) subtotal = %d, total = %d\n", i, sub_total, total)

		//fmt.Println(i)
		//fmt.Println(nu[0])
		//fmt.Println(nu[1])

	}
	fmt.Println(total)
	return total
}
