package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	_max := 0
	for _, num := range nums {
		//fmt.Println(num)
		if _max < num {
			_max = num
		}
	}
	fmt.Println(_max)
}
