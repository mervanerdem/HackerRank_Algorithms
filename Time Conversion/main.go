package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "02:05:45AM"

	a := len(s)
	b := string(s[a-2])

	if string(s[0]) == "1" && string(s[1]) == "2" {
		if b == "P" {
			s = strings.Replace(s, "PM", "", -1)
		} else if b == "A" {
			s = strings.Replace(s, "12", "00", -1)
			s = strings.Replace(s, "AM", "", -1)
		} else {
			fmt.Println("ERROR")
		}
	} else if b == "P" {
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
	} else {
		s = strings.Replace(s, "AM", "", -1)
	}
	fmt.Println(s)
}
