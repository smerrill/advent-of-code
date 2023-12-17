package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func GetCalibrationsWithWords(input string) (int, error) {
	return 0, nil
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
	input, err := ReadFileIntoString("input1.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	sum, err := GetCalibrations(input)
	if err != nil {
		log.Fatalf("Error getting calibrations: %v", err)
	}
	fmt.Println(sum)
}
