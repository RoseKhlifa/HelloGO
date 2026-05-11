package main

import (
	"fmt"
)

// func 函数名(参数名 参数类型,参数名 参数类型)返回值类型{}
func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println(add(42, 13))
}
