package cmd

import (
	"io"
	"strings"
)

// countWords returns the number of whitespace-separated words in r.
func countWords(r io.Reader) (int, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}

	words := strings.Split(string(content), " ")
	return len(words), nil
}
