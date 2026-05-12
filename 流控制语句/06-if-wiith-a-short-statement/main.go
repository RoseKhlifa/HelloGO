package main

import (
	// fmt 用来把结果输出到控制台。
	"fmt"
	// math 提供数学计算函数，这里使用 math.Pow 计算 x 的 n 次方。
	"math"
)

// pow 计算 x 的 n 次方，并把结果和上限 lim 进行比较。
// 如果计算结果小于 lim，就返回真实计算结果；否则返回 lim。
func pow(x, n, lim float64) float64 {
	// Go 的 if 可以在条件判断前写一个简单语句，格式是：
	// if 初始化语句; 条件表达式 {
	//     条件为 true 时执行这里
	// }
	//
	// 这里的 v := math.Pow(x, n) 是 if 的短语句：
	// 1. 先计算 x 的 n 次方。
	// 2. 把计算结果保存到变量 v 中。
	// 3. 再判断 v < lim 是否成立。
	//
	// 注意：变量 v 只在这个 if 语句内部可用，包括 if 后面的 else 分支。
	// 离开 if 以后，外面不能再使用 v。
	if v := math.Pow(x, n); v < lim {
		// 如果 v 小于上限 lim，说明计算结果没有超过限制，直接返回 v。
		return v
	}

	// 如果 if 条件不成立，说明 v 大于或等于 lim。
	// 这时返回 lim，相当于把结果限制在最大值 lim 以内。
	return lim
}

func main() {
	// fmt.Println 可以一次打印多个值，中间会自动用空格分隔。
	fmt.Println(
		// 3 的 2 次方是 9，9 小于 10，所以返回 9。
		pow(3, 2, 10),
		// 3 的 3 次方是 27，27 不小于 20，所以返回上限 20。
		pow(3, 3, 20),
	)
}
