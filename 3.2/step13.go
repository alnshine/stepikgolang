// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var s1, s2 string
// 	fmt.Scan(&s1)
// 	fmt.Scan(&s2)
// 	fmt.Println(adding(s1, s2))
// }
// func adding(s1, s2 string) int64 {
// 	r1 := []rune(s1)
// 	r2 := []rune(s2)
// 	var rg1, rg2 []rune
// 	for _, j := range r1 {
// 		if j >= '0' && j <= '9' {
// 			rg1 = append(rg1, j)
// 		}
// 	}
// 	for _, j := range r2 {
// 		if j >= '0' && j <= '9' {
// 			rg2 = append(rg2, j)
// 		}
// 	}
// 	res1, _ := strconv.Atoi(string(rg1))
// 	res2, _ := strconv.Atoi(string(rg2))
// 	return int64(res1 + res2)
// }
