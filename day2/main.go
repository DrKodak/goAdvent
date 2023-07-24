package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func myGameScore(them string, me string) int {
	score := 0	
	switch me {
	case "ROCK":
		score += 1
		if them == "ROCK" {
			score += 3
		} else if them == "SCISSORS" {
			score += 6
		}
	case "PAPER":
		score += 2
		if them == "PAPER" {
			score +=3
		} else if them == "ROCK" {
			score += 6
		}
	case "SCISSORS":
		score += 3
		if them == "SCISSORS" {
			score += 3
		} else if them == "PAPER" {
			score +=6
		}	
	}
	return score	
}

func convertThem(them string) string {
	val := ""	
	switch them {
	case "A" :
		val = "ROCK"
	case "B":
		val = "PAPER"
	case "C":
		val = "SCISSORS"
	}
	return val
}

func convertMe(them string) string {
	val := ""	
	switch them {
	case "X" :
		val = "ROCK"
	case "Y":
		val = "PAPER"
	case "Z":
		val = "SCISSORS"
	}
	return val
}

func main() {

	// open file
	file, err := os.Open("strat.txt")
	check(err)
	fileScanner := bufio.NewScanner(file)
	score := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		them := convertThem(string(line[0]))
		me := convertMe(string(line[2]))
		// fmt.Println("them vs me, ", them, me)
		score += myGameScore(them, me)	
	}
	fmt.Println(score)
}