package main

import "fmt"

func main() {

	a := 500
	fmt.Printf("%d\t%b\t%#x\n",a,a,a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x",b,b,b)
}
