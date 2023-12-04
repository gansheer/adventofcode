package day02

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gansheer/adventofcode/2023/internal/pkg/utils"
)

type Game struct {
	Id    int
	Red   int
	Blue  int
	Green int
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
	for _, line := range lines {
		game := extractGame(line)
		fmt.Printf("%+v\n", game)
		result += int64(game.Red * game.Blue * game.Green)
	}
	return strconv.FormatInt(result, 10), nil
}

// TODO
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

// TODO
func doTheThingSimple(lines []string) (string, error) {
	var result int64 = int64(0)
	for _, line := range lines {
		game := extractGame(line)
		fmt.Printf("%+v\n", game)
		if isGameUnderLimit(game, LimitGameSimple()) {
			result += int64(game.Id)
		}
	}
	return strconv.FormatInt(result, 10), nil
}

func extractGame(line string) Game {
	game := Game{Red: 0, Blue: 0, Green: 0}
	gameToken := strings.Split(line, ":")[0]
	game.Id, _ = strconv.Atoi(strings.Split(gameToken, " ")[1])
	showTokens := strings.Split(strings.Split(line, ":")[1], ";")
	for _, showToken := range showTokens {
		colorTokens := strings.Split(showToken, ",")
		for _, colorToken := range colorTokens {
			item := strings.Split(strings.TrimSpace(colorToken), " ")
			switch color := item[1]; color {
			case "red":
				if game.Red < utils.ToInt(item[0]) {
					game.Red = utils.ToInt(item[0])
				}
			case "blue":
				if game.Blue < utils.ToInt(item[0]) {
					game.Blue = utils.ToInt(item[0])
				}
			case "green":
				if game.Green < utils.ToInt(item[0]) {
					game.Green = utils.ToInt(item[0])
				}
			}
		}
	}
	return game

}

func isGameUnderLimit(current Game, limit Game) bool {
	return current.Red <= limit.Red && current.Blue <= limit.Blue && current.Green <= limit.Green
}

func LimitGameSimple() Game {
	return Game{Id: -1, Red: 12, Green: 13, Blue: 14}
}
