package main

import "testing"

const input1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const input2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestGetCalibration(t *testing.T) {
	sum, err := GetCalibrations(input1)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if sum != 142 {
		t.Errorf("Expected 142, got %d", sum)
	}
	sum, err = GetCalibrationsWithWords(input2)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if sum != 281 {
		t.Errorf("Expected 281, got %d", sum)
	}
}
