//roughly 1.8x slower than that just loops through string window, but more concise
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

func findMarker(signal string) int {

	for i := 14;i < len(signal);i++{
		hashSet := map[byte]int{} //second data type doesn't really matter
		for j := 0;j < 14;j++{
			if _, ok := hashSet[signal[i - 14 + j]]; ok{
				break
			}
			hashSet[signal[i - 14 + j]] = j
		}
		if len(hashSet) == 14{
			return i
		}
	}
	panic("Could not find marker")
}

func main(){
	input_string := readFile()
	fmt.Println(findMarker(input_string))
}
