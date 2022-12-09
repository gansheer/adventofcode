package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Data struct {
	num_lines  int
	size_lines int
	forest     map[int][]int
	visibility map[int][]int
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

func doTheThing(fileScanner *bufio.Scanner) string {
	counter := 0
	var data Data = Data{forest: make(map[int][]int), visibility: make(map[int][]int)}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("********", line)
		var lineValues []int
		var lineHidden []int
		for _, treeHeight := range line {
			lineValues = append(lineValues, atoi(string(treeHeight)))
			lineHidden = append(lineHidden, 0)
		}
		data.forest[counter] = lineValues
		data.visibility[counter] = lineHidden
		counter += 1
	}
	data.num_lines = counter
	data.size_lines = len(data.forest[0])
	fmt.Println(data)
	fmt.Println("---> nothing")
	initBasicVisibility(&data)
	result := computeForVisibility(&data)
	return strconv.Itoa(result)
}

func initBasicVisibility(data *Data) {
	fmt.Println("----> initBasicVisibility")
	for i := 0; i < data.num_lines; i++ {
		fmt.Println(i)
		data.visibility[i][0] = 0
		data.visibility[i][data.size_lines-1] = 0

	}
	for i := 0; i < data.size_lines; i++ {
		fmt.Println(i)
		data.visibility[0][i] = 0
		data.visibility[data.num_lines-1][i] = 0
	}
	print(data)
}

func computeForVisibility(data *Data) int {
	maxResult := 1

	for i := 0; i < data.num_lines; i++ {
		for j := 0; j < data.size_lines; j++ {
			//if i == 2 && j == 3 {
			pointResult := computeForVisibilityPoint(data, i, j)
			if pointResult > maxResult {
				maxResult = pointResult
			}
			//}
		}
	}
	print(data)
	return maxResult
}

func computeForVisibilityPoint(data *Data, currentColumn int, currentLine int) int {
	fmt.Println("----> computeForVisibilityPoint", currentColumn, currentLine)

	currentSize := data.forest[currentLine][currentColumn]
	fmt.Println(currentSize)

	fmt.Println("starting the column to down")
	resultDown := 0
	if currentLine == 0 || currentLine == data.num_lines-1 {
		resultDown = 1
	} else {
	getDown:
		for j := currentLine + 1; j < data.num_lines; j++ {
			if data.forest[j][currentColumn] < currentSize {
				resultDown += 1
			} else {
				resultDown += 1
				break getDown
			}
		}
	}
	fmt.Println("to down", resultDown)

	fmt.Println("starting the column to up")
	resultUp := 0
	if currentLine == 0 || currentLine == data.num_lines-1 {
		resultUp = 1
	} else {
	getUp:
		for j := currentLine - 1; j >= 0; j-- {
			if data.forest[j][currentColumn] < currentSize {
				resultUp += 1
			} else {
				resultUp += 1
				break getUp
			}
		}
	}
	fmt.Println("to up", resultUp)

	fmt.Println("starting the line to right")
	resultRight := 0
	if currentColumn == 0 || currentColumn == data.size_lines-1 {
		resultRight = 1
	} else {
	getRight:
		for j := currentColumn + 1; j < data.size_lines; j++ {
			if data.forest[currentLine][j] < currentSize {
				resultRight += 1
			} else {
				resultRight += 1
				break getRight
			}
		}
	}
	fmt.Println("to right", resultRight)

	fmt.Println("starting the line to left")
	resultLeft := 0
	if currentColumn == 0 || currentColumn == data.size_lines-1 {
		resultLeft = 1
	} else {
	getLeft:
		for j := currentColumn - 1; j >= 0; j-- {
			if data.forest[currentLine][j] < currentSize {
				resultLeft += 1
			} else {
				resultLeft += 1
				break getLeft
			}
		}
	}
	fmt.Println("to left", resultLeft)

	fmt.Println((resultDown * resultUp * resultRight * resultLeft))
	data.visibility[currentLine][currentColumn] = (resultDown * resultUp * resultRight * resultLeft)
	fmt.Println(data.visibility[currentLine][currentColumn])
	return (resultDown * resultUp * resultRight * resultLeft)

}

func print(data *Data) {
	fmt.Println("------------")
	for i := 0; i < data.num_lines; i++ {
		fmt.Println(data.forest[i])
	}
	fmt.Println("************")
	for i := 0; i < data.num_lines; i++ {
		fmt.Println(data.visibility[i])
	}
	fmt.Println("------------")
	fmt.Println(data.num_lines, data.size_lines)
	fmt.Println("------------")
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
