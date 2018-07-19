// Package cleandoc provides an implementation of Python's inspect.cleandoc.
package cleandoc

import (
	"math"
	"strings"
	"unicode"
)

// Cleandoc removes leading indentation from strings.
//
// Any whitespace that can be uniformly removed from the second line
// onwards is removed.
func Cleandoc(doc string) string {

	// This code is a direct translation of inspect.cleandoc, covered under the PSF license
	lines := strings.Split(strings.Replace(doc, "\t", "        ", -1), "\n")

	// Find minimum indentation of any non-blank lines after first line.
	margin := math.MaxInt32
	for _, line := range lines[1:] {
		content := len(strings.TrimLeftFunc(line, unicode.IsSpace))
		if content > 0 {
			indent := len(line) - content
			if margin > indent {
				margin = indent
			}
		}
	}
	// Remove indentation.
	if len(lines) != 0 {
		lines[0] = strings.TrimLeftFunc(lines[0], unicode.IsSpace)
	}
	if margin < math.MaxInt32 {
		for i, line := range lines[1:] {
			if len(line) < margin {
				lines[i+1] = ""
			} else {
				lines[i+1] = line[margin:]
			}
		}
	}
	// Remove any trailing or leading blank lines.
	for len(lines) != 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	for len(lines) != 0 && len(lines[0]) == 0 {
		lines = lines[1:]
	}
	return strings.Join(lines, "\n")
}
