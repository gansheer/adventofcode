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
	roundScores := roundScores()
	shapeScores := shapeScores()

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// 0, 3, 6
		roundScore := roundScores[line[len(line)-1:]]
		// round
		round := line[:1] + " " + strconv.Itoa(int(roundScore))
		shapeScore := shapeScores[round]

		//fmt.Println("roundscore:", roundScore)
		//fmt.Println("round:", round)
		//fmt.Println("shapescore:", shapeScore)
		result += roundScore
		result += shapeScore
	}
	return result
}

func shapeScores() map[string]int32 {
	return map[string]int32{
		"A 0": 3, // need to loose > rock vs scissors
		"A 3": 1, // need a draw > rock vs rock
		"A 6": 2, // need to win > rock vs paper
		"B 0": 1, // need to loose > paper vs rock
		"B 3": 2, // need to draw > paper vs paper
		"B 6": 3, // need to win > paper vs scissors
		"C 0": 2, // need to loose > scissors vs paper
		"C 3": 3, // need to draw > scissors vs scissors
		"C 6": 1, // need to win > scissors vs rock
	}

}

func roundScores() map[string]int32 {
	return map[string]int32{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
