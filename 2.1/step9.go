package main

import "fmt"

func main() {
	fmt.Println(fibonacci(4))
}
func fibonacci(n int) int {
	f0, f1, fn := 0, 1, 0
	for i := 0; i < n; i++ {
		fn = f0 + f1
		f0, f1 = f1, fn
	}
	return f0
}
