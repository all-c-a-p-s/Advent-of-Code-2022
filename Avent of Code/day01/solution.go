package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)

func main(){

	input, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	defer input.Close()

	highestCalories := 0
	secondCalories := 0
	thirdCalories := 0
	elfCalories := 0//calories of current elf

	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			if elfCalories > highestCalories {//if it is equal it will pass next conidition an become secondCalories
				thirdCalories = secondCalories
				secondCalories = highestCalories
				highestCalories = elfCalories
	
			} else if elfCalories > secondCalories {
				thirdCalories = secondCalories
				secondCalories = elfCalories
	
			} else if elfCalories > thirdCalories {
				thirdCalories = elfCalories
			}
			elfCalories = 0
			continue
		}

		lineCalories, e := strconv.Atoi(line)

		if e != nil {
			log.Fatal(e)
		}

		elfCalories += lineCalories
		
		
			

	}

	fmt.Printf("highest: %d, second: %d, third: %d", highestCalories, secondCalories, thirdCalories)
}
