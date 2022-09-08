package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[int]int{}
	var n int
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		if _, ok := m[n]; !ok {
			m[n] = work(n)
		}
		n = m[n]
		fmt.Print(n)
		fmt.Print(" ")
	}
}
func work(n int) int {
	if n > 3 {
		time.Sleep(time.Millisecond * 500)
		return n + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return n - 1
	}
}
