package main

import "fmt"

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

func main() {
	if 5 != 5 {
		fmt.Println("5 == 5")
	} else {
		fmt.Println("5 != 5")
	}

	if 3 < 4 {
		fmt.Println("3 < 4")
	}else if 4 == 4 {
		fmt.Println("4 == 4")
	}else{
		fmt.Println("4 != 4")
	}
}