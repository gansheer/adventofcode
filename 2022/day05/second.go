package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CargoCrane struct {
	crates map[string][]string
}
type Move struct {
	number string
	source string
	target string
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

	var cargoCraneLines []string = []string{}
	var cargoCrane CargoCrane

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("********", line)
		if len(line) != 0 {
			cargoCraneLines = append(cargoCraneLines, line)
			fmt.Println(cargoCraneLines)
		} else {
			fmt.Println("Generating cargo crane")
			cargoCraneLines = reverseArray(cargoCraneLines)
			fmt.Println(cargoCraneLines)
			cargoCrane = initCargoCrane(cargoCraneLines)
			fillCargoCrane(cargoCraneLines[1:], cargoCrane)
			fmt.Println(cargoCrane)
			print(cargoCrane)
			break
		}
	}
	fmt.Println("---> applyMoves")
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("%%%", line)
		move := extractMove(line)
		fmt.Println(move)
		cargoCrane = applyMove(move, cargoCrane)

	}
	fmt.Println("---> final state")
	fmt.Println(cargoCrane)

	return computeResult(cargoCrane)
}

func initCargoCrane(cargoCraneLines []string) CargoCrane {
	fmt.Println("----> Initializing cargo crane")
	var cargoCrane = CargoCrane{crates: make(map[string][]string)}
	initLine := cargoCraneLines[0]
	fmt.Println(initLine)
	numbers := strings.Split(initLine, "   ")
	fmt.Println(numbers)
	for _, number := range numbers {
		fmt.Println(number)
		cargoCrane.crates[strings.TrimSpace(number)] = []string{}
	}
	return cargoCrane

}

func fillCargoCrane(cargoCraneLines []string, cargoCrane CargoCrane) {
	fmt.Println("----> fillCargoCrane")
	for _, line := range cargoCraneLines {
		fmt.Println(len(line), "|", line, "|")
		for i := 0; i < len(line); i += 4 {
			//fmt.Println("i ", i)
			crate := line[i : i+3]
			crateIndex := (i / 4) + 1
			if len(strings.TrimSpace(crate)) != 0 {
				cargoCrane.crates[strconv.Itoa(crateIndex)] = append(cargoCrane.crates[strconv.Itoa(crateIndex)], crate)
			}
		}
	}
}

func extractMove(line string) Move {
	linesTo := strings.Split(line, "to")
	linesFrom := strings.Split(linesTo[0], "from")
	lines := strings.Split(linesFrom[0], "move")
	return Move{number: strings.TrimSpace(lines[1]), source: strings.TrimSpace(linesFrom[1]), target: strings.TrimSpace(linesTo[1])}

}

func applyMove(move Move, cargoCrane CargoCrane) CargoCrane {
	count, _ := strconv.Atoi(move.number)
	crateSize := len(cargoCrane.crates[move.source])
	objectsToMove := cargoCrane.crates[move.source][crateSize-count : crateSize]
	fmt.Println(crateSize, count, objectsToMove)
	fmt.Println(cargoCrane.crates[move.source])
	fmt.Println(cargoCrane.crates[move.target])

	cargoCrane.crates[move.source] = cargoCrane.crates[move.source][0 : crateSize-count]
	cargoCrane.crates[move.target] = append(cargoCrane.crates[move.target], objectsToMove...)

	fmt.Println(cargoCrane)
	return cargoCrane

}

func computeResult(cargoCrane CargoCrane) string {
	results := []string{}
	//fmt.Println(len(cargoCrane.crates))
	for i := 1; i <= len(cargoCrane.crates); i++ {
		crateIndex := strconv.Itoa(i)
		fmt.Println(crateIndex, cargoCrane.crates[crateIndex])
		crate := cargoCrane.crates[crateIndex][len(cargoCrane.crates[crateIndex])-1]
		//fmt.Println(crate)
		results = append(results, strings.TrimRight(strings.TrimLeft(crate, "["), "]"))
	}
	return strings.Join(results, "")
}

func print(cargoCrane CargoCrane) {
	fmt.Println("************")
	for i := 1; i <= len(cargoCrane.crates); i++ {
		crateIndex := strconv.Itoa(i)
		fmt.Println(crateIndex, cargoCrane.crates[crateIndex])
	}
	fmt.Println("************")
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
