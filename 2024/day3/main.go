package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Panicln("error reading input file")
	}

	log.Println("Test")
	var expressions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println("Line", line)

		chars := strings.Split(line, "")
		//var instruction string

		for i, char := range chars {
			log.Println("Char", char)
			isValid := getMul(chars[i:i+8])	

			if isValid {
				expressions = append(expressions, strings.Join(chars[i:i+12], ""))
			}
				
			log.Println("Valid", isValid)
		}
	}
	log.Println("EXPS", expressions)
}

func getMul(str []string) bool {
	if str[0] = "m" && str[1] == "u" && str[2] == "l" && str[3] == "(" {
		remaining := str[4:len(str)]
		for _, char := range str
	}
}

func isInteger(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

