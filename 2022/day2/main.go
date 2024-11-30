package main

import (
	"bufio"
	"log"
	"os"
)

type Hand struct {
	OppChoice HandName
	Target    HandName
}

type GameMap struct {
	Score      int32
	Defeats    HandName
	DefeatedBy HandName
}

type HandName string

const (
	Rock     HandName = "rock"
	Paper    HandName = "paper"
	Scissors HandName = "scissors"
)

const (
	WinPoints  = 6
	DrawPoints = 3
	LosePoints = 0
)

var gameMap = map[HandName]GameMap{
	"rock":     {Score: 1, Defeats: Scissors, DefeatedBy: Paper},
	"paper":    {Score: 2, Defeats: Rock, DefeatedBy: Scissors},
	"scissors": {Score: 3, Defeats: Paper, DefeatedBy: Rock},
}

func main() {
	var hands []Hand

	// Read the file input
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Println("Error reading file")
	}
	defer file.Close()

	// Scan each line in file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		guide := scanner.Text()
		oppChoice := HandName(guide[0])
		target := HandName(guide[2])

		// Add opponent choice and player target to hands
		hands = append(hands, Hand{oppChoice, target})
	}

	// Run day 1
	resultDayOne(hands)

	// Run day 2
	resultDayTwo(hands)
}

func toHandName(value string) HandName {
	if value == "A" || value == "X" {
		return Rock
	} else if value == "B" || value == "Y" {
		return Paper
	} else {
		return Scissors
	}
}

func resultDayOne(hands []Hand) {
	log.Println("Result Day One")

	var result int32

	for _, hand := range hands {
		// Get points from winning, losing, or drawing
		oppHand := toHandName(string(hand.OppChoice))
		targetHand := toHandName(string(hand.Target))
		if gameMap[oppHand].DefeatedBy == targetHand {
			result += WinPoints + gameMap[targetHand].Score
		} else if gameMap[oppHand].Defeats == targetHand {
			result += LosePoints + gameMap[targetHand].Score
		} else {
			result += DrawPoints + gameMap[targetHand].Score
		}
	}

	log.Println("Day One Result", result)
}

func resultDayTwo(hands []Hand) {
	var result int32

	for _, hand := range hands {
		if hand.Target == "X" {
			oppHand := toHandName(string(hand.OppChoice))
			playerHand := gameMap[oppHand].Defeats
			result += LosePoints + gameMap[playerHand].Score
		} else if hand.Target == "Y" {
			oppHand := toHandName(string(hand.OppChoice))
			playerHand := oppHand
			result += DrawPoints + gameMap[playerHand].Score
		} else {
			oppHand := toHandName(string(hand.OppChoice))
			playerHand := gameMap[oppHand].DefeatedBy
			result += WinPoints + gameMap[playerHand].Score
		}
	}

	log.Println("Day Two Result", result)
}
