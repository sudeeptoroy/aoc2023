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

func get_copy_cards(mycards map[int]int, current int, nu int) {
	for i := current; i < current+nu; i++ {
		mycards[i+1] = mycards[i+1] + mycards[current]
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
	copy_cards := make(map[int]int)
	fmt.Printf("len cards = %d", len(cards))
	for i := 0; i < len(cards)-1; i++ {
		copy_cards[i] = 1
	}
	mycard_match := 0
	for i := 0; i < len(cards)-1; i++ {
		cno := strings.Split(cards[i], ": ")
		//fmt.Println(cno)
		nu := strings.Split(cno[1], " | ")
		//fmt.Println(nu[0])
		//fmt.Println(nu[1])
		sub_total = 0
		mycard_match = 0
		for _, w := range strings.Fields(nu[0]) {
			for _, n := range strings.Fields(nu[1]) {
				//fmt.Println(w, n)
				if w == n {
					mycard_match++
					//fmt.Printf("<-- match ")
					fmt.Println("i: ", i, w, n)
					if sub_total == 0 {
						sub_total = 1
					} else {
						sub_total = sub_total * 2
					}
					//fmt.Printf("subtotal = %d", sub_total)
				}
			}
		}
		total += sub_total
		get_copy_cards(copy_cards, i, mycard_match)
		fmt.Printf(" %d) subtotal = %d, total = %d, this card match = %d\n", i, sub_total, total, mycard_match)

		//fmt.Println(i)
		//fmt.Println(nu[0])
		//fmt.Println(nu[1])

	}
	fmt.Println(total)
	fmt.Println(copy_cards)
	total_count := 0
	for k, v := range copy_cards {
		total_count += v
		fmt.Printf("for key = %d, value = %d, total = %d\n", k, v, total_count)
	}
	fmt.Println("card total", total_count)

	return total
}
