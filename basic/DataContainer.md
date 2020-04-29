#### 1. Arrray

-  数组是值类型，赋值和传参会复制整个数组，而不是指针。
-  数组长度必须是常量，且是类型的组成部分。[2]int 和 [3]int 是不同类型。 • ⽀支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
-  指针数组 [n]*T，数组指针 *[n]T。
- 值拷⻉贝⾏行为会造成性能问题，通常会建议使⽤用 slice，或数组指针。
- 内置函数 len 和 cap 都返回数组长度 (元素数量)。
[Array.go](./Array.go)
#### 2. Slice
``` c
struct Slice {
    byte*    array;// actual data
    uintgo   len;// number of elements
    uintgo   cap;// allocated number of elements
}
```

- 引⽤用类型。但⾃⾝是结构体，值拷贝传递。
- 属性 len 表⽰可⽤用元素数量，读写操作不能超过该限制。 
- 属性 cap 表示最⼤扩张容量，不能超出数组限制。
- append向 slice 尾部添加数据，返回新的 slice 对象。⼀一旦超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
[Slice.go](./Slice.go)
