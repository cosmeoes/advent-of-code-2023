package helpers

import (
	"os"
	"strings"
)

func ReadFileLines(path string) ([]string, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return []string{}, err
    }

    return strings.Split(strings.TrimSpace(string(content)), "\n"), nil
}
