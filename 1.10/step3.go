package main

import "fmt"

func main() {
	var a, b, s int
	fmt.Scan(&a, &b)
	s = 0
	for i := a; i < b; i++ {
		s += i
	}
	fmt.Println(s + b)
}
