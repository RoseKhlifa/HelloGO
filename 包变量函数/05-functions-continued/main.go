package main

import (
	"fmt"
)

// func 函数名 参数类型一致 省略一个参数类型
func add(x, y int) int {
	return x + y
}
func main() {
	fmt.Println(add(42, 13))
}
