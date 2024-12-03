package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Panicln("error reading input file")
	}

	var data [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, strings.Split(line, " "))
	}

	partOne(data)
	partTwo(data)
}

func partOne(data [][]string) {
	var result int

	for _, reports := range data {
		for i, report := range reports {
			currentValue, _ := strconv.Atoi(report)

			var prevValue int
			var nextValue int
			var prevDist int
			var nextDist int

			if i != 0 {
				prevValue, _ = strconv.Atoi(reports[i-1])
				prevDist = prevValue - currentValue
			}

			if i < len(reports)-1 {
				nextValue, _ = strconv.Atoi(reports[i+1])
				nextDist = currentValue - nextValue
			} else {
				result++
				break // Finished
			}

			// Bad report if no change
			if nextDist == 0 {
				break
			}

			// Bad report if change greated than 3
			if nextDist > 3 || nextDist < -3 {
				break
			}

			// Bad report if trend changes
			if prevDist > 0 && nextDist < 0 || prevDist < 0 && nextDist > 0 {
				break
			}
		}
	}
	log.Println("Part One Results", result)
}

func partTwo(data [][]string) {
	var result int

	for _, reports := range data {
		var reportsToCheck []string

		for i := -1; i < len(reports)+1; i++ {

			if i >= 0 && i < len(reports) {
				reportsToCheck = append(reportsToCheck[:i], reportsToCheck[i+1:]...)
				log.Println("We should have removed one report here", reportsToCheck)
				// Loop through reports here
				reportsToCheck = reports
				log.Println("Reports reset", reportsToCheck)
			}

			log.Println("Weve checked all of the reports in the given (modifed) list")
		}
		log.Println("We have checked all report list permutations")
	}

	log.Println("Part Two Result", result)
}

func isReportSafe(prev int, current int, next int) bool {
	return false
}
