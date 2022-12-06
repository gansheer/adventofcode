package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type entryData struct {
	first_1  int
	first_2  int
	second_1 int
	second_2 int
}

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
	var result int = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		data := extract(line)
		fmt.Println(data.first_1, "|", data.first_2, "<>", data.second_1, "|", data.second_2)
		if inBoth(data) {
			result += 1
		}
	}
	return result
}

func extract(line string) entryData {
	pairs := strings.Split(line, ",")
	first := strings.Split(pairs[0], "-")
	second := strings.Split(pairs[1], "-")
	return entryData{atoi(first[0]), atoi(first[1]), atoi(second[0]), atoi(second[1])}
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func inBoth(data entryData) bool {
	secondInfirst := data.first_1 <= data.second_1 && data.first_2 >= data.second_2
	firstInSecond := data.first_2 <= data.second_2 && data.first_1 >= data.second_1
	return secondInfirst || firstInSecond

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
