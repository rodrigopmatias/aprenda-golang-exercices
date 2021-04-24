package main

import "fmt"

type bancon int

var x bancon

func main() {
	fmt.Printf("%T(%v)\n", x, x)

	x := bancon(42)

	fmt.Printf("%T(%v)\n", x, x)
}
