package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}
	if b == 0 {
		fmt.Println("ошибка")
		return
	}
	res, err1 := divide(a, b)
	if err1 == nil {
		fmt.Println(res)
	} else {
		fmt.Println("ошибка")
	}
}
func divide(a, b int) (int, error) {
	return a / b, nil
}
