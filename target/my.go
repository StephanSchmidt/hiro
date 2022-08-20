package main

import (
	"fmt")

func add(x int, y int) <- chan int {
res := make(chan int)
go func(){
 defer close(res)
res <- x+y
}()
return res
}
func main() {
a := make(chan any)
go func() {
defer close(a)
a <- (<- add(2,3))+2}()

b := make(chan any)
go func() {
defer close(b)
b <- (<- add(2,3))}()

fmt.Println((<- add(2,3)))
var _b = <- b
fmt.Println(_b)
var _a = <- a
fmt.Println(_a)
}
