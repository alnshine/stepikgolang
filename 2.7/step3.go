package main

import "fmt"

func main() {
	var s1 string
	fmt.Scan(&s1)
	r := []rune(s1)
	for i := 0; i < len(r); i++ {
		fmt.Print(string(r[i]))
		if i != len(r)-1 {
			fmt.Print(string('*'))
		}
	}
}
