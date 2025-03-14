// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
)

const comma_symbol = ","

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
		fmt.Printf("  %s\n", commaNonRecursive(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaNonRecursive(s string) string {
	var buf bytes.Buffer

	var lastPart = len(s) % 3

	if lastPart == 0 {
		return s
	}

	buf.WriteString(s[:lastPart])

	for i := lastPart; i < len(s); i += 3 {
		buf.WriteString(comma_symbol)
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

//!-
