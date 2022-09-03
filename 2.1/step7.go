package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}
func minimumFromFour() int {
	var min int
	fmt.Scan(&min)
	for i := 0; i < 3; i++ {
		var moment int
		fmt.Scan(&moment)
		if moment < min {
			min = moment
		}
	}
	return min
}
