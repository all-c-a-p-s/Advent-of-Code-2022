package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
	"slices"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getCrate (stack []string, numToMove int) ([]string, string) {//basically the equivalent of .pop() from python
	crateIndex := len(stack) - numToMove
	crate := stack[crateIndex]
	stack = append(stack[:crateIndex], stack[crateIndex + 1:]...)//deletes element from slice

	return stack, crate
}

func move(numToMove int, stackFrom, stackTo *[]string) {

	if numToMove == 0 {		
		return
	}

	crate := ""//empty string for now

	*stackFrom, crate = getCrate(*stackFrom, numToMove)

	*stackTo = append(*stackTo, crate)
	
	move(numToMove - 1, stackFrom, stackTo)//recursive call, exits when numToMove is 0

}

func parseInstruction(instruction string, crateStacks *[][]string) {
	instructionParts := strings.Split(instruction, " ")

	numToMove, err := strconv.Atoi(instructionParts[1])
	check(err)

	moveFrom, err := strconv.Atoi(instructionParts[3])
	check(err)

	moveTo, err := strconv.Atoi(instructionParts[5])
	check(err)

	move(numToMove, &(*crateStacks)[moveFrom - 1], &(*crateStacks)[moveTo - 1])//-1 because input file is not 0-indexed
}

func readCrates() [][]string {
	input, err := os.Open("input.txt")
	check(err)
	defer input.Close()
	
	crateStacks := make([][]string, 9)

	for i := 0; i < 9;i++{
		crateStacks[i] = make([]string, 0)
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			line := scanner.Text()
			if line[1] == '1' {
				input.Seek(0, 0)//reset scanner to beginning of file
				break //line after all crates have been read in
			}
			if line[i * 4 + 1] != ' ' {//where i is the stack currently being read
				crateStacks[i] = append(crateStacks[i], string(line[i * 4 + 1]))
			}
		}
		slices.Reverse(crateStacks[i])//so that crates on top are at the end of the slices	

	}

	return crateStacks
	
}

func main() {
	crateStacks := readCrates()

	input, err := os.Open("input.txt")
	check(err)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {		
		line := scanner.Text()
		
		if len(line) == 0{
			continue //line 10 of input is blank
		}

		if line[0] == 'm' {//'move x from y to z'
			parseInstruction(line, &crateStacks)
		}
	}

	for i := 0; i < 9;i++{
		_, last := getCrate(crateStacks[i], 1)//answer can't have an empty column so this shouldn't ever panic
		fmt.Print(last)
	}

}