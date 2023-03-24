package main

import (
	"fmt"
)

func main() {
	//challenge 1
	num := 21
	j := true
	hexa := 15
	yr := "\u042f"
	ya := 1071
	floating := 123.456

	fmt.Printf("%d\n%T\n%%\n%t\n%t\n%q\n%d\n%o\n%x\n%X\n%U\n%.6f\n%E", num, num, j, !j, yr, num, num, hexa, hexa, ya, floating, floating)
}
