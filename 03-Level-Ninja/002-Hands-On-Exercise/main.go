package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {

	for i := 65; i <= 90; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
