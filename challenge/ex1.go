package main

import "fmt"

func main() {
	for num := 1; num < 21; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("fizz buzz")
		} else if num%5 == 0 {
			fmt.Println("buzz")
		} else if num%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(num)
		}

	}
}
