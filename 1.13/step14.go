package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	f0, f1, fn := 0, 1, 0
	i := 0
	for i < a {
		fn = f0 + f1
		f0, f1 = f1, fn
		i++
		if fn == a {
			fmt.Println(i + 1)
			return
		} else if fn > a || i > a {
			fmt.Println(-1)
			return
		}
	}
}
