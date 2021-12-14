package main

import "fmt"

func newInt1() *int { return new(int) }
func newInt2() *int {
	var i int
	return &i
}

func main() {

	p := new(chan int)
	fmt.Printf("\n%T", p)

	c := make(chan int)
	fmt.Printf("\n%T", c)

	s := make([]string, 10)
	fmt.Printf("\n%T", s)
}
