package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	adventofcode "stevenwmerrill.com/adventofcode/2023/v2"
)

func GetCalibrationsWithWords(input string) (int, error) {
	return 0, nil
}

func ScanNumerals(data []byte, atEOF bool) (advance int, token []byte, err error) {
	digitRegex := regexp.MustCompile(`\d`)
	// Scan single digits
	for i := 0; i < len(data); i += 1 {
		r := data[i]
		/*
			If you want to handle multi-byte Unicode chars, do this:

			for i, width := 0, 0; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
		*/
		if digitRegex.MatchString(string(r)) {
			return i + 1, []byte{r}, nil
		}
	}
	// Request more data.
	return 0, nil, nil
}

func ScanNumeralsAndNumberWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	digitRegex := regexp.MustCompile(`\d`)
	numberWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	// Scan single digits
	for i := 0; i < len(data); {
		r := data[i]
		if digitRegex.MatchString(string(r)) {
			return i + 1, []byte{byte(r)}, nil
		} else {
			for k, v := range numberWords {
				if string(data[i:i+len(v)]) == v {
					// The impedance mismatch of returning a token of []byte instead of an int isn't great.
					return i + (len(v) - 1), []byte(fmt.Sprint(k + 1)), nil
				}
			}
		}
		i += 1
	}
	// Request more data.
	return 0, nil, nil
}

func GetCalibrationsScanner(input string, splitFunc bufio.SplitFunc) (int, error) {
	sum := 0
	for _, j := range strings.Split(input, "\n") {
		var output []string
		scanner := bufio.NewScanner(strings.NewReader(j))
		scanner.Split(splitFunc)
		for scanner.Scan() {
			output = append(output, scanner.Text())
		}
		if l := len(output); l > 0 {
			addition, err := strconv.Atoi(output[0] + output[l-1])
			if err != nil {
				return 0, err
			}
			sum += addition
		}
	}
	return sum, nil
}

func GetCalibrations(input string) (int, error) {
	var sum int
	digitRegex := regexp.MustCompile(`\d`)
	var firstInt string
	var lastInt string
	var addition int
	var err error

	for _, j := range strings.Split(input, "\n") {
		firstInt = ""
		lastInt = ""
		addition = 0
		for _, k := range strings.Split(j, "") {
			if digitRegex.MatchString(k) {
				if firstInt == "" {
					firstInt = k
				}
				lastInt = k
			}
		}
		if firstInt != "" && lastInt != "" {
			addition, err = strconv.Atoi(firstInt + lastInt)
			if err != nil {
				return 0, err
			}
			sum += addition
		}
	}
	return sum, err
}

func main() {
	input, err := adventofcode.ReadFileIntoString("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	sum1, err := GetCalibrationsScanner(input, ScanNumerals)
	if err != nil {
		log.Fatalf("Error getting calibrations: %v", err)
	}
	sum2, err := GetCalibrationsScanner(input, ScanNumeralsAndNumberWords)
	if err != nil {
		log.Fatalf("Error getting calibrations: %v", err)
	}
	fmt.Println(sum1, sum2)
}
