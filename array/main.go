package main

import (
	"fmt"
)

// slice的扩容,引起的len和cap的变化
func main() {
	arr := []int{1, 2, 3, 4, 5}
	result := getArr(arr)
	fmt.Printf("%v", result)
}

// 这个函数会将第一个index的数字放到slice的末尾
func getArr(arr []int) [][]int {
	var result [][]int
	fmt.Printf("pointer=%p, len=%v,cap=%v\n", &arr, len(arr), cap(arr))
	for index := 1; index < len(arr); index++ {
		a := append(arr[index:], arr[:index]...)
		fmt.Printf("pointer=%p, len=%v,cap=%v\n", &a, len(a), cap(a))
		result = append(result, a)
	}
	return result
}
