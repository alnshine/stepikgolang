package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	n10, n100 := n%10, n%100

	var kkk string
	if n10 == 1 && n100 != 11 {
		kkk = "a"
	} else if (n10 == 2 && n100 != 12) || (n10 == 3 && n100 != 13) || (n10 == 4 && n100 != 14) {
		kkk = "y"
	}

	fmt.Printf("%d korov%s\n", n, kkk)
}
