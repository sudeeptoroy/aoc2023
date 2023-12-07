package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type HandType int

const (
	five_of_a_kind HandType = iota
	four_of_a_kind
	full_house
	three_of_a_kind
	two_pair
	one_pair
	high_card
)

var card_strength []rune //("AKQJT98765432")

type hand_bid struct {
	hand      string
	bid       int
	rank      int
	hand_type HandType
}

func get_card_hand(hand string) HandType {
	card_label_tracker := make([]int, 5)

	hand_rune := []rune(hand)

	fmt.Println(hand)

	sort.Slice(hand_rune, func(i, j int) bool {
		return hand_rune[i] < hand_rune[j]
	})

	hand = string(hand_rune)
	fmt.Println(hand)

	//
	card_map := make(map[rune]int)
	for _, c := range hand {
		card_map[c] = card_map[c] + 1
	}

	fmt.Println("map = ", card_map)

	if card_map['J'] != 0 && card_map['J'] != 5 {
		fmt.Println("going to replce cards -- nu ", card_map['J'])

		// Create slice of key-value pairs
		pairs := make([][2]interface{}, 0, len(card_map))
		for k, v := range card_map {
			pairs = append(pairs, [2]interface{}{k, v})
		}

		// Sort slice based on values
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i][1].(int) > pairs[j][1].(int)
		})

		// Extract sorted keys
		keys := make([]rune, len(pairs))
		for i, p := range pairs {
			keys[i] = p[0].(rune)
		}

		// Print sorted map
		for _, k := range keys {
			fmt.Printf("%s: %d\n", string(k), card_map[k])
		}

		fmt.Println("before replacement sorted map = ", card_map)
		//for k, _ := range card_map {
		for _, k := range keys {
			//fmt.Println("hand = ", hand)
			if k == 'J' {
				continue
			}
			fmt.Println("J --> ", string(k), " hand before replacement", hand, card_map['J'])
			hand = strings.Replace(hand, "J", string(k), -1)
			fmt.Println("hand = ", hand)
			break
		}
	}

	card_map1 := make(map[rune]int)
	for _, c := range hand {
		card_map1[c] = card_map1[c] + 1
	}
	fmt.Println("map1 = ", card_map1)
	// Create slice of key-value pairs
	pairs := make([][2]interface{}, 0, len(card_map1))
	for k, v := range card_map1 {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	// Sort slice based on values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1].(int) > pairs[j][1].(int)
	})

	// Extract sorted keys
	keys := make([]rune, len(pairs))
	for i, p := range pairs {
		keys[i] = p[0].(rune)
	}

	// Print sorted map
	for i, k := range keys {
		fmt.Printf("%s: %d\n", string(k), card_map1[k])
		card_label_tracker[i] = card_map1[k]
	}

	fmt.Println(card_label_tracker)
	if card_label_tracker[0] == 5 {
		return five_of_a_kind
	} else if card_label_tracker[0] == 4 {
		return four_of_a_kind
	} else if card_label_tracker[0] == 3 {
		if card_label_tracker[1] == 2 {
			return full_house
		} else {
			return three_of_a_kind
		}
	} else if card_label_tracker[0] == 2 {
		if card_label_tracker[1] == 2 {
			return two_pair
		} else {
			return one_pair
		}
	} else {
		return high_card
	}
}

func main() {
	card_strength := []rune("AKQJT98765432J")
	cards := make([]hand_bid, 1000)
	dat, e := os.ReadFile("./input.txt")
	check(e)

	lines := strings.Split(string(dat), "\n")
	//fmt.Println(lines)

	ci := 0
	for _, line := range lines {
		//fmt.Println(line)
		one_line := strings.Fields(line)
		if len(one_line) == 0 {
			fmt.Println("continue")
			continue
		}
		fmt.Println(one_line)
		cards[ci].hand = one_line[0]
		cards[ci].hand_type = get_card_hand(one_line[0])
		cards[ci].bid, _ = strconv.Atoi(one_line[1])
		ci++
	}

	fmt.Println("before:", cards)

	sort.Slice(cards, func(i, j int) bool {
		if cards[i].hand_type < cards[j].hand_type {
			return false
		} else if cards[i].hand_type > cards[j].hand_type {
			return true
		}
		var card_i [5]int
		var card_j [5]int
		for c, _ := range cards[i].hand {
			for a, b := range card_strength {
				if cards[i].hand[c] == byte(b) {
					card_i[c] = a
				}
				if cards[j].hand[c] == byte(b) {
					card_j[c] = a
				}
			}
		}
		for a, _ := range cards[i].hand {
			if card_i[a] < card_j[a] {
				return false
			} else if card_i[a] > card_j[a] {
				return true
			}
		}
		return true
	})

	fmt.Println("After:", cards)

	winning_total := 0
	for r, c := range cards {
		if c.bid == 0 {
			continue
		}
		winning_total = winning_total + c.bid*(r+1)
	}

	fmt.Printf("Winning total: %d", winning_total)

	fmt.Println("vim-go")
}
