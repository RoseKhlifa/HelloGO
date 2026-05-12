# 06 · Multiple Results — 函数返回多个值

对应 Tour of Go: <https://go.dev/tour/basics/6>

---

## 这一节学什么

**Go 的函数可以一次返回多个值**。这是 Go 区别于 C / Java 的一个标志性设计。

语法:

```go
func 函数名(参数) (返回值类型1, 返回值类型2, ...) {
    return 值1, 值2, ...
}
```

调用时用多重赋值接收:

```go
a, b := 函数名(...)
```

---

## Tour of Go 的例子

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
```

输出:

```
world hello
```

---

## 为啥这个设计重要

**Go 没有 try / catch 异常机制**,错误是通过"额外的返回值"传递的。所以一旦你开始读 Go 标准库,会铺天盖地看到这种模式:

```go
result, err := doSomething()
if err != nil {
    // 处理错误
    return
}
// 用 result
```

后面几节就会碰到。**多返回值不是花哨语法,是 Go 错误处理哲学的基础**。

---

## 我踩的坑

### 1. 只写了 `package main` 就 `go run .`

```
runtime.main_main·f: function main is undeclared in the main package
```

**Go 程序必须有 `func main()` 作为入口**,光声明包名不够。

### 2. 想忽略某个返回值?用 `_`(下划线)

```go
_, err := doSomething()    // 只关心 err,丢弃第一个返回值
result, _ := doSomething() // 反过来也行
```

如果你接收了一个变量却不用,Go 编译会直接报错(`declared and not used`)—— Go 对"无用变量"零容忍,得用 `_` 显式表态。

### 3. 返回值类型只有一个时,括号可以省略

```go
func add(x, y int) int { return x + y }            // 一个返回值,无括号
func swap(x, y string) (string, string) { ... }    // 多个返回值,必须有括号
```

### 4. `fmt.Print` 不支持 `%s` 占位符(踩过)

```go
fmt.Print("%s %s", a, b)     // ❌ 不会替换,会原样打印出 "%s %s"
```

实际输出类似 `%s %sworldhello`,因为 `Print` 只是把每个参数**当值原样输出**,根本不解析格式串。

`fmt` 三兄弟的分工:

| 函数 | 自动加空格 | 末尾换行 | 支持 `%s %d` |
|------|----------|---------|-------------|
| `Print`   | 仅非字符串参数之间 | ❌ | ❌ |
| `Println` | ✅ 所有参数之间    | ✅ | ❌ |
| `Printf`  | 完全按格式串走     | ❌(自己写 `\n`)| ✅ |

修法:

```go
fmt.Printf("%s %s\n", a, b)    // 想用格式化
fmt.Println(a, b)              // 不需要格式化时更简洁
```

**经验**:看到 `%s` `%d` `%v` 这种东西,**第一反应是 `Printf`,不是 `Print`**。

---

## 练习想法

写一个 `divmod(a, b int) (int, int)`,返回 `a / b` 和 `a % b`,在 `main` 里调用并打印结果。

进阶:再写一个 `safeDivide(a, b int) (int, error)`,当 `b == 0` 时返回非 nil 的 error。这一步会用到下一节才学的 `error` 类型,可以先放着,学完回来补。