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

	computeForVisibility(&data)

	result := int(0)
	for i := 0; i < data.size_lines; i++ {
		for j := 0; j < data.num_lines; j++ {
			result += data.visibility[i][j]
		}

	}
	return strconv.Itoa(result)
}

func initBasicVisibility(data *Data) {
	fmt.Println("----> initBasicVisibility")
	for i := 0; i < data.num_lines; i++ {
		fmt.Println(i)
		data.visibility[i][0] = 1
		data.visibility[i][data.size_lines-1] = 1

	}
	for i := 0; i < data.size_lines; i++ {
		fmt.Println(i)
		data.visibility[0][i] = 1
		data.visibility[data.num_lines-1][i] = 1
	}
	print(data)
}

func computeForVisibility(data *Data) {
	fmt.Println("----> computeForVisibility")

	fmt.Println("starting the columns from up, without first and last")
	for i := 1; i < data.num_lines-1; i++ {
		maxSize := data.forest[0][i]
		println("* forest", 0, i, "value", maxSize)
		for j := 1; j < data.size_lines-1; j++ {
			if data.forest[j][i] > maxSize {
				data.visibility[j][i] = 1
				maxSize = data.forest[j][i]
			}
		}

	}
	fmt.Println("starting the columns from down, without first and last")
	for i := 1; i < data.num_lines-1; i++ {
		maxSize := data.forest[data.num_lines-1][i]
		println("* forest", data.num_lines-1, i, "value", maxSize)
		for j := data.size_lines - 2; j >= 1; j-- {
			fmt.Println("forest", "col", i, "line", j, "size", data.forest[j][i], "maxSize", maxSize)
			if data.forest[j][i] > maxSize {
				data.visibility[j][i] = 1
				maxSize = data.forest[j][i]
			}
		}

	}

	fmt.Println("starting the lines from left, without first and last")
	for i := 1; i < data.size_lines-1; i++ {
		maxSize := data.forest[i][0]
		println("* forest", i, 0, "value", maxSize)
		for j := 1; j < data.size_lines-2; j++ {
			fmt.Println("forest", "line", i, "col", j, "size", data.forest[i][j], "maxSize", maxSize)
			if data.forest[i][j] > maxSize {
				fmt.Println("visible")
				data.visibility[i][j] = 1
				maxSize = data.forest[i][j]
			}
		}
	}

	fmt.Println("starting the lines from right, without first and last")
	for i := 1; i < data.size_lines-1; i++ {
		maxSize := data.forest[i][data.size_lines-1]
		println("* forest", i, data.size_lines-1, "value", maxSize)
		for j := data.size_lines - 2; j >= 1; j-- {
			fmt.Println("forest", "line", i, "col", j, "size", data.forest[i][j], "maxSize", maxSize)
			if data.forest[i][j] > maxSize {
				fmt.Println("visible")
				data.visibility[i][j] = 1
				maxSize = data.forest[i][j]
			}
		}

	}

	print(data)
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
