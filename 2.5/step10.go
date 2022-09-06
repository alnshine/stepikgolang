package main

import "fmt"

func main() {
	var s1 string
	fmt.Scan(&s1)
	r1 := []rune(s1)
	var r2 []rune
	for i := 0; i < len(r1); i++ {
		if i%2 != 0 {
			r2 = append(r2, r1[i])
		}
	}
	fmt.Println(string(r2))
}
