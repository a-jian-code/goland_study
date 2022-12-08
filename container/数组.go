package main

// 数组是值类型，也就是拷贝数组，不改变原数组。引用传递会改变原数组
func printArray(arr [5]int) {
	arr[0] = 1000
	for i, v := range arr {
		println(i, v)
	}
}

// 引用传递，将指针传入
func printArrayp(arr *[5]int) {
	arr[0] = 1001
	for i, v := range arr {
		println(i, v)
	}
}

// 可以使用_省略变量
func main() {
	var array1 [5]int
	//array2 := [3]int{1, 2, 3}
	//array3 := [...]int{4, 5, 6}
	//var grid [4][5]bool
	//fmt.Println(array1, array2, array3, grid)
	//原始遍历数组
	//for i := 0; i < len(array3); i++ {
	//	fmt.Println(array3[i])
	//}
	//range关键字遍历数组
	//for i := range array3 {
	//	fmt.Println(array3[i])
	//}
	//遍历数组的下标和值
	//for index, value := range array2 {
	//	fmt.Println(index, value)
	//}
	//printArray(array1)
	printArrayp(&array1)
	//fmt.Println(array1)
}
