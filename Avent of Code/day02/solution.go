package main

import (
	"bufio"
	"fmt"
	"os"
)

func roundScore(round string) int {
	resultScore := make(map[string]int)

	resultScore["A X"] = 3
    resultScore["A Y"] = 4
	resultScore["A Z"] = 8
    resultScore["B X"] = 1
    resultScore["B Y"] = 5
    resultScore["B Z"] = 9
    resultScore["C X"] = 2
    resultScore["C Y"] = 6
    resultScore["C Z"] = 7

	return resultScore[round]
}

func main() {
	input, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	totalScore := 0

	for scanner.Scan(){
		totalScore += roundScore(scanner.Text())
	}

	fmt.Println(totalScore)
	
}