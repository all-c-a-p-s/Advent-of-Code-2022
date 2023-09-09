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

func read_file() string {
	file, err := os.ReadFile("input.txt")
	check(err)
	return string(file)
}

func are_different(signal string, letter int) bool {
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

func find_marker(signal string) int {
	for i := 14;i < len(signal);i++{
		if are_different(signal, i) {
			return i + 1//indexing starts at 0
		}
	}
	return 0
}

func main(){
	input_string := read_file()
	fmt.Printf("%d", find_marker(input_string))
}