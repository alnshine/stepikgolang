package main

import "fmt"

func main() {
	var n int
	a := 1
	a2 := 1
	fmt.Scan(&n)
	for a2 <= n {
		if a2 <= n {
			fmt.Print(a2, " ")
		}
		a *= 2
		a2 = a
	}
}
