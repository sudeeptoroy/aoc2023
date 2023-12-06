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

var (
	config_list_keys [1]string = [1]string{"seeds"}
	config_map_keys  [7]string = [7]string{"seed-to-soil map", "soil-to-fertilizer map", "fertilizer-to-water map", "water-to-light map", "light-to-temperature map", "temperature-to-humidity map", "humidity-to-location map"}
	config_map_name  [7]string = [7]string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
)

type maps_field struct {
	dest      int
	src       int
	range_len int
}

type config_maps struct {
	maps_fields []maps_field
}

func seed_to_location(seeds []int, amap []config_maps) int {

	min_location := 0
	temp_input := -1
	temp_out := 0
	for _, s := range seeds {
		if s == 0 {
			fmt.Println("skip")
			continue
		}

		temp_input = s
		temp_out = -1

		for _, m := range amap {
			for _, fields := range m.maps_fields {
				if temp_input >= fields.src && temp_input <= fields.src+fields.range_len {
					temp_out = fields.dest + (temp_input - fields.src)
					break
				}
			}
			if temp_out != -1 {
				temp_input = temp_out
			}
		}
		if min_location == 0 {
			min_location = temp_out
		} else if min_location > temp_out {
			min_location = temp_out
		}
		//fmt.Println("for seed ", s, "loc is", temp_out, "min loc = ", min_location)
	}
	return min_location
}

func parse_config(lines []string, seeds []int, parsed_config_map []config_maps) {

	conf_key_found := 0
	conf_map_index := -1
	go_to_next_line := 0
	var map_line []string
	var map_line_index = -1

	for _, line := range lines {
		if line == "\n" || line == "" {
			conf_key_found = 0
			conf_map_index = -1
			map_line_index = -1
			continue
		}

		if conf_key_found == 0 {
			conf := strings.Split(line, ":")
			for _, j := range config_list_keys {
				if conf[0] == j {
					map_line = strings.Fields(conf[1])
					for i, s := range map_line {
						seeds[i], _ = strconv.Atoi(string(s))
					}
				}
			}

			for i, j := range config_map_keys {
				if conf[0] == j {
					conf_key_found = 1
					conf_map_index = i
					map_line_index = 0
					go_to_next_line = 1
					break
				} else {
					conf_key_found = 0
					conf_map_index = -1
					map_line_index = -1
					go_to_next_line = 0
				}
			}
		}

		if go_to_next_line == 1 {
			go_to_next_line = 0
			continue
		}

		if conf_key_found == 1 {
			if conf_map_index != -1 {
				map_line = strings.Fields(line)
				parsed_config_map[conf_map_index].maps_fields[map_line_index].dest, _ = strconv.Atoi(map_line[0])
				parsed_config_map[conf_map_index].maps_fields[map_line_index].src, _ = strconv.Atoi(map_line[1])
				parsed_config_map[conf_map_index].maps_fields[map_line_index].range_len, _ = strconv.Atoi(map_line[2])
				map_line_index++
			}
		}
	}
}
func main() {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	default_maps := 10
	lines := strings.Split(string(dat), "\n")
	almanac_map := make([]config_maps, default_maps)
	for i := 0; i < default_maps; i++ {
		almanac_map[i].maps_fields = make([]maps_field, 200)
	}

	input_seeds := make([]int, 100000)
	seeds := make([]int, 1)

	parse_config(lines, input_seeds, almanac_map)

	//seeds_index := 0
	min := 0
	for i := 0; i < len(input_seeds); i += 2 {
		start := input_seeds[i]
		range_v := input_seeds[i+1]
		if start == 0 || range_v == 0 {
			continue
		}
		fmt.Println("between ", start, "and range ", range_v, " start min ", min)
		for j := 0; j <= range_v; j++ {
			seeds[0] = start + j
			ret := seed_to_location(seeds, almanac_map)
			if min == 0 {
				min = ret
			} else if ret < min {
				min = ret
			}
			//seeds[seeds_index] = start + j
			//seeds_index += 1
		}
		fmt.Println("between ", start, "and range ", range_v, " final min ", min)
	}

	fmt.Println("final min = ", min)

	//	seed_to_location(seeds, almanac_map)

	fmt.Println("vim-go")
}
