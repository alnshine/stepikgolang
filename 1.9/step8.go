package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var n1, n2, n3, n4, n5, n6 int
	n1 = n / 100000
	n2 = (n / 10000) % 10
	n3 = (n / 1000) % 10
	n4 = (n / 100) % 10
	n5 = (n / 10) % 10
	n6 = n % 10
	if n1+n2+n3 == n4+n5+n6 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
