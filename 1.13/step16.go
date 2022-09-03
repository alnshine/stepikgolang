package main

import "fmt"

func main() {
	var n, lenght, n1, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	n1 = n
	for n1 > 0 {
		lenght++
		n1 /= 10
	}
	a1 := make([]int, lenght, lenght)
	n1 = n
	for i := 0; i < lenght; {
		n1 = n % 10
		a1[i] = n1
		n /= 10
		i++
	}
	a2 := make([]int, 0, 0)
	for i := range a1 {
		if a1[i] != k {
			a2 = append(a2, a1[i])
		}
	}
	var reserve []int
	for i := len(a2) - 1; i >= 0; i-- {
		reserve = append(reserve, a2[i])
	}
	for i := range reserve {
		fmt.Print(reserve[i])
	}
}
