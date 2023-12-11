package day04

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/gansheer/adventofcode/2023/internal/pkg/utils"
)

type Card struct {
	PlayerNumbers  []int
	WinningNumbers []int
	Index          int
	NumOfCards     int
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
	var result float64 = float64(0)

	var cards []*Card
	for _, line := range lines {
		fmt.Printf("%+v\n", line)
		card := extractCard(line)
		fmt.Printf("%+v\n", card)
		cards = append(cards, card)
	}

	for index, c := range cards {
		playerWins := utils.Intersection(c.PlayerNumbers, c.WinningNumbers)
		if len(playerWins) > 0 {
			for i := 1; i <= len(playerWins); i++ {
				cards[index+i].NumOfCards += cards[index].NumOfCards
			}
		}
	}

	for _, c := range cards {
		result += float64(c.NumOfCards)
	}

	return strconv.FormatInt(int64(result), 10), nil
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
	var result float64 = float64(0)

	//var cards []*Card
	for _, line := range lines {
		fmt.Printf("%+v\n", line)
		card := extractCard(line)
		fmt.Printf("%+v\n", card)
		//cards := append(cards, card)
		playerWins := utils.Intersection(card.PlayerNumbers, card.WinningNumbers)
		fmt.Printf("%+v\n", playerWins)
		if len(playerWins) > 0 {
			result += math.Pow(2, float64(len(playerWins)-1))
		}
		fmt.Printf("%+v\n", result)
	}

	return strconv.FormatInt(int64(result), 10), nil
}

func extractCard(line string) *Card {
	card := Card{NumOfCards: 1}
	card.Index = utils.ToInt(strings.Fields(strings.Split(line, ":")[0])[1])
	datas := strings.TrimSpace(strings.Split(line, ":")[1])
	//fmt.Printf("%+v\n", datas)
	wNumbers := strings.Fields(strings.TrimSpace(strings.Split(datas, "|")[0]))
	for _, wNumber := range wNumbers {
		//fmt.Printf("%+v\n", wNumber)
		card.WinningNumbers = append(card.WinningNumbers, utils.ToInt(strings.TrimSpace(wNumber)))
	}
	pNumbers := strings.Fields(strings.TrimSpace(strings.Split(datas, "|")[1]))
	for _, pNumber := range pNumbers {
		card.PlayerNumbers = append(card.PlayerNumbers, utils.ToInt(strings.TrimSpace(pNumber)))
	}

	return &card
}
