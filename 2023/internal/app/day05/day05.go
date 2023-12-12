package day05

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gansheer/adventofcode/2023/internal/pkg/utils"
)

type XToY struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLenght           int
}

type XToYMap struct {
	Name   string
	Values []XToY
}

type Datas struct {
	Seeds   []int
	Results []int
	Maps    []XToYMap
}

func (x XToYMap) String() string {
	return fmt.Sprintf("\n**Name: %s\n**Values:\n%+v\n", x.Name, x.Values)
}

func (datas Datas) String() string {
	return fmt.Sprintf("Datas:\n%+v\n%+v\n%+v\n", datas.Seeds, datas.Maps, datas.Results)
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
	var result int

	var datas Datas
	var seeds []int
	var currentMap *XToYMap
	var seedToSoilMap XToYMap = XToYMap{Name: "seed-to-soil"}
	var soilToFertilizerMap XToYMap = XToYMap{Name: "soil-to-fertilizer"}
	var fertilizerToWaterMap XToYMap = XToYMap{Name: "fertilizer-to-water"}
	var waterToLightMap XToYMap = XToYMap{Name: "water-to-light"}
	var lightToTemperatureMap XToYMap = XToYMap{Name: "light-to-temperature"}
	var temperatureToHumidityMap XToYMap = XToYMap{Name: "temperature-to-humidity"}
	var humidityToLocationMap XToYMap = XToYMap{Name: "humidity-to-location"}

	for _, line := range lines {
		//fmt.Printf("line: %+v\n", line)
		if strings.HasPrefix(line, "seeds:") {
			seeds = extractSeeds(line)
		} else if strings.Contains(line, "map:") {

			if strings.HasPrefix(line, "seed-to-soil map:") {
				currentMap = &seedToSoilMap
			}
			if strings.HasPrefix(line, "soil-to-fertilizer map:") {
				currentMap = &soilToFertilizerMap
			}
			if strings.HasPrefix(line, "fertilizer-to-water map:") {
				currentMap = &fertilizerToWaterMap
			}
			if strings.HasPrefix(line, "water-to-light map:") {
				currentMap = &waterToLightMap
			}
			if strings.HasPrefix(line, "light-to-temperature map:") {
				currentMap = &lightToTemperatureMap
			}
			if strings.HasPrefix(line, "temperature-to-humidity map:") {
				currentMap = &temperatureToHumidityMap
			}
			if strings.HasPrefix(line, "humidity-to-location map:") {
				currentMap = &humidityToLocationMap
			}
		} else if len(strings.Fields(line)) == 0 {
			currentMap = nil
		} else {
			numbers := extractNumbers(line)
			fmt.Printf("numbers: %+v\n", numbers)
			switch currentMap {
			case &seedToSoilMap:
				seedToSoil := XToY{DestinationRangeStart: numbers[0], SourceRangeStart: numbers[1], RangeLenght: numbers[2]}
				seedToSoilMap.Values = append(seedToSoilMap.Values, seedToSoil)
			case &soilToFertilizerMap:
				soilToFertilizer := XToY{DestinationRangeStart: numbers[0], SourceRangeStart: numbers[1], RangeLenght: numbers[2]}
				soilToFertilizerMap.Values = append(soilToFertilizerMap.Values, soilToFertilizer)
			case &fertilizerToWaterMap:
				fertilizerToWater := XToY{DestinationRangeStart: numbers[0], SourceRangeStart: numbers[1], RangeLenght: numbers[2]}
				fertilizerToWaterMap.Values = append(fertilizerToWaterMap.Values, fertilizerToWater)
			case &waterToLightMap:
				waterToLight := XToY{DestinationRangeStart: numbers[0], SourceRangeStart: numbers[1], RangeLenght: numbers[2]}
				waterToLightMap.Values = append(waterToLightMap.Values, waterToLight)
			case &lightToTemperatureMap:
				lightToTemperature := XToY{DestinationRangeStart: numbers[0], SourceRangeStart: numbers[1], RangeLenght: numbers[2]}
				lightToTemperatureMap.Values = append(lightToTemperatureMap.Values, lightToTemperature)
			case &temperatureToHumidityMap:
				temperatureToHumidity := XToY{DestinationRangeStart: numbers[0], SourceRangeStart: numbers[1], RangeLenght: numbers[2]}
				temperatureToHumidityMap.Values = append(temperatureToHumidityMap.Values, temperatureToHumidity)
			case &humidityToLocationMap:
				humidityToLocation := XToY{DestinationRangeStart: numbers[0], SourceRangeStart: numbers[1], RangeLenght: numbers[2]}
				humidityToLocationMap.Values = append(humidityToLocationMap.Values, humidityToLocation)

			}
		}

	}

	datas.Seeds = seeds
	datas.Maps = append(datas.Maps, seedToSoilMap)
	datas.Maps = append(datas.Maps, soilToFertilizerMap)
	datas.Maps = append(datas.Maps, fertilizerToWaterMap)
	datas.Maps = append(datas.Maps, waterToLightMap)
	datas.Maps = append(datas.Maps, lightToTemperatureMap)
	datas.Maps = append(datas.Maps, temperatureToHumidityMap)
	datas.Maps = append(datas.Maps, humidityToLocationMap)

	computeData(&datas)

	fmt.Printf("\n%+v\n", datas)

	result = utils.MinIntSlice(datas.Results)

	return strconv.FormatInt(int64(result), 10), nil
}

func computeData(datas *Datas) {
	fmt.Print("\n***** computeData ****\n")
	for _, seed := range datas.Seeds {
		fmt.Printf("\nseed %+v", seed)
		target := seed
		for _, transformMap := range datas.Maps {
			target = getTransformedValueInMap(target, transformMap)
		}
		fmt.Printf("\ntarget for seed %+v is %+v\n", seed, target)
		datas.Results = append(datas.Results, target)
	}
}

func getTransformedValueInMap(seed int, xToYMap XToYMap) int {
	target := seed
	for _, item := range xToYMap.Values {
		if seed >= item.SourceRangeStart && seed <= item.SourceRangeStart+item.RangeLenght {
			step := seed - item.SourceRangeStart
			target = item.DestinationRangeStart + step
			fmt.Printf("\nmap %+v gives %+v (+%v)", xToYMap.Name, target, item)
			return target
		}
		fmt.Printf("\nmap %+v gives %+v (+%v)", xToYMap.Name, target, item)
	}
	return target
}

func getSeedFromLowestResult(datas Datas) int {
	lowestResultIndex := 0
	for index, value := range datas.Results {
		if value <= datas.Results[lowestResultIndex] {
			lowestResultIndex = index
		}
	}
	return datas.Seeds[lowestResultIndex]
}

func extractSeeds(line string) []int {
	return extractNumbers(strings.Split(line, ":")[1])
}

func extractNumbers(input string) []int {
	seedsStr := strings.Fields(input)
	var seeds []int
	for _, seedStr := range seedsStr {
		seeds = append(seeds, utils.ToInt(seedStr))
	}
	return seeds
}
