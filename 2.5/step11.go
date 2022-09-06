package main

import "fmt"

func main() {
	var s1 string
	fmt.Scan(&s1)
	r := []rune(s1)
	var count int
	for i := 0; i < len(r); i++ {
		gg := r[i]
		for j := 0; j < len(r); j++ {
			if gg == r[j] {
				count++
			}
		}
		if count < 2 {
			fmt.Print(string(r[i]))
		}
		count = 0
	}
}
