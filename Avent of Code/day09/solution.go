package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	vertical, horizontal int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func moveHead(line string, headPos *position) {
	insructions := strings.Fields(line)

	switch insructions[0] { //update head node position based on instruction
	case "R":
		(*headPos).horizontal += 1
	case "L":
		(*headPos).horizontal -= 1
	case "U":
		(*headPos).vertical += 1	
	case "D":
		(*headPos).vertical -= 1	
	default:
		panic("Invalid instrucion")
	}
}

func moveTail(headPos, tailPos *position, positionsVisited *map[position]bool, tailNode bool) {
	var relativePos = position {
		vertical: (*headPos).vertical - (*tailPos).vertical,
		horizontal: (*headPos).horizontal - (*tailPos).horizontal,
	}	

	vector := 5 * relativePos.vertical + relativePos.horizontal //essentially puts head and tail on a 5x5 grid with tail in the centre
	//this way head will always be on the grid. I did this to simplify if/else logic

	switch vector {
	case 12, 11, 7: //up and to the right - vector of 12 is for if a rope that was a 1-1 diagonal away moves by 1-1 because of the node in front of it
		(*tailPos).horizontal ++
		(*tailPos).vertical ++
	case 8, 9, 3: //up and left
		(*tailPos).horizontal --
		(*tailPos).vertical ++
	case -8, -9, -3: //down and right
		(*tailPos).horizontal ++
		(*tailPos).vertical --
	case -12, -11, -7: //down and left
		(*tailPos).horizontal --
		(*tailPos).vertical --

	case 10: //2 up
		(*tailPos).vertical ++
	case -10: //2 down
		(*tailPos).vertical --
	case 2: //2 right
		(*tailPos).horizontal ++
	case -2: //2 left
		(*tailPos).horizontal --
	}

	if tailNode{ //update hashmap only for the tail node
		(*positionsVisited)[*tailPos] = true
	}
}

func main() {
	input, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(input)

	positionsVisited := map[position]bool{} //hashmap used so that only unique values are counted
	positionsVisited[position{0, 0}] = true

	headPos := position {vertical: 0, horizontal: 0}
	node1 := position {vertical: 0, horizontal: 0}
	node2 := position {vertical: 0, horizontal: 0}
	node3 := position {vertical: 0, horizontal: 0}
	node4 := position {vertical: 0, horizontal: 0}
	node5 := position {vertical: 0, horizontal: 0}
	node6 := position {vertical: 0, horizontal: 0}
	node7 := position {vertical: 0, horizontal: 0}
	node8 := position {vertical: 0, horizontal: 0}
	tailPos := position {vertical: 0, horizontal: 0} //initialise rope nodes

	for scanner.Scan(){
		line := strings.Fields(scanner.Text())
		repeat, err := strconv.Atoi(line[1])
		check(err)

		for i := 0;i < repeat;i++{
			moveHead(scanner.Text(), &headPos)//moveHead() and moveTail() on all nodes
			moveTail(&headPos, &node1, &positionsVisited, false)
			moveTail(&node1, &node2, &positionsVisited, false)
			moveTail(&node2, &node3, &positionsVisited, false)
			moveTail(&node3, &node4, &positionsVisited, false)
			moveTail(&node4, &node5, &positionsVisited, false)
			moveTail(&node5, &node6, &positionsVisited, false)
			moveTail(&node6, &node7, &positionsVisited, false)
			moveTail(&node7, &node8, &positionsVisited, false)
			moveTail(&node8, &tailPos, &positionsVisited, true)
		}
	}

	fmt.Println(len(positionsVisited)) //length of hashmap = unique positions of tail node
}