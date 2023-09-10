package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"	
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func scenicScore(treeGrid [99][99]int, row, column int) int {
	leftVisible := 0
	rightVisible := 0
	upVisible := 0
	downVisible := 0 //initialised at 0 because they are incremented at the start of the loop

	for i := column + 1; i < 99;i++{
		rightVisible ++
		if treeGrid[row][i] >= treeGrid[row][column] {
			break
		}	
	} 

	for i := column - 1; i >= 0;i--{
		leftVisible ++
		if treeGrid[row][i] >= treeGrid[row][column] {
			break
		}
	}

	for i := row + 1; i < 99;i++{
		downVisible ++
		if treeGrid[i][column] >= treeGrid[row][column] {
			break
		}		
	}

	for i := row - 1; i >= 0;i--{
		upVisible ++
		if treeGrid[i][column] >= treeGrid[row][column] {
			break
		}
	}

	return leftVisible * rightVisible * upVisible * downVisible //will return 0 for trees on the edge, which I think is intended by the question
}

func main(){
	input, err := os.Open("input.txt")
	defer input.Close()
	check(err)	

	scanner := bufio.NewScanner(input)

	treeGrid := [99][99]int {} //works because trees are always arranged in a 99x99 square
	lineCount := 0

	for scanner.Scan() { //read input file into 2d array
		for i := 0;i < len(scanner.Text());i++{
			tree := string(scanner.Text()[i])
			treeGrid[lineCount][i], err = strconv.Atoi(tree)
			check(err)
		}
		lineCount ++
	}

	maxScore := 0

	for i := 0;i < 99;i++{ //loop through 2d array updating maxScore if a better scor eis found
		for j := 0;j < 99;j++{
			score := scenicScore(treeGrid, i, j)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println(maxScore)	
}