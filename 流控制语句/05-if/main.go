package main

import (
	// fmt 用来把结果输出到控制台，也用来把数字转换成字符串。
	"fmt"
	// math 提供数学相关函数，这里使用 math.Sqrt 计算平方根。
	"math"
)

// sqrt 根据传入的数字 x 返回它的平方根结果。
// 返回值类型是 string，是为了同时表示普通平方根和负数平方根的虚数形式。
func sqrt(x float64) string {
	// if 用来根据条件决定是否执行代码块。
	// Go 的 if 条件不需要小括号，但条件后面的花括号是必须的。
	// 当 x 小于 0 时，说明 x 是负数，普通实数范围内不能直接计算平方根。
	if x < 0 {
		// 对负数求平方根时，先把它变成正数 -x，再递归调用 sqrt 计算正数部分。
		// 最后拼接 "i"，表示这是一个虚数结果。
		// 例如 sqrt(-4) 会变成 sqrt(4) + "i"，最终得到 "2i"。
		return sqrt(-x) + "i"
	}

	// 如果 x 不是负数，就直接使用 math.Sqrt 计算平方根。
	// math.Sqrt 返回 float64，fmt.Sprint 会把它转换成字符串，方便作为函数返回值。
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	// 调用 sqrt(2) 演示普通正数的平方根。
	// 调用 sqrt(-4) 演示 if 分支处理负数并返回虚数形式。
	fmt.Println(sqrt(2), sqrt(-4))
}
