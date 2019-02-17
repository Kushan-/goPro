package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	wtr io.Writer
}

// Area returns the area of the square
func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i, c := range p {
		fmt.Printf("c = %v\n", c)
		fmt.Printf("i = %v\n", i)
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}
	return c.wtr.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello There i am USing Go")
}
