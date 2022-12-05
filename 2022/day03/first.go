package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type entryData struct {
	line   string
	first  string
	second string
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
		data := cut(line)
		fmt.Println(len(data.first), "|", data.first)
		fmt.Println(len(data.second), "|", data.second)
		common := inBoth(data.first, data.second)
		fmt.Println(common)
		result += score(common)
		fmt.Println("-----------")
	}
	return result
}

func cut(line string) entryData {
	return entryData{line, string(line[0 : len(line)/2]), string(line[len(line)/2:])}
}

func inBoth(first string, second string) []string {
	a := strings.Split(first, "")
	b := strings.Split(second, "")
	//fmt.Println(a)
	//fmt.Println(b)

	m := make(map[string]uint8)
	for _, k := range a {
		m[k] |= (1 << 0)
	}
	for _, k := range b {
		m[k] |= (1 << 1)
	}

	var inAAndB, inAButNotB, inBButNotA []string
	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		switch {
		case a && b:
			inAAndB = append(inAAndB, k)
		case a && !b:
			inAButNotB = append(inAButNotB, k)
		case !a && b:
			inBButNotA = append(inBButNotA, k)
		}
	}
	//fmt.Println(inAAndB)
	//fmt.Println(inAButNotB)
	//fmt.Println(inBButNotA)
	return inAAndB
}

func score(commons []string) int {
	//fmt.Printf("%s ASCII code is : %d\n", string('A'), int('A'))
	//fmt.Printf("%s ASCII code is : %d\n", string('B'), int('B'))
	//fmt.Printf("%s ASCII code is : %d\n", string('a'), int('a'))
	//fmt.Printf("%s ASCII code is : %d\n", string('b'), int('b'))
	//fmt.Println(int(commons[0][0]), " ")
	if commons[0][0] <= 90 {
		// UPPER
		fmt.Println(commons[0], int(commons[0][0])-(38))
		return int(commons[0][0]) - 38
	} else if commons[0][0] > 90 {
		// lower
		fmt.Println(commons[0], int(commons[0][0])-(48+48))
		return int(commons[0][0]) - (48 + 48)
	}
	return 0
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
