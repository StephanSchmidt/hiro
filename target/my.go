package main

import (
	"fmt")

func add(a int, b int) <- chan int {
res := make(chan int)
go func(){
 defer close(res)
res <- a+b
}()
return res
}
func main() {var a = (<- add(2,3))
fmt.Println(a)
}
