package main

import "fmt"

func main() {
	var n, n1, n2, n3 int
	fmt.Scan(&n)
	n1 = n / 100
	n2 = n / 10
	n2 = n2 % 10
	n3 = n % 10

	n3 *= 100
	n2 *= 10
	fmt.Println(n1 + n2 + n3)
}
