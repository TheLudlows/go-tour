#### 语言结构
 - 包声明
  源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

- 引入包
 `import "fmt"` 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素）
 - 函数/方法
  `func main()` 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的
 - 变量/标识符
 标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
 - 表达式/语句不需要;结尾
 


另外变量声明之后不使用也会报错,除非使用_标识符来接受，此标识符表示抛弃变量
#### 地址
&获取变量地址,*使用是指
例子：/basic/Basic.go:20

#### 变量定义
三种方式
1. `var identifier1, identifier2 type`
2. `var v_name = value`
3. `v_name := value`
#### 常量

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。常量的定义格式：
const identifier [type] = value

#### 权限
Go语言没有像其它语言一样有public、protected、private等访问控制修饰符，它是通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、函数等的名称是大写字母开头表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头）

#### 匿名字段
匿名字段不过是⼀种语法糖，从根本上说，就是⼀个与成员类型同名 (不含包名) 的字段。
被匿名嵌⼊的可以是任何类型，当然也包括指针。
```go
type User struct {
 name string
}
type Manager struct {
 User
 title string
}
```
#### 函数定义
```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```
函数可以返回多个值。
另一种定义方式
```go
addFunc := func(a int, b int) int {
    return  a+b
}
println(addFunc(999 , 1))
```
函数可以作为参数传递,
参数列表可以改为（a,b int）

⽅法总是绑定对象实例，并隐式将实例作为第⼀实参,可⽤实例 value 或 pointer 调⽤全部⽅法，编译器⾃动转换。可以像字段成员那样访问匿名字段⽅法，编译器负责查找。

#### 变量作用域
- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

#### 结构体
Go语言的结构体(struct)和其他语言的类(class)有同等的地位
定义格式
```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```
#### 类型方法
在Go语言中，可以给任意类型(包括内置类型，但不包括指针类型)添加相应的方法，
```go
type Rect struct { 
    x, y float64
    width, height float64
}
func (r *Rect) Area() float64 { 
    return r.width * r.height
}
```
#### 可见性
没有private、protected、public这样的关键字。要使某个符号对其他包可见，需要将该符号定义为以大写字母开头。
成员方法的可访问性遵循同样的规则

#### 面向对象
一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口，
```go
type IFile interface {
    Read(buf []byte) (n int, err error)
    Write(buf []byte) (n int, err error)
    Seek(off int64, whence int) (pos int64, err error) Close() error
}
```

任何对象实例都满足空接口interface{},类型可实现多个接⼝。

`var v1 interface{} = 1`

#### 异常Error defer panic recover
error接口时内部定义的，使用只需要实现Error方法，具体看：[Error.go](src/basic/Error.go)
defer类似于Java中的final语句块，即被defer申明的语句一定会被执行，并且可以声明多个方法，执行顺序为先进后出。
panic是go内部的一个方法，类似于Java的throw关键字，正常的函数执行流程将立即终止，并将错误信息上报，defer声名的语句会被执行
recover方法是接受panic的错误信息，类似于catch
三者配合使用见[Error.go](src/basic/Error2.go)
#### 类型
值类型：变量直接存储的值，内存通常在栈中分配，int、float、bool、string、数组、struct。
引用类型：变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，通过GC回收。
ref r--->地址--->值，包括 指针、slice、map、chan、接口