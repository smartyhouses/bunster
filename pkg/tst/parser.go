package tst

import (
	"errors"
	"fmt"
	"strings"
)

// TestCase represents a single test case with input and output
type Test struct {
	Label  string
	Input  string
	Output string
}

const (
	START  = iota
	INPUT  = iota
	OUTPUT = iota
)

// Parse reads from an io.Reader and parses the test case format
func Parse(in string) ([]Test, error) {
	var lines = splitIntoLines(in)

	var step = START
	var tests []Test
	var test Test

	for _, line := range lines {
		if step == START {
			if strings.TrimSpace(line) == "" {
				continue
			}

			label := strings.TrimSpace(line)

			label, ok := strings.CutPrefix(label, "#(TEST:")
			if !ok {
				return nil, fmt.Errorf("bad test syntax, coundl't find test header '#(TEST: ...)', found %q", label)
			}

			label, ok = strings.CutSuffix(label, ")")
			if !ok {
				return nil, fmt.Errorf("bad test syntax, unclosed test header '#(TEST: ...)'")
			}

			test.Label = strings.TrimSpace(label)
			step = INPUT
			continue
		}

		if step == INPUT {
			if strings.TrimSpace(line) == "#(RESULT)" {
				step = OUTPUT
				continue
			}

			test.Input += line
			continue
		}

		if step == OUTPUT {
			if strings.TrimSpace(line) == "#(ENDTEST)" {
				step = START
				tests = append(tests, test)
				continue
			}

			test.Output += line
			continue
		}
	}

	if step == INPUT {
		return nil, errors.New("bad test syntax, coundl't find #(RESULT) section")
	}
	if step == OUTPUT {
		return nil, errors.New("bad test syntax, unclosed test, missing '#(ENDTEST)'")
	}

	return tests, nil
}

// SplitIntoLines splits a string into a slice of lines, preserving newlines.
func splitIntoLines(input string) []string {
	var lines []string
	var currentLine strings.Builder

	for _, char := range input {
		currentLine.WriteRune(char)
		if char == '\n' {
			lines = append(lines, currentLine.String())
			currentLine.Reset()
		}
	}

	// Append the last line if it doesn't end with a newline
	if currentLine.Len() > 0 {
		lines = append(lines, currentLine.String())
	}

	return lines
}
