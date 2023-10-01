package main

import "fmt"

func main() {
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	fmt.Print(square, "\n")
	fmt.Print(isToday, "\n")
	fmt.Print(numeric, "\n")
	fmt.Print(floatNum, "\n")

	const three, four, five string = "tiga", "empat", "lima"
	fmt.Print(three, " ", four, " ", five)
}
