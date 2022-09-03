package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)
	var h int
	h = d / 30
	var m int
	m = 2 * (d % 30)
	fmt.Println("It is", h, "hours", m, "minutes.")
}
