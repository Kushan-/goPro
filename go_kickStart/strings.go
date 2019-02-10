package main

import "fmt"

func main() {
	book := "The pig is nasty"
	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// unit8 is a byte
	//Strings in go are immutable

	fmt.Println(book[4:7])
	fmt.Println(book[5:])
	fmt.Println(book[:7]) //after 4

	//concatenate
	fmt.Println("t" + book[1:]) //after 4

	//Multi line
	poem := `
		Some nasty peoem is down the road
		red i grewy
		blah blah blah
		^^^^^
	`
	fmt.Println(poem) //after 4

}
