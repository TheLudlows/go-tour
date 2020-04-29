package main

import "fmt"

func main() {

	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组⻓长度。
	c := [5]int{2: 100, 4: 200} // 使⽤用索引号初始化元素。

	d := [...]struct {
		// 可省略元素类型。
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}

	fmt.Println(a)
	updateArr(&a) // 默认值传递
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	e := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	f := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能⽤用 "..."。

	fmt.Println(e)
	fmt.Println(f)

}

func updateArr(a *[3]int) {
	a[0] = 88
}
