package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	resultDayOne := 0
	resultDayTwo := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("unable to open input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.Atoi(line)
		if err != nil {
			log.Panic("error converting string to number")
		}

		// Day One
		resultDayOne += value/3 - 2

		// Day Two
		result := value
		for result > 0 {
			result = result/3 - 2
			log.Println("Next", result)

			if result > 0 {
				resultDayTwo += result
			}
		}
	}

	log.Println("Day One", resultDayOne)
	log.Println("Day Two", resultDayTwo)
}
