package main

import "fmt"

func main() {
	var n, n1, n2 int
	fmt.Scan(&n)
	for n >= 1 {
		n1 = n % 10
		n /= 10
		n2 += n1
	}
	n = n2 % 10
	n1 = n2 / 10
	fmt.Println(n + n1)
}
