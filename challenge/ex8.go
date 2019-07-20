/*

Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.




*/




package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	//return map["%q", strings.Fields(m)]

	fmt.Println(m)
	count := 0
	for _, v := range strings.Fields(s) {
		//fmt.Printf("%d = %s\n", i, v)
		if _, ok := m[v]; ok{
			fmt.Printf("excecuting %d = %s\n", count, v)
			
			m[v] = count+1

		}else{
			count = 1
			m[v] = count
		}
		

		
	}
	fmt.Println(m)
	return m
}

func main() {
	wc.Test(WordCount)
}
