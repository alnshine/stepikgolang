package main

import "fmt"

func main() {
	var a, b, s int
	s = 0
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b > 9 && b < 100 && b%8 == 0 {
			s += b
		}
	}
	fmt.Println(s)
}
