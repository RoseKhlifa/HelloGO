# 11 · Basic Types — Go 的基础类型

对应 Tour of Go: <https://go.dev/tour/basics/11>

---

## 全部基础类型一览

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

---

## 分类速记

| 类别 | 类型 | 备注 |
|------|------|------|
| 布尔 | `bool` | 只有 `true` / `false` |
| 字符串 | `string` | 不可变,底层是 UTF-8 字节序列 |
| 有符号整数 | `int8 / int16 / int32 / int64` | 数字代表位数 |
| 无符号整数 | `uint8 / uint16 / uint32 / uint64` | 不能存负数,范围翻倍 |
| 平台相关整数 | `int / uint` | 32 位机器 32 位,64 位机器 64 位 |
| 指针整数 | `uintptr` | 给 `unsafe` 包用,日常忘掉 |
| 别名 | `byte` = `uint8`,`rune` = `int32` | 不是新类型,只是换名字表意 |
| 浮点 | `float32 / float64` | 默认用 `float64` |
| 复数 | `complex64 / complex128` | 几乎用不到,知道存在就行 |

---

## 必须知道的注意事项

### 1. `int` 跟 `uint` 的位数取决于平台

- 64 位系统:`int` = 64 位(范围 -2⁶³ ~ 2⁶³-1)
- 32 位系统:`int` = 32 位

**业务代码无脑用 `int` 就行**;要精确控制(网络协议、二进制格式、嵌入式)再用带位数的 `int32` / `int64`。

### 2. `byte` 和 `rune` 是别名,不是新类型

```go
var b byte = 'A'        // 等价于 var b uint8 = 65
var r rune = '中'        // 等价于 var r int32 = 20013(Unicode 码点)
```

意图区分:
- `byte` 用来表达**一个原始字节**(读文件、网络流、二进制)
- `rune` 用来表达**一个 Unicode 字符**(处理文本)

### 3. ⚠️ string 索引拿的是字节,不是字符(中文必踩)

```go
s := "中"
fmt.Println(len(s))             // 3  ← "中"在 UTF-8 里占 3 字节
fmt.Println(s[0])               // 228 ← 一个字节,根本不是"中"
fmt.Println(len([]rune(s)))     // 1  ← 1 个 Unicode 字符
```

按字符遍历的正确姿势:

```go
for i, r := range s {
    fmt.Printf("byte index %d: %c\n", i, r)
}
```

`range string` 会按 **rune** 解码,`i` 是字节下标,`r` 是当前 rune。

**规则**:`string` 在 Go 底层就是只读的 `[]byte`(UTF-8 编码)。要按字符处理,**先转 `[]rune`**。

### 4. 没有隐式类型转换,一切都要显式

```go
var x int = 10
var y float64 = 3.14
// z := x + y                  // ❌ 编译错误:mismatched types
z := float64(x) + y             // ✅

var a int32 = 10
var b int64 = 20
// c := a + b                  // ❌ int32 和 int64 互不兼容
c := int64(a) + b               // ✅
```

**Go 的哲学:宁可啰嗦,不能模糊**。跟 JS / Python 的"自动转换"相反。带来的好处是**不会有隐式精度丢失或溢出**。

### 5. `string` 不可变

```go
s := "hello"
// s[0] = 'H'                  // ❌ 编译错误
s = "Hello"                     // ✅ 整个重新赋值可以
```

想"原地改",转成 `[]byte` 或 `[]rune` 改完再转回去:

```go
b := []byte(s)
b[0] = 'H'
s = string(b)
```

### 6. `uintptr` 和复数类型几乎用不到

- `uintptr`:存指针的整数,给 `unsafe` 包用,**初学完全不用碰**。
- `complex64` / `complex128`:科学计算用,**业务代码一辈子不会用**。

---

## Tour of Go 这一节的代码

```go
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
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

这段代码顺手介绍了几个新东西:

| 写法 | 含义 |
|------|------|
| `var ( ... )` | **变量分组声明**,跟 `import ( ... )` 一样的语法 |
| `%T` | Printf 动词,**打印类型** |
| `%v` | Printf 动词,**打印值**(万能默认) |
| `1<<64 - 1` | 位运算,1 左移 64 位再减 1 = `uint64` 最大值(64 个 1) |
| `cmplx.Sqrt` | 复数开方,顺带演示复数类型 |

---

## Printf 动词速查(常用)

| 动词 | 用途 |
|------|------|
| `%v` | 默认格式,**不知道用啥就 %v** |
| `%T` | 类型(调试神器) |
| `%d` | 十进制整数 |
| `%s` | 字符串 |
| `%f` | 浮点数(默认 6 位小数) |
| `%t` | 布尔 |
| `%c` | Unicode 字符(给 rune) |
| `%q` | 带引号的字符串/字符 |
| `%x` / `%X` | 十六进制(小写 / 大写) |
| `%b` | 二进制 |
| `%p` | 指针地址 |

---

## 核心结论(背下来)

1. **`int` 看平台,要精确就用 `int32` / `int64`**
2. **`byte` = `uint8`,`rune` = `int32`(Unicode 码点)**
3. **`string` 是 UTF-8 字节序列,`s[i]` 取的是字节;处理字符用 `[]rune` 或 `range`**
4. **Go 没有隐式类型转换,该 `T(x)` 就 `T(x)`**
5. **`%v` 和 `%T` 是调试神器**

---

## 练习想法

写一个 `main`,声明几个不同类型的变量(比如一个 `int`、一个 `float64`、一个 `string`、一个 `rune`),用 `fmt.Printf("%T = %v\n", x, x)` 打印它们的类型和值,直观感受一下。

进阶:写一个函数 `countRunes(s string) int` 返回字符串里的字符(rune)个数,**不能用 `len(s)`**(那是字节数)。试试输入 `"hello"`、`"中文"`、`"a中b"`,看看结果对不对。