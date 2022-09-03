package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(a[3])
}
