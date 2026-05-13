# Slices 切片

这一节代码：

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

运行结果：

```text
[3 5 7]
```

切片是 Go 里最常用的数据结构之一。和数组相比，切片更灵活，也更容易在日常开发中出现“共享底层数组”的坑。

## 切片是什么

切片是对数组某一段的引用。

数组是真正存放数据的地方。

切片本身不直接保存所有元素，它保存的是对底层数组某一段的描述。

可以把切片理解成三部分：

1. 指向底层数组某个位置的指针。
2. 长度 `len`。
3. 容量 `cap`。

简化理解：

```text
slice = 指针 + 长度 + 容量
```

## 切片类型

数组类型带长度：

```go
[6]int
```

切片类型不带长度：

```go
[]int
```

所以：

```go
var s []int
```

表示 `s` 是一个 `int` 切片。

它可以表示 0 个、3 个、100 个 `int`，长度可以变化。

## 从数组创建切片

本节示例：

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4]
```

`primes` 是数组：

```text
索引:      0  1  2  3   4   5
primes:   2  3  5  7  11  13
```

`primes[1:4]` 表示从下标 `1` 开始，到下标 `4` 之前结束。

包含下标：

```text
1, 2, 3
```

不包含下标：

```text
4
```

所以结果是：

```text
[3 5 7]
```

切片区间规则：

```go
a[low:high]
```

含义：

```text
从 low 开始
到 high 之前结束
包含 low
不包含 high
```

这个规则和很多语言里的左闭右开区间一样：

```text
[low, high)
```

## 切片不是复制数组

这一点非常重要。

```go
s := primes[1:4]
```

不会复制 `primes[1]` 到 `primes[3]`。

它只是创建了一个切片，指向数组 `primes` 的这一段。

可以理解成：

```text
primes: [2 3 5 7 11 13]
            ↑     ↑
            s 引用这里：[3 5 7]
```

所以修改切片会影响底层数组。

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
s := primes[1:4]

s[0] = 100

fmt.Println(s)
fmt.Println(primes)
```

输出：

```text
[100 5 7]
[2 100 5 7 11 13]
```

因为 `s[0]` 对应的是 `primes[1]`。

## 多个切片可以共享同一个底层数组

```go
a := [4]int{1, 2, 3, 4}

s1 := a[0:2]
s2 := a[1:3]

s1[1] = 100

fmt.Println(s1)
fmt.Println(s2)
fmt.Println(a)
```

输出：

```text
[1 100]
[100 3]
[1 100 3 4]
```

因为：

```text
s1[1] 对应 a[1]
s2[0] 也对应 a[1]
```

所以改一个切片，另一个切片看到的内容也变了。

这是切片最核心、也最容易踩坑的地方。

## len 和 cap

切片有长度和容量。

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
s := primes[1:4]

fmt.Println(len(s))
fmt.Println(cap(s))
```

`s` 是：

```text
[3 5 7]
```

所以长度 `len(s)` 是：

```text
3
```

容量 `cap(s)` 是从切片起点到底层数组末尾的长度。

`s` 从 `primes[1]` 开始，底层数组到末尾还有：

```text
primes[1], primes[2], primes[3], primes[4], primes[5]
```

共 `5` 个元素。

所以：

```text
len(s) = 3
cap(s) = 5
```

## 切片表达式的省略写法

如果从开头开始，可以省略 `low`：

```go
a[:4]
```

等价于：

```go
a[0:4]
```

如果切到末尾，可以省略 `high`：

```go
a[2:]
```

等价于：

```go
a[2:len(a)]
```

如果整个都要：

```go
a[:]
```

等价于：

```go
a[0:len(a)]
```

## 从切片再切片

切片还能继续切。

```go
s := []int{2, 3, 5, 7, 11, 13}
t := s[1:4]
u := t[1:3]

fmt.Println(t)
fmt.Println(u)
```

输出：

```text
[3 5 7]
[5 7]
```

注意：`t` 和 `u` 仍然可能共享同一个底层数组。

## 切片字面量

可以直接创建切片：

```go
s := []int{2, 3, 5, 7, 11, 13}
```

注意这里没有写长度。

```go
[]int{...}
```

是切片。

```go
[6]int{...}
```

是数组。

```go
[...]int{...}
```

是让编译器推断长度的数组。

这三个很容易混：

```go
a := [3]int{1, 2, 3}   // 数组，类型是 [3]int
b := [...]int{1, 2, 3} // 数组，类型是 [3]int
c := []int{1, 2, 3}    // 切片，类型是 []int
```

## make 创建切片

`make` 常用来创建切片、map、channel。

创建长度为 3 的切片：

```go
s := make([]int, 3)
fmt.Println(s)
fmt.Println(len(s), cap(s))
```

输出：

```text
[0 0 0]
3 3
```

创建长度为 3、容量为 5 的切片：

```go
s := make([]int, 3, 5)
fmt.Println(s)
fmt.Println(len(s), cap(s))
```

输出：

```text
[0 0 0]
3 5
```

这里：

```go
make([]int, len, cap)
```

含义是：

```text
创建一个 int 切片
长度是 len
容量是 cap
```

长度表示当前能直接访问的元素数量。

容量表示在不重新分配底层数组的情况下最多能扩展到多长。

## append 追加元素

切片可以使用 `append` 追加元素。

```go
s := []int{1, 2, 3}
s = append(s, 4)

fmt.Println(s)
```

输出：

```text
[1 2 3 4]
```

注意：`append` 的返回值必须接住。

正确：

```go
s = append(s, 4)
```

错误：

```go
append(s, 4) // 编译错误：返回值没有使用
```

为什么必须接住？

因为 `append` 可能会创建新的底层数组，并返回一个新的切片。

## append 和底层数组扩容

如果切片容量够，`append` 可能直接在原底层数组后面追加。

如果容量不够，`append` 会分配新的底层数组，把旧数据复制过去，再追加新元素。

示例：

```go
s := make([]int, 0, 3)
fmt.Println(len(s), cap(s))

s = append(s, 1)
s = append(s, 2)
s = append(s, 3)
fmt.Println(s, len(s), cap(s))

s = append(s, 4)
fmt.Println(s, len(s), cap(s))
```

可能输出：

```text
0 3
[1 2 3] 3 3
[1 2 3 4] 4 6
```

最后一次追加超过了原容量，所以发生扩容。

容量增长策略是运行时实现细节，不要写依赖具体扩容倍数的代码。

## nil 切片

切片的零值是 `nil`。

```go
var s []int

fmt.Println(s == nil)
fmt.Println(len(s))
fmt.Println(cap(s))
```

输出：

```text
true
0
0
```

nil 切片可以安全 `append`：

```go
var s []int
s = append(s, 1)
fmt.Println(s)
```

输出：

```text
[1]
```

## nil 切片和空切片

nil 切片：

```go
var a []int
```

空切片：

```go
b := []int{}
c := make([]int, 0)
```

它们的长度都是 `0`：

```go
len(a) == 0
len(b) == 0
len(c) == 0
```

但：

```go
a == nil // true
b == nil // false
c == nil // false
```

日常使用时，大多数情况下它们都可以当作空列表。

但在 JSON 编码时可能有区别：

```go
var a []int      // nil
b := []int{}     // empty
```

编码结果可能是：

```json
null
[]
```

所以写 API 时要注意对方期望的是 `null` 还是 `[]`。

## 切片不能直接比较

切片只能和 `nil` 比较。

```go
var s []int
fmt.Println(s == nil) // 可以
```

不能比较两个切片：

```go
a := []int{1, 2}
b := []int{1, 2}

fmt.Println(a == b) // 错误
```

如果要比较内容，需要自己遍历，或者使用标准库辅助函数。

比如 Go 1.21+ 可以使用：

```go
slices.Equal(a, b)
```

## 遍历切片

### 1. 普通 for

```go
s := []int{10, 20, 30}

for i := 0; i < len(s); i++ {
	fmt.Println(i, s[i])
}
```

### 2. range

```go
for i, v := range s {
	fmt.Println(i, v)
}
```

如果只要值：

```go
for _, v := range s {
	fmt.Println(v)
}
```

如果只要下标：

```go
for i := range s {
	fmt.Println(i)
}
```

## range 的值是副本

这是 Go 切片里非常常见的坑。

```go
s := []int{1, 2, 3}

for _, v := range s {
	v = v * 10
}

fmt.Println(s)
```

输出仍然是：

```text
[1 2 3]
```

因为 `v` 是元素值的副本。

如果要修改原切片，应该用下标：

```go
for i := range s {
	s[i] = s[i] * 10
}

fmt.Println(s)
```

输出：

```text
[10 20 30]
```

## 切片作为函数参数

切片传给函数时，切片头会被复制。

切片头里包含指向底层数组的指针，所以函数里修改元素会影响外面。

```go
func change(s []int) {
	s[0] = 100
}

func main() {
	nums := []int{1, 2, 3}
	change(nums)
	fmt.Println(nums)
}
```

输出：

```text
[100 2 3]
```

但如果函数里对切片变量本身重新赋值，不会改变外面的切片变量。

```go
func appendOne(s []int) {
	s = append(s, 4)
}

func main() {
	nums := []int{1, 2, 3}
	appendOne(nums)
	fmt.Println(nums)
}
```

外面的 `nums` 不一定会变成 `[1 2 3 4]`。

更正确的写法是返回新切片：

```go
func appendOne(s []int) []int {
	return append(s, 4)
}

func main() {
	nums := []int{1, 2, 3}
	nums = appendOne(nums)
	fmt.Println(nums)
}
```

核心记忆：

> 函数可以通过切片修改底层数组的元素，但如果改变切片长度，通常要返回新切片。

## copy 复制切片内容

如果不想共享底层数组，可以复制一份。

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))

copy(dst, src)

dst[0] = 100

fmt.Println(src)
fmt.Println(dst)
```

输出：

```text
[1 2 3]
[100 2 3]
```

`copy(dst, src)` 会把 `src` 的元素复制到 `dst`。

复制数量是 `len(dst)` 和 `len(src)` 中较小的那个。

## 删除切片元素

Go 没有内置的 `remove` 函数，但可以用 `append` 拼接。

删除下标 `i` 的元素：

```go
s = append(s[:i], s[i+1:]...)
```

示例：

```go
s := []int{10, 20, 30, 40}
i := 1

s = append(s[:i], s[i+1:]...)
fmt.Println(s)
```

输出：

```text
[10 30 40]
```

注意：这种写法通常会复用原底层数组。

如果元素里保存的是指针，并且你很关心内存释放，删除时可能还需要把最后一个元素置零，避免旧引用残留。

## 插入元素

在下标 `i` 插入元素 `x`：

```go
s = append(s[:i], append([]int{x}, s[i:]...)...)
```

更清楚的写法：

```go
s = append(s, 0)
copy(s[i+1:], s[i:])
s[i] = x
```

示例：

```go
s := []int{10, 30, 40}
x := 20
i := 1

s = append(s, 0)
copy(s[i+1:], s[i:])
s[i] = x

fmt.Println(s)
```

输出：

```text
[10 20 30 40]
```

## 切片共享底层数组的坑

看这个例子：

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]
c := a[2:4]

b[1] = 100

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
```

输出：

```text
[1 2 100 4 5]
[2 100]
[100 4]
```

因为 `b[1]` 和 `c[0]` 都对应 `a[2]`。

所以使用切片时要时刻记住：

> 切片可能共享同一块底层数组。

如果你需要完全独立的数据，就用 `copy`。

## reslice 扩展切片

切片可以在容量范围内重新切片。

```go
a := []int{1, 2, 3, 4, 5}
s := a[1:3]

fmt.Println(s, len(s), cap(s))

s = s[:4]
fmt.Println(s)
```

`s := a[1:3]` 是：

```text
[2 3]
```

它的容量从 `a[1]` 到末尾，所以容量是 `4`。

因此 `s[:4]` 可以扩展到：

```text
[2 3 4 5]
```

但是不能超过容量：

```go
s = s[:5] // 运行时 panic
```

## 限制切片容量：完整切片表达式

Go 还有三下标切片：

```go
s := a[low:high:max]
```

它的长度是：

```text
high - low
```

容量是：

```text
max - low
```

示例：

```go
a := []int{1, 2, 3, 4, 5}
s := a[1:3:3]

fmt.Println(s)
fmt.Println(len(s), cap(s))
```

输出：

```text
[2 3]
2 2
```

这可以限制 `s` 的容量，减少后续 `append` 意外覆盖底层数组其他部分的风险。

## append 可能影响原数组

```go
a := []int{1, 2, 3, 4}
s := a[:2]

s = append(s, 100)

fmt.Println(s)
fmt.Println(a)
```

可能输出：

```text
[1 2 100]
[1 2 100 4]
```

因为 `s` 的容量还够，`append` 直接写进了原底层数组。

如果不希望影响原数组，可以先复制：

```go
s := append([]int(nil), a[:2]...)
s = append(s, 100)
```

## 预分配容量

如果你大概知道要追加多少元素，可以提前分配容量。

```go
s := make([]int, 0, 100)

for i := 0; i < 100; i++ {
	s = append(s, i)
}
```

这样可以减少扩容和复制次数。

这里长度是 `0`，容量是 `100`：

```go
make([]int, 0, 100)
```

表示现在还没有元素，但预留了 100 个元素的空间。

## 切片和数组的区别

| 对比项 | 数组 | 切片 |
| --- | --- | --- |
| 类型写法 | `[3]int` | `[]int` |
| 长度 | 固定 | 可变 |
| 长度是否属于类型 | 是 | 否 |
| 赋值 | 复制整个数组 | 复制切片头，共享底层数组 |
| 传参 | 复制整个数组 | 复制切片头，共享底层数组 |
| 是否能 append | 不能 | 能 |
| 日常使用频率 | 较少 | 很高 |

## 本节 main.go 逐步讲解

### 第 1 步：创建数组

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
```

创建一个长度为 `6` 的整数数组。

```text
索引:      0  1  2  3   4   5
primes:   2  3  5  7  11  13
```

### 第 2 步：从数组创建切片

```go
var s []int = primes[1:4]
```

`s` 的类型是：

```go
[]int
```

`primes[1:4]` 取的是：

```text
下标 1 到下标 4 之前
```

也就是：

```text
3, 5, 7
```

所以：

```go
fmt.Println(s)
```

输出：

```text
[3 5 7]
```

### 第 3 步：理解 s 和 primes 的关系

`s` 不是一个全新的独立数组。

它引用的是 `primes` 的一部分。

```text
primes: [2 3 5 7 11 13]
            └───s───┘

s:          [3 5 7]
```

如果修改 `s`：

```go
s[0] = 100
```

那么 `primes[1]` 也会变。

## 常见坑

### 1. 以为切片会复制数据

```go
s := a[1:3]
```

这不会复制数据，只是引用底层数组的一段。

### 2. 忘记接住 append 返回值

必须写：

```go
s = append(s, x)
```

### 3. range 修改不到原元素

错误：

```go
for _, v := range s {
	v = v * 10
}
```

正确：

```go
for i := range s {
	s[i] = s[i] * 10
}
```

### 4. nil 切片和空切片在 JSON 中可能不同

```go
var a []int  // JSON: null
b := []int{} // JSON: []
```

### 5. append 可能影响共享的底层数组

如果多个切片共享底层数组，一个切片 `append` 后可能影响另一个切片或原数组。

需要独立数据时，使用 `copy`。

## 总结

切片的核心知识点：

1. 切片类型写作 `[]T`。
2. 切片是对底层数组某一段的引用。
3. `a[low:high]` 包含 `low`，不包含 `high`。
4. `len` 是当前切片长度。
5. `cap` 是从切片起点到底层数组末尾的容量。
6. 切片赋值和传参会复制切片头，但共享底层数组。
7. 修改切片元素可能影响其他切片或原数组。
8. `append` 可能复用原数组，也可能分配新数组。
9. 改变切片长度时，要接住并返回新的切片。
10. 需要独立数据时，使用 `copy`。

一句话记忆：

> 数组是数据本体，切片是对数组一段数据的窗口；窗口可以变长，但多个窗口可能看到同一块数据。
