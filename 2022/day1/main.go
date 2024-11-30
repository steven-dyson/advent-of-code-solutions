package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type ElfInv struct {
	Count int32
	Total int32
	Items []int32
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Println("Error reading file")
	}
	defer file.Close()

	var elves []ElfInv
	var current ElfInv

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		string := scanner.Text()

		if string != "" {
			value, err := strconv.ParseInt(string, 10, 32)
			if err != nil {
				log.Println("error parsing string to number")
			}

			current.Total += int32(value)
			current.Count++
			current.Items = append(current.Items, int32(value))
		} else {
			elves = append(elves, current)
			current = ElfInv{}
		}
	}

	// Handle last elf
	if current.Count > 0 {
		elves = append(elves, current)
	}

	// Run part one result
	partOneResult(elves)

	// RUn part two result
	partTwoResult(elves)

	if err := scanner.Err(); err != nil {
		log.Println("error reading file %b\n", err)
	}
}

func partOneResult(elves []ElfInv) {
	var highest int32

	for _, elf := range elves {
		if elf.Total > highest {
			highest = elf.Total
		}
	}

	log.Println("Part One", highest)
}

func partTwoResult(elves []ElfInv) {
	total := int32(0)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Total > elves[j].Total
	})

	for i := 0; i < 3 && i < len(elves); i++ {
		total += elves[i].Total
	}

	log.Println("Part Two", total)
}
