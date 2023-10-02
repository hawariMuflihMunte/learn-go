package main

import "fmt"

func main() {
	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You can be better!")
		}
	}
}
