package main

import "fmt"

func main() {
	grades := []int{73, 67, 38, 33}
	a := len(grades)
	for i := 0; i < a; i++ {
		if grades[i] > 35 {
			b := grades[i] % 5
			if b == 3 {
				grades[i] = grades[i] + 2
			} else if b == 4 {
				grades[i] = grades[i] + 1
			} else {
				continue
			}
		}
	}
	fmt.Println(grades)
}
