package main

import (
	"fmt"
	"os"
	"time"
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

func find_marker(signal string) int {

	for i := 14;i < len(signal);i++{
		hashSet := map[byte]int{} //data types don't really matter
		for j := 0;j < 14;j++{
			hashSet[signal[i - 14 + j]] = j
		}
		if len(hashSet) == 14{
			return i
		}
	}
	panic("Could not find marker")
}

func main(){
	start := time.Now()
	input_string := read_file()
	fmt.Println(find_marker(input_string))
	duration := time.Since(start)
	fmt.Println(duration)
}
