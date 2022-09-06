package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	if len(r) < 5 {
		fmt.Println("Wrong password")
		return
	}
	for i := 0; i < len(r); i++ {
		if check(r[i]) {
			fmt.Println("Wrong password")
			return
		}
	}
	fmt.Println("Ok")
}
func check(r rune) bool {
	if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return false
	}
	return true
}
