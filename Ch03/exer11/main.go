package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(addComma("-12345.15"))

}

// Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an optional sign.
func addComma(s string) string {
	var start, end int

	var startStr, endStr string
	if strings.HasPrefix(s, "-") {
		start = 1
		startStr = s[:start]
	} else {
		start = 0
	}

	if strings.Contains(s, ".") {
		end = strings.Index(s, ".")
		endStr = s[end:]
	} else {
		end = len(s)
	}
	s = s[start:end]

	return startStr + comma(s) + endStr
}

func comma(s string) string {
	var buf bytes.Buffer
	i := (3 - len(s)%3) % 3
	for _, v := range s {
		if i == 3 {
			buf.WriteByte(',')
			i = 0
		}
		buf.WriteRune(v)
		i++
	}
	return buf.String()
}
