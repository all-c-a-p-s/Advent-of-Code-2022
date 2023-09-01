package main

import (
	"fmt"
	"os"
	"bufio"
)

func findDuplicate(items, items2, items3 string) rune {

	firstRucksacks := []rune(items)
	secondRucksacks := []rune(items2)
	thirdRucksacks := []rune(items3)

	for i := 0; i < len(firstRucksacks); i++{
		for j := 0; j < len(secondRucksacks); j++{
			for k := 0; k < len(thirdRucksacks); k++{
				if firstRucksacks[i] == secondRucksacks[j] && secondRucksacks[j] == thirdRucksacks[k]{
					return firstRucksacks[i]
				}
			}

		}
	}

	fmt.Println("Could not find duplicate")
	return '#'

}

func Priority(item rune) int {
	ascii_code := int(item)
	
	if ascii_code > 96 {//lowercase
		return ascii_code - 96
	}

	return ascii_code - 38 //uppercase

}

func main(){

	input, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file")
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	total := 0

	lineCount := 0

	var firstRucksacks string
	var secondRucksacks string
	var thirdRucksacks string

	for scanner.Scan() {

		badgePriority := 0
		
		if lineCount % 3 == 0 {
			firstRucksacks = scanner.Text()
		} else if lineCount % 3 == 1 {
			secondRucksacks = scanner.Text()
		} else {
			thirdRucksacks = scanner.Text()
			Duplicate := findDuplicate(firstRucksacks, secondRucksacks, thirdRucksacks)
			badgePriority = Priority(Duplicate)
		}

		total += badgePriority
		lineCount ++
	}

	fmt.Println(total)
}