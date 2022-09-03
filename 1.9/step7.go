package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 9 {
		fmt.Println(n)
	} else if n > 9 && n < 100 {
		n /= 10
		fmt.Println(n)
	} else if n > 100 && n < 1000 {
		n /= 100
		fmt.Println(n)
	} else if n > 1000 && n < 10000 {
		n /= 1000
		fmt.Println(n)
	} else {
		n /= 10000
		fmt.Println(n)
	}
}
