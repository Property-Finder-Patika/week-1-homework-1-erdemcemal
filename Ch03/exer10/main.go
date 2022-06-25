package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345"))

}

// Exercise 3.10 Exercise 3.10: Write a non-recursive version of comma, using bytes.
// Buffer instead of string concatenation.
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
