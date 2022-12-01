package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func doTheThing(fileScanner *bufio.Scanner) int {
	var result []int = []int{0, 0, 0}

	var currentElve int = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("---------")
		fmt.Println(currentElve)
		if calories, err := strconv.Atoi(line); err == nil {
			currentElve += int(calories)
		} else {
			sort.Ints(result)
			for i, v := range result {
				fmt.Println(i, v)
				if v < currentElve {
					result[i] = currentElve
					break
				}
			}

			currentElve = 0
		}

	}
	fmt.Println("---------")
	fmt.Println(currentElve)
	sort.Ints(result)
	for i, v := range result {
		fmt.Println(i, v)
		if v < currentElve {
			result[i] = currentElve
			break
		}
	}

	currentElve = 0
	fmt.Println(result)
	return addArray(result...)
}

func addArray(numbs ...int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
