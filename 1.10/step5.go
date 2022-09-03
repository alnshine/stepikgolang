package main

import "fmt"

func main() {
	var a, b, c int
	b = 0
	c = 0
	for {
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		if b < a {
			c = 0
			b = a
		}
		if a == b {
			c++
		}
	}
	fmt.Println(c)
}
