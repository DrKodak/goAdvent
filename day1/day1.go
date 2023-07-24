package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("calories.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	max := 0
	val := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()	
		if line != "" {
			tmp, _ := strconv.Atoi(line)
			// check(err)
			val += tmp
			if val > max {
				max = val
			}
		} else {
			// fmt.Println("hello new line")
			val = 0	
		}
	}

	readFile.Close()
	fmt.Println(max)

	part2()
}