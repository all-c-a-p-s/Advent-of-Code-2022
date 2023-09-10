package main

import (
	"fmt"
	"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func readFile() string {
	file, err := os.ReadFile("input.txt")
	check(err)
	return string(file)
}

func areDifferent(signal string, letter int) bool {
	for i := 0;i < 14;i++ {
		for j := 0;j < 14;j++{
			if i == j {//otherwise letters at the same index will be compared
				continue
			}
			if signal[letter - i] == signal[letter - j]{
				return false
			}
		}
	}
	return true
}

func findMarker(signal string) int {
	for i := 14;i < len(signal);i++{
		if areDifferent(signal, i) {
			return i + 1//indexing starts at 0
		}
	}
	panic("Could not find marker")
}

func main(){
	input := readFile()
	fmt.Print(findMarker(input))
}
