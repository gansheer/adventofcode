package day03

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gansheer/adventofcode/2023/internal/pkg/utils"
)

type Position struct {
	StartColumn int
	EndColumn   int
	LineIndex   int
	Number      int
}

// TODO
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

// TODO
func doTheThingComplex(lines []string) (string, error) {
	var result int64 = int64(0)

	var numbers []*Position
	for index, line := range lines {
		extractedPositions := findNumbers(line)
		for _, extractedPosition := range extractedPositions {
			extractedPosition.LineIndex = index
			fmt.Printf("number: %+v\n", extractedPosition)
		}
		numbers = append(numbers, extractedPositions...)
		fmt.Printf("%+v\n", len(numbers))
	}

	var symboles []*Position
	for index, line := range lines {
		extractedPositions := findSymboles(line)
		for _, extractedPosition := range extractedPositions {
			extractedPosition.LineIndex = index
			fmt.Printf("symbole: %+v\n", extractedPosition)
		}
		symboles = append(symboles, extractedPositions...)
		fmt.Printf("%+v\n", len(symboles))
	}

	fmt.Printf("** There are %+v numbers\n", len(numbers))
	fmt.Printf("** There are %+v symboles\n", len(symboles))

	for i, symbole := range symboles {
		fmt.Printf("symbole %+v: %+v\n", i, symbole)

		numbersNear := numbersTouchingSymbol(symbole, numbers)
		fmt.Printf("%+v numbersNear\n", len(numbersNear))
		if len(numbersNear) == 2 {
			fmt.Printf("%+v valide\n", symbole)
			result += int64(numbersNear[0].Number) * int64(numbersNear[1].Number)
		}

	}

	return strconv.FormatInt(result, 10), nil
}

func computeSimple(inputFile string) (string, error) {
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
	var positions []*Position

	for index, line := range lines {
		extractedPositions := findNumbers(line)
		for _, extractedPosition := range extractedPositions {
			extractedPosition.LineIndex = index
			fmt.Printf("%+v\n", extractedPosition)
		}
		positions = append(positions, extractedPositions...)
		fmt.Printf("%+v\n", len(positions))
	}

	for _, position := range positions {
		if isValid(position, lines) {
			result += int64(position.Number)
		}
	}

	return strconv.FormatInt(result, 10), nil
}

func findNumbers(line string) []*Position {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringSubmatchIndex(line, -1)

	result := make([]*Position, len(matches))
	for i, match := range matches {
		start := match[0]
		end := match[1]
		number := utils.ToInt(line[start:end])

		result[i] = &Position{
			Number:      number,
			StartColumn: start,
			EndColumn:   end,
		}
	}

	return result
}

func findSymboles(line string) []*Position {
	result := make([]*Position, 0)

	for i, char := range line {
		if char == '*' {
			result = append(result, &Position{
				StartColumn: i,
				EndColumn:   i,
			})
		}
	}

	return result
}

func isValid(position *Position, lines []string) bool {
	from := position.StartColumn - 1
	if from < 0 {
		from = 0
	}
	to := position.EndColumn + 1
	if to > len(lines[0]) {
		to = len(lines[0])
	}

	for looplines := position.LineIndex - 1; looplines <= position.LineIndex+1; looplines++ {
		if looplines < 0 || looplines >= len(lines) {
			continue
		}
		symbolFound := strings.IndexAny(lines[looplines][from:to], "+#$*@/=%-&")
		if symbolFound > -1 {
			return true
		}
	}

	return false
}

func numbersTouchingSymbol(symbole *Position, numbers []*Position) []*Position {
	fmt.Printf("numbersTouchingSymbol %+v\n", symbole)
	var result []*Position
	for _, number := range numbers {
		valid := true
		if number.LineIndex < (symbole.LineIndex-1) || number.LineIndex > (symbole.LineIndex+1) {
			valid = false
		}
		if valid && number.EndColumn < symbole.StartColumn {
			valid = false
		}

		if valid && number.StartColumn > symbole.EndColumn+1 {
			valid = false
		}
		if valid {
			fmt.Printf("numbersTouchingSymbol %+v valid\n", number)
			result = append(result, number)
		}
	}

	return result
}
