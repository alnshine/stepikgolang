package main

import "fmt"

func main() {
	var n, h, m int
	fmt.Scan(&n)
	h = n / 3600
	n -= h * 3600
	m = n / 60
	fmt.Printf("It is %v hours %v minutes.", h, m)
}
