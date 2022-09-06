package main

import "fmt"

func main() {
	x := 5
	y := 8
	test(&x, &y)
}
func test(x1 *int, x2 *int) {
	var a int
	a = *x1 * *x2
	fmt.Println(a)
}
