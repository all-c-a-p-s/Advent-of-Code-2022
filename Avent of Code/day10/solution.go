package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func updateDisplay(clockTicks, x int) { //doesn't actually recognise the letter, you need to be able to read them

	if clockTicks % 40 == 0 {//new row
		fmt.Println()
	}

	if (x - 1 == clockTicks % 40) || (x == clockTicks % 40) || (x + 1 == clockTicks % 40){
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

}

func main() {
	input, err := os.Open("input.txt")
	defer input.Close()

	check(err)

	scanner := bufio.NewScanner(input)

	x, clockTicks := 1, 0

	for scanner.Scan() {

		instruction := strings.Fields(scanner.Text())

		clockTicks ++
		updateDisplay(clockTicks, x)		

		if instruction[0] == "addx" {

			add, err := strconv.Atoi(instruction[1])
			check(err)

			x += add

			clockTicks ++

			updateDisplay(clockTicks, x)

		}
	}
}
