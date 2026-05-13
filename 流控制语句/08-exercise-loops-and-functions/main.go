package main

import (
	"fmt"
	"math"
)

// Sqrt 用牛顿法实现平方根函数
func Sqrt(x float64) float64 {
	// 初始猜测值 z = 1.0
	z := 1.0

	// 循环：直到变化足够小（停止变化）
	for i := 0; i < 10; i++ { // 先做10次迭代
		z -= (z*z - x) / (2 * z)
		fmt.Printf("迭代 %d: z = %.10f\n", i+1, z)
	}
	return z
}

func main() {
	x := 2.0
	fmt.Printf("\n自己实现的 Sqrt(%v) = %v\n", x, Sqrt(x))
	fmt.Printf("系统 math.Sqrt(%v) = %v\n", x, math.Sqrt(x))
}
