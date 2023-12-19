package adventofcode

import (
	"io"
	"os"
)

func ReadFileIntoString(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ReadStdinToString() (string, error) {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		return "", err
	}
	return string(stdin), nil
}
