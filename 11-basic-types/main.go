package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	s := "你好！"
	//遍历string字符正确方式
	for i, r := range s {
		fmt.Printf("byte index %d:%c\n", i, r)
	}
	//修改string字符的正确方式
	b := []rune(s)
	b[0] = 'n'
	s = string(b)
	for i, r := range s {
		fmt.Printf("byte index %d:%c\n", i, r)
	}
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
