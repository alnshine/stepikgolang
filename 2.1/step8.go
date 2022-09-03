package main

import "fmt"

func main() {
	fmt.Println(vote(0, 0, 1))
}
func vote(x int, y int, z int) int {
	var ed int
	var nol int
	if x == 0 {
		nol++
	} else {
		ed++
	}
	if y == 0 {
		nol++
	} else {
		ed++
	}
	if z == 0 {
		nol++
	} else {
		ed++
	}
	if nol > ed {
		return 0
	} else {
		return 1
	}
}
