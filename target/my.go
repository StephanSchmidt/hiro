package main

import (
	"fmt"
)

func add(a int, b int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		res <- a + b
	}()
	return res
}
func main() {
	var b_done = 3
	a := make(chan any)
	go func() {
		defer close(a)
		a <- (<-add(2, 3)) + 2
	}()

	fmt.Println((<-add(2, 3)))
	var a_done = <-a
	fmt.Println(a_done)
	fmt.Println(b_done)
}
