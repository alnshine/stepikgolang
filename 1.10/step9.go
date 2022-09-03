package main

import "fmt"

func main() {
	var x, p, y, g int
	fmt.Scan(&x, &p, &y)
	for ; x < y; g++ {
		x += x * p / 100
	}
	fmt.Println(g)

}
