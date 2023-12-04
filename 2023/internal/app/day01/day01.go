package day01

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gansheer/adventofcode/2023/internal/pkg/utils"
)

func computeComplex(inputFile string) (string, error) {
	// see https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559 for openfile code
	lines, err := utils.ReadFileLines(inputFile)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	result, err := doTheThingComplex(lines)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return result, nil
}

func doTheThingComplex(lines []string) (string, error) {
	var result int64 = int64(0)
	for _, line := range lines {
		modifiedLine := utils.ReplaceStringNumberByDigit(line)
		numbers := utils.ExtractNumbers(modifiedLine)
		numberString := strings.Join(extractFirstLastNumber(numbers), "")
		number, err := strconv.ParseInt(numberString, 10, 64)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		fmt.Println(line, " ", modifiedLine, " ", utils.ExtractNumbers(line), " ", numberString)
		result += number

	}
	return strconv.FormatInt(result, 10), nil
}

func computeSimple(inputFile string) (string, error) {
	// see https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go/16615559#16615559 for openfile code
	lines, err := utils.ReadFileLines(inputFile)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	result, err := doTheThingSimple(lines)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return result, nil
}

func doTheThingSimple(lines []string) (string, error) {
	var result int64 = int64(0)
	for _, line := range lines {
		numbers := utils.ExtractNumbers(line)
		numberString := strings.Join(extractFirstLastNumber(numbers), "")
		number, err := strconv.ParseInt(numberString, 10, 64)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		fmt.Println(utils.ExtractNumbers(line), " ", numberString)
		result += number

	}
	return strconv.FormatInt(result, 10), nil
}

func extractFirstLastNumber(numbers []string) []string {
	return []string{numbers[0], numbers[len(numbers)-1]}
}
