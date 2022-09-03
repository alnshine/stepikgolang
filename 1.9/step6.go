package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a, b, c int
	a = n / 100
	c = n % 10
	n /= 10
	b = n % 10
	if a == b || a == c || b == c {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
