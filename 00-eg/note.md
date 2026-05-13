# fmt 常用输出和字符串格式化函数

Go 的 `fmt` 包常用来做两类事情：

1. 把内容输出到控制台。
2. 把内容拼接或格式化成字符串。

这几个函数可以分成两组：

| 函数 | 作用 | 是否直接输出 | 是否返回字符串 |
| --- | --- | --- | --- |
| `fmt.Print()` | 直接输出内容，不自动换行 | 是 | 否 |
| `fmt.Println()` | 直接输出内容，自动换行 | 是 | 否 |
| `fmt.Printf()` | 按格式直接输出内容 | 是 | 否 |
| `fmt.Sprint()` | 拼接成字符串，不自动换行 | 否 | 是 |
| `fmt.Sprintln()` | 拼接成字符串，末尾自动加换行 | 否 | 是 |
| `fmt.Sprintf()` | 按格式生成字符串 | 否 | 是 |

## fmt.Print()

`fmt.Print()` 用来把内容直接输出到控制台。

它不会自动在末尾添加换行。

```go
fmt.Print("hello")
fmt.Print("world")
```

输出结果：

```text
helloworld
```

如果想自己换行，需要手动加 `\n`：

```go
fmt.Print("hello\n")
fmt.Print("world\n")
```

适合场景：

1. 想连续输出内容，不想自动换行。
2. 想自己控制换行位置。
3. 做简单的控制台提示，比如输入提示。

示例：

```go
fmt.Print("请输入用户名: ")
```

## fmt.Println()

`fmt.Println()` 用来把内容直接输出到控制台，并且自动在末尾添加换行。

```go
fmt.Println("hello")
fmt.Println("world")
```

输出结果：

```text
hello
world
```

如果传入多个参数，`Println` 会自动在参数之间添加空格。

```go
name := "Go"
age := 15

fmt.Println(name, age)
```

输出结果：

```text
Go 15
```

适合场景：

1. 最常用的控制台输出。
2. 打印一行调试信息。
3. 输出多个变量，并希望它们之间自动用空格分隔。

## fmt.Printf()

`fmt.Printf()` 用来按照指定格式输出内容。

它不会自动换行，想换行需要在格式字符串里写 `\n`。

```go
name := "Go"
age := 15

fmt.Printf("name=%s, age=%d\n", name, age)
```

输出结果：

```text
name=Go, age=15
```

常见格式占位符：

| 占位符 | 含义 | 示例 |
| --- | --- | --- |
| `%s` | 字符串 | `"hello"` |
| `%d` | 十进制整数 | `100` |
| `%f` | 浮点数 | `3.140000` |
| `%.2f` | 保留 2 位小数 | `3.14` |
| `%t` | 布尔值 | `true` |
| `%v` | 按默认格式输出任意值 | 通用 |
| `%T` | 输出值的类型 | `int`、`string` |
| `%#v` | 输出 Go 语法形式的值 | 调试常用 |

适合场景：

1. 输出内容需要固定格式。
2. 需要控制小数位数。
3. 需要把变量嵌入一段文本中。
4. 调试时想同时看变量值和类型。

示例：

```go
price := 12.5
fmt.Printf("价格：%.2f 元\n", price)
```

输出结果：

```text
价格：12.50 元
```

## fmt.Sprint()

`fmt.Sprint()` 不会把内容输出到控制台，而是把内容拼接成一个字符串并返回。

```go
s := fmt.Sprint("hello", "world")
fmt.Println(s)
```

输出结果：

```text
helloworld
```

注意：`Sprint` 不会像 `Println` 那样自动在参数之间添加空格。

如果想要空格，需要自己写：

```go
s := fmt.Sprint("hello", " ", "world")
```

适合场景：

1. 想把多个值拼接成字符串。
2. 不想立刻输出，只想得到一个字符串。
3. 函数需要返回字符串结果。

示例：

```go
name := "Go"
msg := fmt.Sprint("hello ", name)
fmt.Println(msg)
```

## fmt.Sprintln()

`fmt.Sprintln()` 不会直接输出，而是把内容拼接成字符串并返回。

它和 `fmt.Println()` 类似，会在参数之间添加空格，并且在字符串末尾添加换行。

```go
s := fmt.Sprintln("hello", "world")
fmt.Print(s)
```

输出结果：

```text
hello world
```

这里用 `fmt.Print(s)` 输出，是因为 `s` 本身已经带了换行。

适合场景：

1. 想生成一整行字符串。
2. 希望参数之间自动加空格。
3. 希望生成的字符串末尾自带换行。
4. 拼接日志文本、多行文本时使用。

## fmt.Sprintf()

`fmt.Sprintf()` 不会直接输出，而是按照指定格式生成字符串并返回。

它可以理解为：`Printf` 是格式化后直接打印，`Sprintf` 是格式化后返回字符串。

```go
name := "Go"
age := 15

s := fmt.Sprintf("name=%s, age=%d", name, age)
fmt.Println(s)
```

输出结果：

```text
name=Go, age=15
```

适合场景：

1. 想按固定格式生成字符串。
2. 需要把格式化结果保存到变量。
3. 需要把格式化后的字符串作为函数返回值。
4. 生成错误信息、日志内容、文件名、接口参数等。

示例：

```go
id := 7
filename := fmt.Sprintf("user_%03d.txt", id)
fmt.Println(filename)
```

输出结果：

```text
user_007.txt
```

## Print 系列和 Sprint 系列的区别

### Print 系列：直接输出

`Print`、`Println`、`Printf` 都会直接把内容输出到控制台。

```go
fmt.Print("hello")
fmt.Println("hello")
fmt.Printf("hello %s\n", "Go")
```

适合：程序运行时直接展示内容。

### Sprint 系列：返回字符串

`Sprint`、`Sprintln`、`Sprintf` 不会直接输出，而是返回一个字符串。

```go
s1 := fmt.Sprint("hello")
s2 := fmt.Sprintln("hello")
s3 := fmt.Sprintf("hello %s", "Go")
```

适合：先生成字符串，再保存、返回、拼接或传给其他函数。

## 三组对应关系

可以这样记：

| 直接输出 | 生成字符串 | 特点 |
| --- | --- | --- |
| `fmt.Print()` | `fmt.Sprint()` | 普通拼接，不自动换行 |
| `fmt.Println()` | `fmt.Sprintln()` | 参数之间加空格，末尾加换行 |
| `fmt.Printf()` | `fmt.Sprintf()` | 使用格式占位符 |

## 怎么选择

如果只是简单打印一行内容：

```go
fmt.Println("hello")
```

如果想不换行地打印：

```go
fmt.Print("loading...")
```

如果想控制输出格式：

```go
fmt.Printf("score=%.1f\n", 98.5)
```

如果想得到一个拼接后的字符串：

```go
s := fmt.Sprint("hello", " ", "Go")
```

如果想得到一行带换行的字符串：

```go
s := fmt.Sprintln("hello", "Go")
```

如果想得到一个格式化后的字符串：

```go
s := fmt.Sprintf("score=%.1f", 98.5)
```

## 总结

最常用的是：

1. `fmt.Println()`：简单打印一行。
2. `fmt.Printf()`：按格式打印。
3. `fmt.Sprintf()`：按格式生成字符串。

一句话记忆：

> `Print` 系列负责“打印出去”，`Sprint` 系列负责“生成字符串”；带 `ln` 的会处理换行，带 `f` 的支持格式化。
