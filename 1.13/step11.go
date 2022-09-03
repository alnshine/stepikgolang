package main

import "fmt"

func main() {
	var a, b, res int
	fmt.Scan(&a, &b)
	for i := a; i < b; i++ {
		if b%7 == 0 {
			res = b
		} else if i%7 == 0 {
			res = i
		}
	}
	if res < a {
		fmt.Println("NO")
	} else {
		fmt.Println(res)
	}

}
