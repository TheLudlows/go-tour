#### 1. 语言结构
 - 包声明
  源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

- 引入包
 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素）
 - 函数/方法
  func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的
 - 变量/标识符
 标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
 - 表达式/语句
 不需要;结尾
 
#### 变量定义
三种方式
1. var identifier1, identifier2 type
2. var v_name = value
3. v_name := value

另外变量声明之后不使用也会报错,除非使用_标识符来接受，此标识符表示抛弃变量
#### 地址
&获取变量地址
*使用是指
例子：/basic/Basic.go:20


#### 常量

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：

const identifier [type] = value

#### 函数定义
```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```
函数可以返回多个值
另一种定义方式
```go
addFunc := func(a int, b int) int {
    return  a+b
}
println(addFunc(999 , 1))
```
函数可以作为参数传递


#### 变量作用域
- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

#### 结构体
定义格式
```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```


