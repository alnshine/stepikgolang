package main

import "fmt"

func main() {
	var y, u, u1, u2 int
	fmt.Scan(&y)
	u = y % 400
	u1 = y % 4
	u2 = y % 100
	if u == 0 {
		fmt.Println("YES")
	} else if u1 == 0 && u2 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
