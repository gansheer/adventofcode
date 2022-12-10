package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	commands   []string
	cycleModif []int
	x          []int
	display    [][]string
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
	currentCycle := 1
	var data Data = Data{commands: make([]string, 250), cycleModif: make([]int, 250)}
	data.cycleModif[0] = 1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("starting after cycle", currentCycle)
		switch string(line[0]) {
		case "a":

			execAddX(&data, line, currentCycle)
			currentCycle += 2
		case "n":
			execNoop(&data, line, currentCycle)
			currentCycle += 1
		default:
			return "ERROR"

		}
	}
	//print(&data)

	fmt.Println("---> nothing")
	compute(&data)
	fmt.Println("---> result")
	display(&data)
	print(&data)
	printDisplay(&data)
	fmt.Println("---> result cycles")
	fmt.Println(data.x[20])
	fmt.Println(data.x[60])
	fmt.Println(data.x[100])
	fmt.Println(data.x[140])
	fmt.Println(data.x[180])
	fmt.Println(data.x[220])
	result := data.x[20]*20 + data.x[60]*60 + data.x[100]*100 +
		data.x[140]*140 + data.x[180]*180 + data.x[220]*220

	return strconv.Itoa(result)
}

func execNoop(data *Data, command string, currentCycle int) {
	fmt.Println("----> execAddX")
	data.commands = append(data.commands, command)
	//data.cycleModif[currentCycle] = data.cycleModif[currentCycle]
	print(data)
}

func execAddX(data *Data, command string, currentCycle int) {
	fmt.Println("----> execAddX")
	data.commands = append(data.commands, command)
	value := atoi(strings.Split(command, " ")[1])
	// can be dangerous if already something
	data.cycleModif[currentCycle+1] = 0
	data.cycleModif[currentCycle+2] = data.cycleModif[currentCycle+1] + value
	//fmt.Println(data.cycleModif)
}

func compute(data *Data) {
	data.x = make([]int, 250)
	for index, value := range data.cycleModif {
		if index == 0 {
			data.x[index] = value
		} else {
			data.x[index] = data.x[index-1] + value
		}
	}
}

func display(data *Data) {
	data.display = make([][]string, 6)
	for i := 0; i < 6; i++ {
		data.display[i] = make([]string, 40)
		for j := 0; j < 40; j++ {
			data.display[i][j] = "."
		}
	}

	for index, value := range data.x[1:] {
		line := int(index / 40)
		column := int(index % 40)
		fmt.Println(index, line, column, value)
		if line > 5 {
			break
		} else {
			if column-1 == value || column == value || column+1 == value {
				data.display[line][column] = "#"
			} else {
				data.display[line][column] = "."
			}

		}
	}

}

func print(data *Data) {
	fmt.Println("------------")
	for i := 0; i < len(data.commands); i++ {
		fmt.Println(data.commands[i])
	}
	fmt.Println("************")
	/*for i := 0; i < len(data.cycleModif); i++ {
		fmt.Println(data.cycleModif[i])
	}*/
	fmt.Println(data.cycleModif)
	fmt.Println("************")
	/*for i := 0; i < len(data.cycleValue); i++ {
		fmt.Println(i, data.cycleValue[i])
	}*/
	fmt.Println(data.x)
	fmt.Println("------------")
}

func printDisplay(data *Data) {
	for i := 0; i < 6; i++ {
		fmt.Println(data.display[i])
	}
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
