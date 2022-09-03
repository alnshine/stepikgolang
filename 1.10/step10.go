package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	for _, a1 := range a {
		for _, b1 := range b {
			if a1 == b1 {
				fmt.Print(string(a1) + " ")
			}
		}
	}
}
