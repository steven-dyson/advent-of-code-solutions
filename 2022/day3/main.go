package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panicln("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	var rucksacks []string

	for scanner.Scan() {
		line := scanner.Text()
		rucksacks = append(rucksacks, line)
	}

	// Part One
	partOne(rucksacks)

	// Part partTwo
	partTwo(rucksacks)
}

func partOne(rucksacks []string) {
	var result int

	for _, rucksack := range rucksacks {
		half := len(rucksack) / 2

		first := rucksack[0:half]
		firstChars := strings.Split(first, "")
		second := rucksack[half:]
		secondChars := strings.Split(second, "")

		for _, char := range firstChars {
			found := false

			for _, secondChar := range secondChars {
				if char == secondChar {
					found = true
					rune := []rune(char)[0]
					charValue := charToInt(rune)
					result += charValue
					break
				}
			}

			if found {
				break
			}
		}
	}

	log.Println("Result Day One", result)
}

func partTwo(rucksacks []string) {
	var result int

	for i := 0; i < len(rucksacks)-2; i = i + 3 {
		rucksackOne := rucksacks[i]
		rucksackOneChars := strings.Split(rucksackOne, "")
		rucksackTwo := rucksacks[i+1]
		rucksackThree := rucksacks[i+2]

		for _, char := range rucksackOneChars {
			if strIncludes(rucksackTwo, char) && strIncludes(rucksackThree, char) {
				rune := []rune(char)[0]
				value := charToInt(rune)
				result += value
				break
			}
		}
	}

	log.Println("Result Day Two", result)
}

func strIncludes(str string, value string) bool {
	chars := strings.Split(str, "")

	for _, char := range chars {
		if value == char {
			return true
		}
	}
	return false
}

func charToInt(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char - 'a' + 1)
	} else if char >= 'A' && char <= 'Z' {
		return int(char - 'A' + 27)
	}
	return 0
}
