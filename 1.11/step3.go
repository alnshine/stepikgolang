package main

import "fmt"

func main() {
	var n float32
	fmt.Scan(&n)
	if n > 0 && n <= 10000 {
		n *= n
		fmt.Printf("%.4f", n)
	} else if n > 10000 {
		fmt.Printf("%e", n)
	} else {
		fmt.Printf("число %2.2f не подходит", n)
	}
}
