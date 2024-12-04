package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panicln("error reading input file")
	}

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			firstValue, _ := strconv.Atoi(match[1])
			secondValue, _ := strconv.Atoi(match[2])
			result += firstValue * secondValue // expressions = append(expressions, match)
		}
	}
	log.Println("Day One Result", result)
}

func isInteger(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
