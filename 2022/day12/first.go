package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Data struct {
	input map[Vertex]string
	edges []Edge
	maxX  int
	maxY  int
	start Vertex
	end   Vertex
}

type Vertex struct {
	x int
	y int
}

type Edge struct {
	from Vertex
	to   Vertex
}

func main() {

	fileName := "input0.txt"

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
	var data Data = Data{input: make(map[Vertex]string), edges: make([]Edge, 0)}
	counter := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		for index, value := range line {
			point := Vertex{x: index, y: counter}
			data.input[point] = string(value)
			if string(value) == "S" {
				data.start = point
			}
			if string(value) == "E" {
				data.end = point
			}

		}
		data.maxX = len(line) - 1
		counter += 1
	}
	data.maxY = counter - 1
	computeEdges(&data)
	//fmt.Println(data)
	print(&data)

	//displayAsciiValue("S")
	//displayAsciiValue("E")
	//displayAsciiValue("a")
	//displayAsciiValue("z")

	fmt.Println("---> nothing")

	return strconv.Itoa(counter)
}

func computeEdges(data *Data) {
	for vertex, value := range data.input {
		fmt.Println(vertex, value)

		if vertex.x > 0 {
			targetMinusY := Vertex{x: vertex.x - 1, y: vertex.y}
			targetValue := data.input[targetMinusY]
			fmt.Println(value, targetValue)
			if math.Abs(float64(decodeIntValue(targetValue)-decodeIntValue(value))) <= 1 {
				edge := Edge{vertex, targetMinusY}
				data.edges = append(data.edges, edge)
			}

		}

		if vertex.x < data.maxX-1 {
			targetPlusX := Vertex{x: vertex.x + 1, y: vertex.y}
			targetValue := data.input[targetPlusX]
			if math.Abs(float64(decodeIntValue(targetValue)-decodeIntValue(value))) <= 1 {
				edge := Edge{vertex, targetPlusX}
				data.edges = append(data.edges, edge)
			}
		}

		if vertex.y > 0 {
			targetMinusY := Vertex{x: vertex.x, y: vertex.y - 1}
			targetValue := data.input[targetMinusY]
			if math.Abs(float64(decodeIntValue(targetValue)-decodeIntValue(value))) <= 1 {
				edge := Edge{vertex, targetMinusY}
				data.edges = append(data.edges, edge)
			}
		}
		if vertex.y < data.maxY-1 {
			targetPlusY := Vertex{x: vertex.x, y: vertex.y + 1}
			targetValue := data.input[targetPlusY]
			if math.Abs(float64(decodeIntValue(targetValue)-decodeIntValue(value))) <= 1 {
				edge := Edge{vertex, targetPlusY}
				data.edges = append(data.edges, edge)
			}
		}
	}

}

func decodeIntValue(valueStr string) int {
	fmt.Println(valueStr)
	if valueStr == "S" {
		return 96
	}
	if valueStr == "E" {
		return 123
	}
	return int(valueStr[0])

}

func displayAsciiValue(valueStr string) {
	fmt.Printf("%s ASCII code is : %d\n", string(valueStr[0]), int(valueStr[0]))
}

func print(data *Data) {
	fmt.Println("------------")
	fmt.Println(data.input)
	fmt.Println("------------")
	for _, value := range data.edges {
		fmt.Println("from", value.from, "to", value.to)
	}
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
