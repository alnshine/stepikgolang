package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	stroka := string(text)
	Right := "Right"
	Wrong := "Wrong"
	r := []rune(stroka)
	if unicode.IsUpper(r[0]) && strings.HasSuffix(string(r), ".\n") {
		fmt.Println(Right)
	} else {
		fmt.Println(Wrong)
	}
}
