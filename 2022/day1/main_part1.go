package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Println("Error reading file")
	}
	defer file.Close()

	type ElfInv struct {
		Count int32
		Total int32
		Items []int32
	}

	var elves []ElfInv
	var total int32
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
			total++
			current = ElfInv{}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Println("error reading file %b\n", err)
	}

	var highest int32

	for _, elf := range elves {
		if elf.Total > highest {
			highest = elf.Total
		}
	}

	log.Println("Total Elves", total)
	log.Println("Highest", highest)
}
