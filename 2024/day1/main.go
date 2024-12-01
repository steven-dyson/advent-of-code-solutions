package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Println("error reading file")
	}

	scanner := bufio.NewScanner(file)
	var colOne []int64
	var colTwo []int64

	for scanner.Scan() {
		line := scanner.Text()

		cols := strings.Split(line, "   ")

		colOne = append(colOne, toInt(cols[0]))
		colTwo = append(colTwo, toInt(cols[1]))
	}

	// Part One
	partOne(colOne, colTwo)

	// Part Two
	partTwo(colOne, colTwo)
}

func toInt(value string) int64 {
	new, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println("error parsing number")
		return 0
	}
	return new
}

func partOne(colOne []int64, colTwo []int64) {
	sort.Slice(colOne, func(i, j int) bool {
		return colOne[i] < colOne[j]
	})

	sort.Slice(colTwo, func(i, j int) bool {
		return colTwo[i] < colTwo[j]
	})

	var result int64

	for i := 0; i < len(colOne); i++ {
		distance := colOne[i] - colTwo[i]
		if distance < 0 {
			distance = distance * -1
		}
		result += distance
	}

	log.Println("Day One Result", result)
}

func partTwo(colOne []int64, colTwo []int64) {
	var result int64

	for _, oneValue := range colOne {
		var foundTimes int64

		for _, twoValue := range colTwo {
			if oneValue == twoValue {
				foundTimes++
			}
		}

		result += int64(oneValue) * foundTimes
	}

	log.Println("Part Two Result", result)
}
