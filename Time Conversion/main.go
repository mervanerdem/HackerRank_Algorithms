package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "07:05:45PM"

	a := len(s)
	b := string(s[a-2])
	if b == "P" {
		b := int(s[0]) - 48
		c := int(s[1]) - 48
		d := b*10 + c + 12
		b = d / 10
		c = d - b*10

		str := strconv.Itoa(b)
		x := string(s[0])
		s = strings.Replace(s, x, str, 1)

		str = strconv.Itoa(c)
		y := string(s[1])
		s = strings.Replace(s, y, str, 1)

		s = strings.Replace(s, "PM", "", -1)
	}
	fmt.Println(s)
}
