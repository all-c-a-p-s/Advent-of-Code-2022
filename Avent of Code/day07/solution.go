package main 

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
)

type directory struct {
	size int//directories will be read in with size 0, files will have their size read in
	name string
	parent *directory
	children map[string]*directory //uses directory name as key
}

func check(err error) {
	if err != nil{
		log.Fatal(err)
	}
}

func directorySize(currentDirectory *directory) (size int) {
	if len((*currentDirectory).children) == 0{//file
		return currentDirectory.size
	} 

	for _, childDirectory := range (*currentDirectory).children {
		size += directorySize(childDirectory)//recursive call on children of this directory
	}

	return size
}

func main() {
	input, err := os.Open("input.txt")
	defer input.Close()
	check(err)

	scanner := bufio.NewScanner(input)

	var homeDirectory = directory { size: 0, name: "home", parent: nil, children: map[string]*directory{}}

	currentDirectory := &homeDirectory 

	directories := []*directory{currentDirectory}//initialise with only home directory 

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if line[0] == "$"{//$ ls command doesn't matter
			if line[1] == "cd"{
				if line[2] == ".."{
					currentDirectory = currentDirectory.parent
				} else if line[2] != "/"{
					currentDirectory = currentDirectory.children[line[2]]
				}
			}
			continue //deals with $ ls command, denests
		}
		if line[0] == "dir"{ //used to add new directories to children map
			var childDirectory = &directory {name: line[1], size: 0, parent: currentDirectory, children: map[string]*directory{}}			
			currentDirectory.children[line[1]] = childDirectory
			directories = append(directories, childDirectory)	
			continue //denests			
		}
		//must be a file
		fileSize, err := strconv.Atoi(line[0])
		check(err)

		var childDirectory = &directory {name: line[1], size: fileSize, parent: currentDirectory, children: nil}
		currentDirectory.children[line[1]] = childDirectory
	}

	sizeToFree := 30_000_000 - (70_000_000 - directorySize(directories[0])) //30m - current available space
	bestDelete := directorySize(directories[0]) - sizeToFree //best solution found so far
	
	for i := 0; i < len(directories);i++{
		if directorySize(directories[i]) < sizeToFree {
			continue
		}
		extraDeleted := directorySize(directories[i]) //additional space deleted by deleting the directory
		if extraDeleted < bestDelete {
			bestDelete = extraDeleted
		}
	}

	fmt.Println(bestDelete)
}