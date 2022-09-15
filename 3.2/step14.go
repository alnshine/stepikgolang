package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")

	r := []rune(text)
	var r1, r2 []rune
	for i := 0; i < len(r); i++ {
		if r[0] == ';' {
			r = r[1:]
			continue
		}
		r1 = append(r1, r[0])
		r = r[1:]
	}
	for _, j := range r {
		r2 = append(r2, j)
	}
	r1 = antip(r1)
	r2 = antip(r2)
	r1 = point(r1)
	r2 = point(r2)
	result1, _ := strconv.ParseFloat(string(r1), 64)
	result2, _ := strconv.ParseFloat(string(r2), 64)
	res := result1 / result2
	s := fmt.Sprintf("%.4f", res)
	fmt.Println(s)
}
func antip(r []rune) []rune {
	var rg []rune
	for i := 0; i < len(r); i++ {
		if r[0] == ' ' {
			r = r[1:]
		}
		rg = append(rg, r[0])
		r = r[1:]
	}
	return rg
}
func point(r []rune) []rune {
	for i := 0; i < len(r); i++ {
		if r[i] == ',' {
			r[i] = '.'
		}
	}
	return r
}
