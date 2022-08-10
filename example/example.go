package main

import (
	"fmt"
)

type x struct {
	y int32
}

func s_add(a int, b int) int {
	return a + b
}

func add(a int, b int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		// body
		ret := a + b
		// end body
		res <- ret
	}()
	return res
}

func main() {
	c := <-add(1, 2)
	fmt.Println(c)
}
