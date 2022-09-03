package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	count := 0
	for _, i := range a {
		if i > 0 {
			count++
		}
	}
	fmt.Println(count)

}
