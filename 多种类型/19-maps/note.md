# Maps 映射

这一节代码：

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

运行结果：

```text
{40.68433 -74.39967}
```

## map 是什么

`map` 是 Go 里的键值对集合。

可以理解为：

```text
key -> value
```

比如本节示例：

```go
map[string]Vertex
```

意思是：

```text
key 的类型是 string
value 的类型是 Vertex
```

也就是：

```text
用地点名字找到对应坐标
```

例如：

```text
"Bell Labs" -> Vertex{40.68433, -74.39967}
```

如果你有 C++ 基础，可以把 Go 的 `map` 先粗略类比成：

```cpp
std::unordered_map<Key, Value>
```

它通常基于哈希表思想，用 key 快速找到 value。

## map 的类型写法

基本格式：

```go
map[KeyType]ValueType
```

示例：

```go
map[string]int
```

表示：

```text
key 是 string
value 是 int
```

再比如：

```go
map[int]string
```

表示：

```text
key 是 int
value 是 string
```

本节示例：

```go
map[string]Vertex
```

表示：

```text
key 是 string
value 是 Vertex 结构体
```

## 本节里的 Vertex

```go
type Vertex struct {
	Lat, Long float64
}
```

这里定义了一个结构体 `Vertex`。

它有两个字段：

```text
Lat  float64
Long float64
```

`Lat` 通常表示纬度。

`Long` 通常表示经度。

所以一个 `Vertex` 可以表示一个地理坐标。

## 声明 map

本节代码中：

```go
var m map[string]Vertex
```

这行只是声明了一个 map 变量。

此时 `m` 的零值是：

```go
nil
```

也就是说，它还没有真正创建底层哈希表。

可以读取 nil map，但不能向 nil map 写入。

```go
var m map[string]int

fmt.Println(m["a"]) // 可以，读到 int 的零值 0
m["a"] = 1          // panic：assignment to entry in nil map
```

所以写入前必须先初始化。

## 使用 make 创建 map

本节代码中：

```go
m = make(map[string]Vertex)
```

`make` 会创建一个可以使用的 map。

创建后，就可以添加、修改、查询、删除元素了。

常见写法：

```go
m := make(map[string]int)
```

也可以预估容量：

```go
m := make(map[string]int, 100)
```

第二个参数是初始容量提示，表示你大概会放多少元素。

注意：这只是给运行时的提示，不是固定长度限制。map 仍然可以继续增长。

## 添加或修改元素

本节代码：

```go
m["Bell Labs"] = Vertex{
	40.68433, -74.39967,
}
```

意思是：

```text
把 key "Bell Labs" 对应的 value 设置为 Vertex{40.68433, -74.39967}
```

如果 key 不存在，就是添加。

如果 key 已经存在，就是修改。

示例：

```go
ages := make(map[string]int)

ages["Tom"] = 18 // 添加
ages["Tom"] = 20 // 修改

fmt.Println(ages["Tom"])
```

输出：

```text
20
```

## 读取元素

本节代码：

```go
fmt.Println(m["Bell Labs"])
```

表示读取 key `"Bell Labs"` 对应的 value。

输出：

```text
{40.68433 -74.39967}
```

如果读取一个不存在的 key，会得到 value 类型的零值。

```go
ages := map[string]int{
	"Tom": 18,
}

fmt.Println(ages["Jerry"])
```

输出：

```text
0
```

因为 `int` 的零值是 `0`。

这会带来一个问题：无法只通过返回值判断 key 是不存在，还是 key 存在但 value 本来就是零值。

## 判断 key 是否存在：comma ok

Go 中读取 map 时可以使用两个返回值：

```go
value, ok := m[key]
```

示例：

```go
ages := map[string]int{
	"Tom":   18,
	"Jerry": 0,
}

age, ok := ages["Jerry"]
fmt.Println(age, ok)

age, ok = ages["Alice"]
fmt.Println(age, ok)
```

输出：

```text
0 true
0 false
```

解释：

```text
0 true  -> key 存在，value 是 0
0 false -> key 不存在，返回的是 int 零值
```

这是 Go map 必须掌握的知识点。

常见写法：

```go
if v, ok := m[key]; ok {
	fmt.Println("存在：", v)
} else {
	fmt.Println("不存在")
}
```

## 删除元素

使用内置函数 `delete`：

```go
delete(m, key)
```

示例：

```go
ages := map[string]int{
	"Tom": 18,
}

delete(ages, "Tom")

fmt.Println(ages["Tom"])
```

输出：

```text
0
```

删除不存在的 key 也不会报错：

```go
delete(ages, "NotExist")
```

这是安全的。

## map 字面量

除了 `make`，还可以用字面量创建 map。

```go
ages := map[string]int{
	"Tom":   18,
	"Jerry": 20,
}
```

本节示例也可以写成：

```go
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
}
```

如果 value 类型已经明确，可以省略里面的类型名：

```go
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
}
```

这一点在 map value 是结构体时很常见。

## nil map 和空 map

nil map：

```go
var a map[string]int
```

空 map：

```go
b := map[string]int{}
c := make(map[string]int)
```

区别：

```go
fmt.Println(a == nil) // true
fmt.Println(b == nil) // false
fmt.Println(c == nil) // false
```

读取：

```go
fmt.Println(a["x"]) // 可以，返回 0
```

写入：

```go
a["x"] = 1 // panic
b["x"] = 1 // 可以
c["x"] = 1 // 可以
```

所以：

```text
nil map 可以读，不能写。
空 map 可以读，也可以写。
```

## map 的 len

`len(m)` 返回 map 中 key-value 对的数量。

```go
m := map[string]int{
	"a": 1,
	"b": 2,
}

fmt.Println(len(m))
```

输出：

```text
2
```

添加元素后，长度会增加。

删除元素后，长度会减少。

## 遍历 map

使用 `range` 遍历：

```go
ages := map[string]int{
	"Tom":   18,
	"Jerry": 20,
}

for name, age := range ages {
	fmt.Println(name, age)
}
```

注意：map 的遍历顺序是不固定的。

你不能依赖遍历顺序。

同一段程序多次运行，输出顺序也可能不一样。

如果需要稳定顺序，要先取出 key，然后排序。

```go
keys := make([]string, 0, len(ages))
for key := range ages {
	keys = append(keys, key)
}

sort.Strings(keys)

for _, key := range keys {
	fmt.Println(key, ages[key])
}
```

## map 的 key 类型要求

map 的 key 必须是可以比较的类型。

可以作为 key 的常见类型：

1. `string`
2. `int`
3. `bool`
4. 指针
5. 数组
6. 结构体，前提是结构体所有字段都可以比较

不能作为 key 的常见类型：

1. slice
2. map
3. function

因为这些类型不能用 `==` 比较。

示例：

```go
map[[]int]string // 错误，slice 不能作为 key
```

如果你需要用一组数字作为 key，可以考虑使用数组：

```go
m := map[[2]int]string{}
m[[2]int{1, 2}] = "point"
```

## map 的 value 类型

map 的 value 可以是很多类型。

比如：

```go
map[string]int
map[string]string
map[string]bool
map[string]Vertex
map[string][]int
map[string]map[string]int
```

本节就是：

```go
map[string]Vertex
```

表示 key 是字符串，value 是结构体。

## 修改 map 中结构体字段的坑

如果 map 的 value 是结构体，不能直接修改结构体字段。

例如：

```go
type User struct {
	Age int
}

users := map[string]User{
	"Tom": {Age: 18},
}

users["Tom"].Age = 20 // 错误
```

为什么？

因为 `users["Tom"]` 取出来的是一个临时副本，不能直接修改它的字段。

正确写法一：取出来，改完，再放回去。

```go
u := users["Tom"]
u.Age = 20
users["Tom"] = u
```

正确写法二：value 使用结构体指针。

```go
users := map[string]*User{
	"Tom": {Age: 18},
}

users["Tom"].Age = 20
```

这时 `users["Tom"]` 是 `*User`，可以通过指针修改原结构体。

是否使用指针 value 要看语义：

1. value 小且希望 map 保存独立值，用结构体值。
2. value 大或需要频繁修改字段，用结构体指针。

## map 是引用类型吗

日常可以把 map 理解为“引用底层数据结构的描述值”。

把 map 赋值给另一个变量，不会复制所有 key-value。

```go
a := map[string]int{"x": 1}
b := a

b["x"] = 100

fmt.Println(a["x"])
```

输出：

```text
100
```

因为 `a` 和 `b` 指向同一个底层 map 数据。

函数传参也是类似：

```go
func change(m map[string]int) {
	m["x"] = 100
}

func main() {
	a := map[string]int{"x": 1}
	change(a)
	fmt.Println(a["x"])
}
```

输出：

```text
100
```

但注意：map 变量本身仍然是按值传递的。

如果函数里把参数重新赋值为一个新 map，不会改变外面的 map 变量。

```go
func reset(m map[string]int) {
	m = make(map[string]int)
	m["x"] = 100
}

func main() {
	a := map[string]int{"x": 1}
	reset(a)
	fmt.Println(a["x"])
}
```

输出：

```text
1
```

函数里改的是参数变量 `m` 指向的新 map，不是外面的变量 `a`。

## map 和并发

Go 的普通 map 不是并发安全的。

如果多个 goroutine 同时读写同一个 map，可能会 panic 或产生数据竞争。

错误场景：

```go
go func() {
	m["a"] = 1
}()

go func() {
	fmt.Println(m["a"])
}()
```

如果需要并发访问 map，常见选择：

1. 用 `sync.Mutex` 或 `sync.RWMutex` 保护普通 map。
2. 使用 `sync.Map`，适合特定高并发读写场景。
3. 用 channel 把 map 限制在单个 goroutine 内管理。

初学阶段先记住：

> 普通 map 不要无保护地并发读写。

## map 和 JSON

map 很常用于 JSON 或动态字段。

```go
m := map[string]any{
	"name": "Tom",
	"age":  18,
}
```

编码成 JSON：

```go
data, _ := json.Marshal(m)
fmt.Println(string(data))
```

输出类似：

```json
{"age":18,"name":"Tom"}
```

注意：JSON 对象的字段顺序不应该作为业务逻辑依赖。

如果数据结构固定，通常更推荐使用 struct。

如果字段不固定，map 更灵活。

## map 适合哪些场景

### 1. 根据 key 快速查找 value

```go
users := map[int]string{
	1: "Tom",
	2: "Jerry",
}

fmt.Println(users[1])
```

### 2. 统计次数

```go
counts := map[string]int{}

for _, word := range words {
	counts[word]++
}
```

如果 key 不存在，`counts[word]` 会返回 `0`，所以可以直接 `++`。

### 3. 去重集合

Go 没有内置 set，常用 map 实现。

```go
set := map[string]bool{}
set["go"] = true
```

更常见的集合写法：

```go
set := map[string]struct{}{}
set["go"] = struct{}{}
```

`struct{}` 表示不关心 value，只关心 key 是否存在。

判断：

```go
_, ok := set["go"]
```

### 4. 缓存

```go
cache := map[string]Result{}
```

用输入参数或 ID 作为 key，把计算结果作为 value。

### 5. 分组

```go
groups := map[string][]User{}

for _, user := range users {
	groups[user.City] = append(groups[user.City], user)
}
```

key 是城市，value 是这个城市的用户列表。

## map 和 C++ 的对比

可以大致这样类比：

```go
map[string]int
```

类似：

```cpp
std::unordered_map<std::string, int>
```

相似点：

1. 都是 key-value 结构。
2. 都可以根据 key 快速查找 value。
3. key 通常需要可比较或可哈希。

重要区别：

1. Go map 是内建类型，语法更直接。
2. Go map 的遍历顺序不固定。
3. Go map 的 nil 值可以读但不能写。
4. Go map 普通情况下不是并发安全的。
5. Go map 不能直接取元素地址。

## 为什么不能取 map 元素地址

下面这种写法不允许：

```go
p := &m["key"] // 错误
```

原因是 map 扩容或重排时，元素位置可能变化。

如果允许你拿到元素地址，后续 map 内部移动元素时，这个地址就可能不安全。

所以 Go 禁止直接取 map 元素地址。

这也解释了为什么 map value 是结构体时不能直接修改字段。

## 本节 main.go 逐步讲解

### 第 1 步：定义坐标结构体

```go
type Vertex struct {
	Lat, Long float64
}
```

定义一个 `Vertex` 类型，表示经纬度。

### 第 2 步：声明 map 变量

```go
var m map[string]Vertex
```

声明一个 map。

key 是 `string`。

value 是 `Vertex`。

此时 `m` 是 nil map，还不能写入。

### 第 3 步：初始化 map

```go
m = make(map[string]Vertex)
```

创建真正可用的 map。

### 第 4 步：写入一条数据

```go
m["Bell Labs"] = Vertex{
	40.68433, -74.39967,
}
```

把 `"Bell Labs"` 这个地点映射到一个经纬度结构体。

可以读作：

```text
m 中 key 为 "Bell Labs" 的 value 是 Vertex{40.68433, -74.39967}
```

### 第 5 步：读取并打印

```go
fmt.Println(m["Bell Labs"])
```

根据 key `"Bell Labs"` 找到对应的 `Vertex` 并打印。

输出：

```text
{40.68433 -74.39967}
```

## 常见坑

### 1. nil map 不能写

```go
var m map[string]int
m["a"] = 1 // panic
```

写之前要：

```go
m = make(map[string]int)
```

### 2. 读取不存在的 key 会返回零值

```go
age := ages["unknown"]
```

如果 `age` 是 `0`，不一定代表 key 存在且值为 `0`。

要用：

```go
age, ok := ages["unknown"]
```

### 3. map 遍历顺序不固定

不要依赖：

```go
for k, v := range m
```

的输出顺序。

需要顺序时，先排序 key。

### 4. map value 是结构体时不能直接改字段

```go
users["Tom"].Age = 20 // 错误
```

要么取出改完放回，要么使用指针 value。

### 5. 普通 map 不能无保护并发读写

并发访问时要加锁、用 `sync.Map`，或者用 channel 管理。

## 总结

map 的核心知识点：

1. `map[K]V` 表示 key 类型是 `K`，value 类型是 `V`。
2. map 需要用 `make` 或字面量初始化后才能写入。
3. nil map 可以读，不能写。
4. `m[key] = value` 用于添加或修改。
5. `v := m[key]` 用于读取。
6. `v, ok := m[key]` 用于判断 key 是否存在。
7. `delete(m, key)` 用于删除。
8. `len(m)` 获取 key-value 数量。
9. `range` 可以遍历 map，但顺序不固定。
10. key 必须是可比较类型。
11. map 赋值和传参不会复制所有数据，多个变量可能共享同一个底层 map。
12. map 不是并发安全的。

一句话记忆：

> map 是根据 key 快速找到 value 的容器；写入前要初始化，读取要注意零值和 ok，遍历不要相信顺序。
