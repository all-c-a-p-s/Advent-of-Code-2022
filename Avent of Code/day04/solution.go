package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"log"
)

func check(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func Overlap(pairs string) bool {
	firstRange := strings.Split(pairs, ",")[0]
	secondRange := strings.Split(pairs, ",")[1]

	firstMin := strings.Split(firstRange, "-")[0]
	firstMax := strings.Split(firstRange, "-")[1]

	secondMin := strings.Split(secondRange, "-")[0]
	secondMax := strings.Split(secondRange, "-")[1]

	min1, err1 := strconv.Atoi(firstMin)
	check(err1)

	max1, err2 := strconv.Atoi(firstMax)
	check(err2)

	min2, err3 := strconv.Atoi(secondMin)
	check(err3)

	max2, err4 := strconv.Atoi(secondMax)
	check(err4)

	if min2 <= max1 && min2 >= min1 {
		return true
	}

	if min1 <= max2 && min1 >= min2 {
		return true
	}

	return false

}

func main(){
	input, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(input)

	total := 0

	for scanner.Scan() {
		if Overlap(scanner.Text()) {
			total ++
		}

	}

	fmt.Println(total)

}