package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	roundScores := roundScores()
	shapeScores := shapeScores()

	for fileScanner.Scan() {
		line := fileScanner.Text()
		result += roundScores[line]
		result += shapeScores[line[len(line)-1:]]
	}
	return result
}

func roundScores() map[string]int32 {
	return map[string]int32{
		"A X": 3,
		"A Y": 6,
		"A Z": 0,
		"B X": 0,
		"B Y": 3,
		"B Z": 6,
		"C X": 6,
		"C Y": 0,
		"C Z": 3,
	}
}

func shapeScores() map[string]int32 {
	return map[string]int32{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
