package main

import (
	"fmt"
	"math"
)

var k, p, v float64

func main() {
	fmt.Scan(&k, &p, &v)
	fmt.Println(T())
}
func T() float64 {
	t := 6 / W()
	return t
}
func W() float64 {
	res := k / M()
	w := math.Sqrt(res)
	return w
}
func M() float64 {
	return p * v
}
