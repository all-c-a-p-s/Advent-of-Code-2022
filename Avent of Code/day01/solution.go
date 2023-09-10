package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"sort"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main(){

	input, err := os.Open("input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	elves := []int{}
	elfCalories := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, elfCalories) //append current elf to elves{}
			elfCalories = 0
			continue
		} 
		calories, err := strconv.Atoi(scanner.Text())
		check(err)
		
		elfCalories += calories
	}

	sort.Slice(elves, func(i, j int) bool{ //sort by descending size
		return elves[i] > elves[j]
	})

	fmt.Println(elves[0] + elves[1] + elves[2]) //sum of first 3
}
