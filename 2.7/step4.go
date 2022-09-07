package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	max := r[0]
	for i := 0; i < len(r); i++ {
		if r[i] > max {
			max = r[i]
		}
	}
	fmt.Println(string(max))
}
