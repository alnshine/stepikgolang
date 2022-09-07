package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	for _, v := range r {
		chislo, _ := strconv.Atoi(string(v))
		chislo *= chislo
		fmt.Print(chislo)
	}
}
