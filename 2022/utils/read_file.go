package utils

import (
	"os"
)

func ReadFile(path string) (string, error) {
	contents, error := os.ReadFile(path)

	if error != nil {
		return "", error
	}

	return string(contents), nil
}
