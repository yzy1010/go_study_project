package main

import "fmt"

// addTen 接收一个整数指针，将指针指向的值增加 10。
func addTen(n *int) {
	*n += 10
}

// doubleSlice 接收一个整数切片的指针，将切片中的每个元素乘以 2。
func doubleSlice(s *[]int) {
	for i := range *s {
		(*s)[i] *= 2
	}
}

func main() {
	// 指针示例：传入变量地址，函数内部修改原值
	x := 5
	fmt.Println("原始 x:", x)
	addTen(&x)
	fmt.Println("调用 addTen(&x) 后 x:", x)

	// 切片示例：传入切片指针，在函数内部修改切片元素
	arr := []int{1, 2, 3, 4}
	fmt.Println("原始 切片:", arr)
	doubleSlice(&arr)
	fmt.Println("调用 doubleSlice(&arr) 后 切片:", arr)
}
