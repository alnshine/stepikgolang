package main

import "fmt"

func main() {
	var n, n1, count int
	fmt.Scan(&n)
	number := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&n1)
		number[i] = n1
	}
	min := number[0]
	for _, i := range number {
		if i < min {
			min = i
		}
	}
	for _, x := range number {
		if x == min {
			count++
		}
	}
	fmt.Println(count)
}
