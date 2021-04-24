package main

import "fmt"

type bancon int

var x bancon
var y int

func main() {
	fmt.Printf("%T(%v)\n", x, x)

	x := bancon(42)
	fmt.Printf("%T(%v)\n", x, x)

	y = int(x)
	fmt.Printf("%T(%v)\n", y, y)
}
