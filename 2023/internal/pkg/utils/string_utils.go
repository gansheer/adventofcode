package utils

import (
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
)

func ExtractNumbers(input string) []string {
	re := regexp.MustCompile("[0-9]")
	return re.FindAllString(input, -1)
}

func FindFirst(input string, tokens []string) (string, int) {
	resultToken := ""
	resultIndex := len(input)

	for _, token := range tokens {
		index := strings.Index(input, token)
		if index != -1 && resultIndex > index {
			resultIndex = index
			resultToken = token
		}
	}
	return resultToken, resultIndex
}

func ReplaceStringNumberByDigit(input string) string {
	var NUMBERS_STR = map[string]string{
		"one":   "o1ne",
		"two":   "tw2o",
		"three": "thr3ee",
		"four":  "fo4ur",
		"five":  "fi5ve",
		"six":   "s6ix",
		"seven": "se7ven",
		"eight": "ei8ght",
		"nine":  "ni9ne",
	}
	keepOn := true
	counter := 0
	for keepOn && counter < 30 {
		token, _ := FindFirst(input, maps.Keys(NUMBERS_STR))
		counter += 1
		if token == "" {
			keepOn = false
		} else {
			input = strings.Replace(input, token, NUMBERS_STR[token], 1)
		}
	}
	return input
}
