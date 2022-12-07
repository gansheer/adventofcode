package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	//fileName := "input1.txt"
	// get file from terminal
	inputFile := "input1.txt"
	// declare chunk size
	maxSz := 1
	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	defer file.Close()

	// create buffer
	b := make([]byte, maxSz)

	startPacket := []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}

	counter := int(0)
	for {
		counter++
		// read content to buffer
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fileContent := string(b[:readTotal])
		// print content from buffer
		fmt.Println(counter, fileContent)
		startPacket = append(startPacket[1:14], fileContent)
		fmt.Println(startPacket)
		if has_dup(dup_count(startPacket)) {
			fmt.Println("duplicates found")
		} else if counter < 14 {
			fmt.Println("too early")
		} else {
			fmt.Println("valid result", counter)
			break
		}
	}

}

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func has_dup(dup_map map[string]int) bool {
	for _, v := range dup_map {
		if v > 1 {
			return true
		}
	}
	return false
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
