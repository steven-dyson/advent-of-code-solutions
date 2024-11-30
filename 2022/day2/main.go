package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Println("Error reading file")
	}
	defer file.Close()

	var total int32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		guide := scanner.Text()
		choice := string(guide[0])
		counter := string(guide[2])

		// Determine if player won
		winPoints := getWinPoints(choice, counter)
		total += winPoints

		// Determine player getHandScore
		handPoints := getHandPoints(counter)
		total += handPoints

		log.Println("Result", total)
	}
}

func getWinPoints(choice string, counter string) int32 {
	playerHand := getHandPoints(counter)
	oppHand := getHandPoints(choice)

	if playerHand == 1 && oppHand == 3 {
		return 6
	} else if playerHand == 2 && oppHand == 1 {
		return 6
	} else if playerHand == 3 && oppHand == 2 {
		return 6
	} else if playerHand == oppHand {
		return 3
	} else {
		return 0
	}
}

func getHandPoints(value string) int32 {
	switch value {
	case "A":
		return 1
	case "X":
		return 1
	case "B":
		return 2
	case "Y":
		return 2
	case "C":
		return 3
	case "Z":
		return 3
	}
	return 0
}

func resultDayOne() {
	log.Println("Result Day Onej")
}

func resultDayTwo() {
	log.Println("Day Two")
}
