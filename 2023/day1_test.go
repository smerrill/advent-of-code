package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const day1Input1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const day1Input2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestGetCalibration(t *testing.T) {
	sum, err := GetCalibrations(day1Input1)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if sum != 142 {
		t.Errorf("Expected 142, got %d", sum)
	}
	sum, err = GetCalibrationsScanner(day1Input1, ScanNumerals)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if sum != 142 {
		t.Errorf("Expected 142, got %d", sum)
	}
	sum, err = GetCalibrationsScanner(day1Input2, ScanNumeralsAndNumberWords)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if sum != 281 {
		t.Errorf("Expected 281, got %d", sum)
	}
}

func TestNumeralScanner(t *testing.T) {
	testCases := map[string][]string{
		"a1z3r5": {"1", "3", "5"},
		"1":      {"1"},
		"":       nil,
		"11":     {"1", "1"},
		"a":      nil,
	}

	for k, v := range testCases {
		passed, output, err := testScanner(k, ScanNumerals, v)
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}
		if !passed {
			t.Errorf("Expected `%v`, got `%v`", output, v)
		}
	}
}

func TestNumeralAndNumberWordScanner(t *testing.T) {
	testCases := map[string][]string{
		"a1z3r5":      {"1", "3", "5"},
		"1":           {"1"},
		"":            nil,
		"11":          {"1", "1"},
		"a":           nil,
		"one":         {"1"},
		"1two3":       {"1", "2", "3"},
		"1twotwoo3":   {"1", "2", "2", "3"},
		"1twoone3":    {"1", "2", "1", "3"},
		"one1twoone3": {"1", "1", "2", "1", "3"},
		"xtwone3four": {"2", "1", "3", "4"},
	}

	for k, v := range testCases {
		passed, output, err := testScanner(k, ScanNumeralsAndNumberWords, v)
		if err != nil {
			t.Errorf("Error: %v\n", err)
		}
		if !passed {
			t.Errorf("Expected `%v`, got `%v`", v, output)
		}
	}
}

func testScanner(input string, splitFunc bufio.SplitFunc, expectedOutput []string) (bool, []string, error) {
	var output []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(splitFunc)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if scanner.Err() != nil {
		return false, nil, scanner.Err()
	}
	return reflect.DeepEqual(output, expectedOutput), output, nil
}
