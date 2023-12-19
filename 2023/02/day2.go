package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	adventofcode "stevenwmerrill.com/adventofcode/2023/v2"
)

func GetPossibleGames(input string) (int, error) {
	var output int
	lines := strings.Split(input, "\n")
OUTER:
	for _, line := range lines {
		gameResult, err := ParseLine(line)
		if err != nil {
			return 0, err
		}
		for _, cubeCount := range gameResult.CubeCounts {
			if cubeCount.Red > 12 || cubeCount.Green > 13 || cubeCount.Blue > 14 {
				continue OUTER
			}
		}
		output += gameResult.Id
	}

	return output, nil
}

type GameResult struct {
	Id         int
	CubeCounts []CubeCounts
}
type CubeCounts struct {
	Red   int
	Green int
	Blue  int
}

type ParseLineFunc func(line string) GameResult

func ParseLine(line string) (GameResult, error) {
	var err error
	output := GameResult{}
	lineRegex := regexp.MustCompile(`Game (\d+): (.*)`)
	rgbRegex := regexp.MustCompile(`(?:(\d+) (red|green|blue))+`)
	lineMatches := lineRegex.FindAllStringSubmatch(line, 3)

	if len(lineMatches) == 0 {
		return output, nil
	}

	if len(lineMatches[0]) > 0 {
		output.Id, err = strconv.Atoi(lineMatches[0][1])
		if err != nil {
			return output, err
		}
		lineSegments := strings.Split(lineMatches[0][2], ";")
		for _, segment := range lineSegments {
			rgbMatches := rgbRegex.FindAllStringSubmatch(segment, 3)
			if len(rgbMatches) > 0 {
				cubeCount := CubeCounts{}
				for _, match := range rgbMatches {
					switch match[2] {
					case "red":
						cubeCount.Red, err = strconv.Atoi(match[1])
						if err != nil {
							return output, err
						}
					case "green":
						cubeCount.Green, err = strconv.Atoi(match[1])
						if err != nil {
							return output, err
						}
					case "blue":
						cubeCount.Blue, err = strconv.Atoi(match[1])
						if err != nil {
							return output, err
						}
					}
				}
				output.CubeCounts = append(output.CubeCounts, cubeCount)
			}
		}
	}

	return output, nil
}

func main() {
	input, err := adventofcode.ReadFileIntoString("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	sum1, err := GetPossibleGames(input)
	if err != nil {
		log.Fatalf("Error getting possible games: %v", err)
	}
	// sum2, err := GetCalibrationsScanner(input, ScanNumeralsAndNumberWords)
	// if err != nil {
	// 	log.Fatalf("Error getting calibrations: %v", err)
	// }
	// fmt.Println(sum1, sum2)
	fmt.Println(sum1)
}
