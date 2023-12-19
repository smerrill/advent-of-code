package main

import (
	"reflect"
	"testing"
)

const input1 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestGamesPossible(t *testing.T) {
	idsSum, err := GetPossibleGameIdSums(input1)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	expectedIdsSum := 1 + 2 + 5
	if idsSum != expectedIdsSum {
		t.Errorf("Expected `%v`, got `%v`", expectedIdsSum, idsSum)
	}

	idsPower, err := GetPowerOfMinCubes(input1)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	expectedIdsPower := 2286
	if idsPower != expectedIdsPower {
		t.Errorf("Expected `%v`, got `%v`", expectedIdsPower, idsPower)
	}
}

var testCases = map[string]GameResult{
	`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green`:           {1, []CubeCounts{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}}},
	`Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue`: {2, []CubeCounts{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}}},
}

func TestParseLine(t *testing.T) {
	for k, v := range testCases {
		output, err := ParseLine(k)

		if err != nil {
			t.Errorf("Error: %v\n", err)
		}
		if !reflect.DeepEqual(output, v) {
			t.Errorf("Expected `%v`, got `%v`", v, output)
		}
	}
}
