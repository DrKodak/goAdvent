package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func part2() {
	readFile, err := os.Open("calories.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	
	max1, max2, max3 := 0,0,0
	val := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			tmpInt, _ := strconv.Atoi(line)
			val += tmpInt
		} else {
			if val > max1 {
				max3 = max2
				max2 = max1 
				max1 = val
			} else if val > max2 {
				max3 = max2
				max2 = val
			} else if val > max3 {
				max3 = val
			}
			val = 0
		}	
	}
	fmt.Println(max1 + max2 + max3)

	readFile.Close()
}