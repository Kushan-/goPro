package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}
	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not too big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x > 5 && x < 15 { //logical OR
		fmt.Println("x is out right")
	}

	a, b := 11.0, 12.0
	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more half b")
	}
}
