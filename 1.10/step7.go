package main

import "fmt"

func main() {
	var a, b, c, i int
	fmt.Scan(&a, &c, &b)
	for i <= a {
		if i%c == 0 && i%b != 0 {
			fmt.Println(i)
			break
		}
		i++
	}
}
