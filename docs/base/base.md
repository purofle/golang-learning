## base

本小节来学习基本语法。

### 注释
Go 的注释使用 `//` 或者 `/**/` 来声明，与大多数语言一致，变量内的内容不会被编译。同时 Go 的注释还能起到类似于注解的作用(?)

### 变量

Go 的变量有两种声明方式，比如使用 `var` 关键词：

```go
var var type
```

还可以同时声明多个相同类型的变量：

```go
var var1, var2 type
```

这就类似于某些语言的：

```java
Type var1 var2;
```

还可以使用海象运算符[<sup>1</sup>](#refer)来声明：

```go
var := value
```

但是注意在使用海象运算符的时候，必须有至少一个新的变量被声明，不然编译器会报错 `no new variables on left side of :=`，错误示范见 [bad.go](bad.go)。

### 函数声明

Go 的函数声明使用 `func` 关键词：

```go
func func_name(parammeter type) return_type {
    // func body
}
```

### defer 关键词

使用 `defer` 关键词的表达式会在函数返回之前进行调用。

### OOP

#### 结构体

Go 的结构体定义如下：

```go
type struct_name struct {
    name type
}
```

结构体可以被初始化：

```go
type Shit struct {
    Color string
    Taste string
}
```

结构体可以有他的方法，是这么定义的：

```go
func (recv type) methodName(parameter_list) (return_value_list) { … }
```

一个完整的例子：

```go
type Shit struct {
	Color string
	Taste string
}

func (s Shit) eat() string {
	return fmt.Sprintf("You ate a %s shit! It smells like %s", s.Color, s.Taste)
}

func main() {
	fmt.Println(Shit{Color: "green", Taste: "rotten eggs"}.eat())
}
```

我们还可以利用指针，声明出类似于其他语言的 OOP 中的*静态函数*。

```go
func (s *Shit) init() *Shit {
	return &Shit{Color: "green", Taste: "rotten eggs"}
}
```


### 引用/声明

<div id="refer" />

[1] 海象运算符原版是 Python 里对 `:=` 符号的描述，这里借用来描述 `:=` 符号。