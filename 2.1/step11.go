package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}
func sumInt(lol ...int) (int, int) {
	var sum int
	for _, elem := range lol {
		sum += elem
	}
	return len(lol), sum
}
