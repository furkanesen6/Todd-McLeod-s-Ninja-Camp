package main

import "fmt"

func main() {
	x := 7

	switch x {
	case 4:
		fmt.Println("1. case")
	case 9:
		fmt.Println("2. case")
	case 7:
		fmt.Println("3. case")
	default:
		fmt.Println("default")
	}
}