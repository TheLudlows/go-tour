package main

import "fmt"

func main() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	printSlice(numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 1)
	fmt.Println(&numbers1[0])
	printSlice(numbers1)
	numbers1 = append(numbers1, 4)
	fmt.Println(&numbers1[0])
	printSlice(numbers1)
	numbers1 = append(numbers1, 2)
	fmt.Println(&numbers1[0])
	printSlice(numbers1)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
