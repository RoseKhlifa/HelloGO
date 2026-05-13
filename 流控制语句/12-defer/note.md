# defer

## 示例代码

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

运行结果：

```text
hello
world
```

## defer 是什么

`defer` 的意思是“延迟执行”。

在 Go 语言中，`defer` 后面通常跟着一个函数调用。这个函数调用不会立刻执行，而是会等到当前函数即将结束时再执行。

这个示例中：

```go
defer fmt.Println("world")
```

表示：先记住 `fmt.Println("world")` 这件事，等 `main` 函数快结束时再执行。

所以程序会先执行：

```go
fmt.Println("hello")
```

打印：

```text
hello
```

然后 `main` 函数准备结束，这时再执行前面被 `defer` 延迟的代码，打印：

```text
world
```

## 执行流程

这段代码的执行顺序可以理解为：

1. 进入 `main` 函数。
2. 遇到 `defer fmt.Println("world")`。
3. Go 不会马上打印 `world`，而是把这次函数调用保存起来。
4. 继续向下执行，遇到 `fmt.Println("hello")`。
5. 立刻打印 `hello`。
6. `main` 函数执行完毕，准备退出。
7. 执行被 `defer` 保存的函数调用，打印 `world`。

所以最终输出顺序是：

```text
hello
world
```

## defer 的核心知识点

### 1. defer 会在当前函数结束前执行

`defer` 不是等整个程序结束才执行，而是等“当前函数”结束前执行。

在这个示例中，当前函数是 `main`，所以 `defer` 的内容会在 `main` 结束前执行。

### 2. defer 后面必须是函数调用

常见写法是：

```go
defer 函数名(参数)
```

例如：

```go
defer fmt.Println("world")
```

这里的 `fmt.Println("world")` 是一个函数调用。

### 3. defer 适合处理收尾工作

`defer` 最大的价值是：把“打开”和“关闭”、“申请”和“释放”写在一起，避免忘记清理资源。

比如打开文件后，马上写：

```go
file, err := os.Open("test.txt")
if err != nil {
	return
}
defer file.Close()
```

这样只要函数结束，文件就会自动关闭。

### 4. 多个 defer 会后进先出

如果一个函数里有多个 `defer`，它们的执行顺序是“后进先出”，也就是最后写的 `defer` 会最先执行。

示例：

```go
defer fmt.Println("1")
defer fmt.Println("2")
defer fmt.Println("3")
```

输出顺序是：

```text
3
2
1
```

可以把它理解成一摞盘子：后放上去的，先拿下来。

## defer 的用途

### 1. 关闭文件

读取或写入文件时，经常需要在使用完后关闭文件：

```go
file, err := os.Open("data.txt")
if err != nil {
	return
}
defer file.Close()
```

好处是：不用在函数的每个返回位置都手动写一次 `file.Close()`。

### 2. 释放锁

并发编程中，经常会先加锁，再解锁：

```go
mu.Lock()
defer mu.Unlock()
```

这样可以保证函数结束时一定会释放锁，减少忘记解锁造成的问题。

### 3. 关闭网络连接

网络请求、数据库连接等资源使用完后也需要关闭：

```go
resp, err := http.Get("https://example.com")
if err != nil {
	return
}
defer resp.Body.Close()
```

这样函数结束时，请求响应体会被关闭。

### 4. 统一打印日志或调试信息

有时可以用 `defer` 记录函数退出：

```go
fmt.Println("start")
defer fmt.Println("end")
```

这样无论函数中间怎么返回，退出前都会执行 `end`。

## 使用场景

`defer` 常见于这些场景：

1. 函数结束前必须做清理工作。
2. 代码中有多个 `return`，但每个出口都需要执行同样的收尾逻辑。
3. 打开了文件、网络连接、数据库连接等资源。
4. 加锁后需要确保函数退出时解锁。
5. 想让“资源申请”和“资源释放”的代码靠近，提升可读性。

## 注意点

`defer` 虽然写在前面，但真正执行是在当前函数结束前。

所以这个示例里：

```go
defer fmt.Println("world")
fmt.Println("hello")
```

不是先打印 `world`，而是先打印 `hello`，最后打印 `world`。

学习 `defer` 时，最重要的是记住一句话：

> `defer` 会把函数调用延迟到当前函数结束前执行。
