package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	gip(a, b)
}
func gip(a, b float64) {
	var res float64
	res = a*a + b*b
	fmt.Println(math.Sqrt(res))
}
