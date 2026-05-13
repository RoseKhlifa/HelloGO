# Structs 结构体

这一节代码很短：

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

运行结果：

```text
{1 2}
```

虽然示例只有几行，但 `struct` 是 Go 里非常核心的知识点。Go 没有传统面向对象语言里的 `class`，很多时候会用“结构体 + 方法 + 接口”来组织程序。

## struct 是什么

`struct` 是一种自定义复合类型，用来把多个字段组合成一个整体。

比如二维坐标点可以有两个字段：

```go
type Vertex struct {
	X int
	Y int
}
```

这里定义了一个新的类型：`Vertex`。

它里面有两个字段：

```text
X int
Y int
```

可以理解为：

```text
Vertex = {
    X: 一个 int
    Y: 一个 int
}
```

如果你有 C++ 基础，可以先把它类比成：

```cpp
struct Vertex {
    int X;
    int Y;
};
```

但 Go 的结构体和 C++ 的 struct/class 有不少差异，后面会单独讲。

## 基本语法

定义结构体的格式：

```go
type 类型名 struct {
	字段名 字段类型
	字段名 字段类型
}
```

示例：

```go
type User struct {
	Name string
	Age  int
}
```

这表示定义了一个 `User` 类型，它有两个字段：

```text
Name string
Age int
```

## 字段类型相同时可以合并写

下面两种写法等价。

分开写：

```go
type Vertex struct {
	X int
	Y int
}
```

合并写：

```go
type Vertex struct {
	X, Y int
}
```

如果多个字段类型相同，合并写会更简洁。

## 创建结构体值

### 1. 按字段顺序初始化

本节示例使用的是按字段顺序初始化：

```go
Vertex{1, 2}
```

因为 `Vertex` 的字段顺序是：

```go
type Vertex struct {
	X int
	Y int
}
```

所以：

```go
Vertex{1, 2}
```

等价于：

```text
X = 1
Y = 2
```

完整写法：

```go
v := Vertex{1, 2}
fmt.Println(v)
```

输出：

```text
{1 2}
```

注意：按顺序初始化时，值的数量和顺序都必须和结构体字段匹配。

### 2. 按字段名初始化

更推荐的写法是写字段名：

```go
v := Vertex{X: 1, Y: 2}
```

好处是更清楚：

```text
X 是 1
Y 是 2
```

即使字段顺序变化，也不容易写错。

实际项目里，尤其是字段较多时，更推荐使用这种写法。

### 3. 只初始化部分字段

使用字段名初始化时，可以只写一部分字段。

```go
v := Vertex{X: 1}
```

没有写到的字段会使用零值。

因为 `Y` 是 `int`，`int` 的零值是 `0`，所以：

```go
fmt.Println(v)
```

输出：

```text
{1 0}
```

### 4. 使用零值创建

可以只声明，不手动赋值：

```go
var v Vertex
fmt.Println(v)
```

因为 `int` 的零值是 `0`，所以输出：

```text
{0 0}
```

Go 很重视“零值可用”这个设计思想。很多类型即使不手动初始化，也应该处于一个安全、可用或至少可预测的状态。

## 访问结构体字段

使用点号 `.` 访问字段：

```go
v := Vertex{X: 1, Y: 2}

fmt.Println(v.X)
fmt.Println(v.Y)
```

输出：

```text
1
2
```

## 修改结构体字段

结构体变量可以修改字段：

```go
v := Vertex{X: 1, Y: 2}
v.X = 100
v.Y = 200

fmt.Println(v)
```

输出：

```text
{100 200}
```

如果结构体变量本身是不可修改的上下文，字段也不能改。比如函数返回值不能直接修改字段：

```go
func makeVertex() Vertex {
	return Vertex{X: 1, Y: 2}
}

func main() {
	makeVertex().X = 10 // 错误
}
```

原因是 `makeVertex()` 返回的是一个临时值，不能直接对这个临时值的字段赋值。

应该先接住：

```go
v := makeVertex()
v.X = 10
```

## 打印结构体

本节示例：

```go
fmt.Println(Vertex{1, 2})
```

输出：

```text
{1 2}
```

这是默认打印方式，只显示字段值，不显示字段名。

如果想显示字段名，可以用：

```go
fmt.Printf("%+v\n", Vertex{X: 1, Y: 2})
```

输出：

```text
{X:1 Y:2}
```

如果想显示 Go 语法形式，可以用：

```go
fmt.Printf("%#v\n", Vertex{X: 1, Y: 2})
```

输出类似：

```text
main.Vertex{X:1, Y:2}
```

调试结构体时，`%+v` 和 `%#v` 很常用。

## 结构体的零值

结构体的零值是：每个字段都取自己的零值。

```go
type User struct {
	Name    string
	Age     int
	Active  bool
	Address *string
}

var u User
fmt.Printf("%+v\n", u)
```

输出类似：

```text
{Name: Age:0 Active:false Address:<nil>}
```

字段零值规则：

| 类型 | 零值 |
| --- | --- |
| `int` | `0` |
| `float64` | `0` |
| `string` | `""` |
| `bool` | `false` |
| 指针 | `nil` |
| slice | `nil` |
| map | `nil` |
| struct | 每个字段都是零值 |

这点很重要，因为 Go 里经常直接使用零值，而不是强制写构造函数。

## 结构体是值类型

Go 的结构体默认是值类型。

赋值时，会复制整个结构体。

```go
v1 := Vertex{X: 1, Y: 2}
v2 := v1

v2.X = 100

fmt.Println(v1)
fmt.Println(v2)
```

输出：

```text
{1 2}
{100 2}
```

`v2 := v1` 会复制一份新的 `Vertex`。

修改 `v2.X` 不会影响 `v1.X`。

这和 C++ 中普通对象赋值会复制对象的感觉比较像。

## 结构体指针

如果想通过一个变量修改原来的结构体，可以使用指针。

```go
v := Vertex{X: 1, Y: 2}
p := &v

p.X = 100

fmt.Println(v)
```

输出：

```text
{100 2}
```

注意 Go 这里没有 `->`。

在 C++ 里可能写：

```cpp
p->X = 100;
```

Go 里写：

```go
p.X = 100
```

虽然 `p` 是 `*Vertex`，但 Go 允许你直接用 `p.X`。

它等价于：

```go
(*p).X = 100
```

Go 自动帮你做了结构体指针的解引用。

## 结构体指针初始化

可以直接取地址：

```go
p := &Vertex{X: 1, Y: 2}
```

这里 `p` 的类型是：

```go
*Vertex
```

也可以使用 `new`：

```go
p := new(Vertex)
```

这会创建一个零值 `Vertex`，并返回它的指针。

等价理解：

```go
p := &Vertex{}
```

然后可以修改字段：

```go
p.X = 10
p.Y = 20
```

实际项目里，`&Vertex{...}` 比 `new(Vertex)` 更常见，因为它可以在创建时顺手设置字段。

## 函数传结构体：值传递

Go 函数参数默认是值传递。

```go
func move(v Vertex) {
	v.X = v.X + 10
}

func main() {
	v := Vertex{X: 1, Y: 2}
	move(v)
	fmt.Println(v)
}
```

输出：

```text
{1 2}
```

因为传进去的是一份拷贝。

如果希望函数修改原来的结构体，要传指针：

```go
func move(v *Vertex) {
	v.X = v.X + 10
}

func main() {
	v := Vertex{X: 1, Y: 2}
	move(&v)
	fmt.Println(v)
}
```

输出：

```text
{11 2}
```

## 什么时候传结构体指针

常见情况：

1. 函数需要修改原结构体。
2. 结构体比较大，复制成本较高。
3. 需要表达“可能没有值”，也就是 `nil`。
4. 要和方法的指针接收者保持一致。

如果结构体很小，而且不需要修改原值，直接传值也很好。

比如本节的 `Vertex` 只有两个 `int` 字段，传值完全没问题。

不要因为“指针看起来高级”就到处用指针。Go 代码更看重清晰和语义。

## 结构体方法

Go 结构体里面不能像 C++ class 那样直接写成员函数。

Go 的方法定义在结构体外面。

```go
type Vertex struct {
	X, Y int
}

func (v Vertex) Sum() int {
	return v.X + v.Y
}
```

这里：

```go
func (v Vertex) Sum() int
```

表示给 `Vertex` 类型定义一个方法 `Sum`。

调用：

```go
v := Vertex{X: 1, Y: 2}
fmt.Println(v.Sum())
```

输出：

```text
3
```

## 值接收者和指针接收者

方法接收者有两种常见写法。

值接收者：

```go
func (v Vertex) Move(dx, dy int) {
	v.X += dx
	v.Y += dy
}
```

指针接收者：

```go
func (v *Vertex) Move(dx, dy int) {
	v.X += dx
	v.Y += dy
}
```

区别非常重要。

### 值接收者会复制一份

```go
func (v Vertex) Move(dx, dy int) {
	v.X += dx
	v.Y += dy
}

func main() {
	v := Vertex{X: 1, Y: 2}
	v.Move(10, 20)
	fmt.Println(v)
}
```

输出仍然是：

```text
{1 2}
```

因为 `Move` 修改的是副本。

### 指针接收者可以修改原值

```go
func (v *Vertex) Move(dx, dy int) {
	v.X += dx
	v.Y += dy
}

func main() {
	v := Vertex{X: 1, Y: 2}
	v.Move(10, 20)
	fmt.Println(v)
}
```

输出：

```text
{11 22}
```

因为方法拿到的是原结构体地址。

## 方法接收者怎么选

一般经验：

1. 方法需要修改结构体，使用指针接收者。
2. 结构体很大，使用指针接收者避免复制。
3. 结构体包含 `sync.Mutex` 这类不能随便复制的字段，使用指针接收者。
4. 其他方法已经用了指针接收者，通常保持一致。
5. 小结构体只读方法，可以使用值接收者。

比如：

```go
type Point struct {
	X, Y int
}

func (p Point) Distance() int {
	return p.X*p.X + p.Y*p.Y
}
```

这种只读方法用值接收者也很自然。

## 结构体字段的导出规则

Go 用首字母大小写控制包外可见性。

```go
type User struct {
	Name string
	age  int
}
```

这里：

```text
Name 首字母大写，包外可以访问
age 首字母小写，只能在当前包内访问
```

如果另一个包导入了这个类型：

```go
u := other.User{}
u.Name = "Tom" // 可以
u.age = 18     // 不可以
```

这点在写库、写 API、写 JSON 结构体时非常重要。

## JSON 和结构体标签

结构体字段后面可以写标签。

最常见的是 JSON 标签：

```go
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}
```

标签本身是字符串，通常被 `encoding/json`、ORM、配置库等工具读取。

比如：

```go
u := User{ID: 1, Name: "Tom", Age: 18}
data, _ := json.Marshal(u)
fmt.Println(string(data))
```

输出：

```json
{"id":1,"name":"Tom","age":18}
```

常见 JSON 标签：

```go
`json:"name"`
```

表示 JSON 字段名叫 `name`。

```go
`json:"age,omitempty"`
```

表示如果 `age` 是零值，就可以省略。

```go
`json:"-"`
```

表示这个字段不参与 JSON 编码。

注意：JSON 编码通常只能处理导出字段，也就是首字母大写的字段。

下面这个字段即使写了标签，也不会正常导出：

```go
type User struct {
	name string `json:"name"`
}
```

因为 `name` 首字母小写，是未导出字段。

## 匿名结构体

如果只在一个很小的地方临时使用，可以不单独定义类型。

```go
user := struct {
	Name string
	Age  int
}{
	Name: "Tom",
	Age:  18,
}

fmt.Println(user)
```

这叫匿名结构体。

适合：

1. 测试代码。
2. 临时组织数据。
3. 小范围使用，不值得单独定义类型。

如果一个结构体会被多个地方使用，就应该用 `type` 起名字。

## 嵌套结构体

结构体字段可以是另一个结构体。

```go
type Address struct {
	City string
}

type User struct {
	Name    string
	Address Address
}
```

使用：

```go
u := User{
	Name: "Tom",
	Address: Address{
		City: "Shanghai",
	},
}

fmt.Println(u.Address.City)
```

这种叫组合。Go 更鼓励组合，而不是继承。

## 匿名字段和嵌入

Go 结构体可以嵌入另一个类型。

```go
type Address struct {
	City string
}

type User struct {
	Name string
	Address
}
```

这里 `Address` 没有写字段名，所以它是匿名字段，也常被称为嵌入字段。

可以这样访问：

```go
u := User{
	Name: "Tom",
	Address: Address{
		City: "Shanghai",
	},
}

fmt.Println(u.Address.City)
fmt.Println(u.City)
```

`u.City` 可以直接访问，是因为 `Address` 的字段被提升了。

但要注意：这不是继承。

更准确地说，Go 的嵌入是一种组合语法糖。

## 结构体比较

如果结构体的所有字段都可以比较，那么结构体也可以比较。

```go
type Point struct {
	X, Y int
}

p1 := Point{X: 1, Y: 2}
p2 := Point{X: 1, Y: 2}

fmt.Println(p1 == p2)
```

输出：

```text
true
```

但是如果结构体中包含不能比较的字段，比如 slice、map、function，那么结构体不能直接用 `==` 比较。

```go
type User struct {
	Name string
	Tags []string
}
```

这种结构体不能直接比较：

```go
u1 == u2 // 错误
```

如果要比较复杂结构体，可以使用：

```go
reflect.DeepEqual(a, b)
```

或者自己写比较逻辑。

实际项目中，自己写比较逻辑通常更清楚，因为你可以决定哪些字段参与比较。

## 空结构体 struct{}

Go 里有一个特殊结构体：

```go
struct{}
```

它没有任何字段，通常不占用额外存储空间。

常见用途：

### 1. 只关心集合中的 key

```go
seen := map[string]struct{}{}
seen["go"] = struct{}{}
```

这里使用 `map[string]struct{}` 表示一个集合。

如果只关心 `"go"` 是否出现过，不关心对应的值，用 `struct{}` 比 `bool` 更强调“值不重要”。

### 2. channel 信号

```go
done := make(chan struct{})

go func() {
	// do something
	close(done)
}()

<-done
```

`struct{}` 可以作为“没有数据，只传递信号”的类型。

## 结构体和内存布局

结构体字段在内存中通常按定义顺序排列，但编译器会为了对齐插入填充字节。

比如：

```go
type A struct {
	B bool
	I int64
}
```

和：

```go
type B struct {
	I int64
	B bool
}
```

字段一样，但内存占用可能不同。

这是因为 CPU 访问某些类型时需要对齐，编译器会插入 padding。

初学阶段不用太纠结这个，但要知道：

1. 字段顺序可能影响结构体大小。
2. 对性能非常敏感时，可以关注字段排列。
3. 普通业务代码优先考虑可读性。

## 结构体和 C++ 的重要差异

### 1. Go 没有 class

Go 没有 `class` 关键字。

通常用：

```text
struct + method + interface
```

来组织代码。

### 2. Go 的方法不写在结构体里面

C++：

```cpp
struct User {
    string name;
    void hello() {}
};
```

Go：

```go
type User struct {
	Name string
}

func (u User) Hello() {
	fmt.Println(u.Name)
}
```

### 3. Go 没有构造函数语法

Go 没有专门的 constructor。

常见写法是普通函数：

```go
func NewUser(name string, age int) User {
	return User{Name: name, Age: age}
}
```

如果结构体较大，或者希望避免复制，也可以返回指针：

```go
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}
```

### 4. Go 没有析构函数

Go 没有 destructor。

资源释放通常靠显式方法配合 `defer`：

```go
file, err := os.Open("data.txt")
if err != nil {
	return err
}
defer file.Close()
```

### 5. Go 不靠继承复用代码

Go 没有类继承。

Go 更鼓励组合：

```go
type Logger struct{}

func (Logger) Print(msg string) {
	fmt.Println(msg)
}

type Service struct {
	Logger
}
```

`Service` 嵌入了 `Logger`，可以直接调用：

```go
s := Service{}
s.Print("hello")
```

但这不是继承，只是组合和方法提升。

## 结构体设计建议

### 1. 字段名尽量清楚

不推荐：

```go
type User struct {
	N string
	A int
}
```

更推荐：

```go
type User struct {
	Name string
	Age  int
}
```

除非是非常短、非常局部的结构体，否则字段名应该表达含义。

### 2. 对外暴露的结构体要慎重设计字段

如果字段首字母大写，外部包可以直接访问和修改。

```go
type User struct {
	Name string
}
```

外部可以：

```go
u.Name = "Jerry"
```

如果不想外部直接改，可以使用小写字段，再提供方法：

```go
type User struct {
	name string
}

func (u User) Name() string {
	return u.name
}
```

### 3. 初始化字段多时，优先使用字段名

不推荐：

```go
u := User{"Tom", 18, true, "Shanghai"}
```

更推荐：

```go
u := User{
	Name:   "Tom",
	Age:    18,
	Active: true,
	City:   "Shanghai",
}
```

字段名写法更长，但更安全、更可读。

### 4. 注意值复制

结构体赋值和传参会复制。

如果结构体里有指针、slice、map，复制结构体时，这些字段本身会被复制，但它们指向的底层数据可能还是共享的。

例子：

```go
type User struct {
	Tags []string
}

u1 := User{Tags: []string{"go"}}
u2 := u1

u2.Tags[0] = "cpp"

fmt.Println(u1.Tags)
fmt.Println(u2.Tags)
```

输出：

```text
[cpp]
[cpp]
```

原因是 slice 字段被复制了，但两个 slice 仍然指向同一个底层数组。

这点很容易踩坑。

### 5. 不要把所有东西都做成指针字段

比如：

```go
type User struct {
	Name *string
	Age  *int
}
```

如果没有“可选值”的需求，这样会让代码更麻烦：

1. 使用前经常要判断 `nil`。
2. 访问时要解引用。
3. 更容易出现空指针问题。

普通字段够用时，就用普通字段：

```go
type User struct {
	Name string
	Age  int
}
```

只有确实需要表达“没有值”和“零值”区别时，再考虑指针字段。

## 本节 main.go 的知识点

回到本节示例：

```go
type Vertex struct {
	X int
	Y int
}
```

这定义了一个结构体类型 `Vertex`。

```go
fmt.Println(Vertex{1, 2})
```

这里做了两件事：

1. 使用结构体字面量 `Vertex{1, 2}` 创建一个 `Vertex` 值。
2. 使用 `fmt.Println` 打印这个结构体。

因为字段顺序是 `X`、`Y`，所以：

```go
Vertex{1, 2}
```

表示：

```go
Vertex{
	X: 1,
	Y: 2,
}
```

打印结果：

```text
{1 2}
```

如果想看得更清楚，可以改成：

```go
fmt.Printf("%+v\n", Vertex{X: 1, Y: 2})
```

输出：

```text
{X:1 Y:2}
```

## 总结

结构体是 Go 里组织数据的核心工具。

你需要重点掌握：

1. `type Name struct { ... }` 定义结构体类型。
2. `Name{...}` 创建结构体值。
3. `Name{Field: value}` 使用字段名初始化，更推荐。
4. `v.X` 访问字段。
5. `v.X = 10` 修改字段。
6. 结构体默认是值类型，赋值和传参会复制。
7. 需要修改原值时使用结构体指针。
8. Go 没有 `->`，结构体指针也用 `.` 访问字段。
9. 方法定义在结构体外面，通过接收者绑定到类型。
10. 字段首字母大小写决定是否能被包外访问。
11. 结构体标签常用于 JSON、数据库、配置解析等场景。
12. Go 更鼓励组合，而不是继承。

一句话记忆：

> `struct` 用来把一组相关字段组织成一个类型；Go 通过结构体保存数据，通过方法描述行为，通过接口表达能力。
