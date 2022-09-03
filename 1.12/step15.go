package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i, x := range a {
		if i%2 == 0 {
			fmt.Print(x, " ")
		}
	}
}
