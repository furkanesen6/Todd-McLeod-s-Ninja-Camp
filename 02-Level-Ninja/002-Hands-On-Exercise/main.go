package main

import "fmt"

var x int
var y string
var z bool

func main() {

	a := (4==4)
	b := (4 <= 5)
	c := (3 >= 3)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)

	fmt.Println(a,b,c,d,e,f)
}
