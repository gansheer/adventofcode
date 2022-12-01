package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fileName := "input1.txt"

	// open the file
	file, err := os.Open(fileName)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	// read line by line
	fmt.Println(doTheThing(fileScanner))

	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()

}

func doTheThing(fileScanner *bufio.Scanner) int32 {
	var result int32 = 0

	var currentElve int32 = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if calories, err := strconv.Atoi(line); err == nil {
			currentElve += int32(calories)
		} else {
			if currentElve > result {
				result = currentElve
			}
			currentElve = 0
		}
	}
	return result
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
