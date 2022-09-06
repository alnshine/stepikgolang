package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var slovo string
	fmt.Scan(&slovo)
	dlina := utf8.RuneCountInString(slovo)
	r := []rune(slovo)
	for i := 0; i < dlina/2; i++ {
		if r[i] != r[len(r)-i-1] {
			fmt.Println("Нет")
			return
		}
	}
	fmt.Println("Палиндром")
}
