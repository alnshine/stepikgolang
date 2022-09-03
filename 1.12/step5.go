workArray := [10]uint8{}
for i := range workArray {
	fmt.Scan(&workArray[i])
}
var a, b uint
for i := 0; i < 3; i++ {
	fmt.Scan(&a, &b)
	workArray[a], workArray[b] = workArray[b], workArray[a]
}
for _, x := range workArray {
	fmt.Print(x, " ")
}
fmt.Printf("type ok")
return
