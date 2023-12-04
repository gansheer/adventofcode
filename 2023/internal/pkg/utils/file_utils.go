// see https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559 for openfile code
package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLines(inputFile string) ([]string, error) {
	var lines []string
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return lines, nil
}
