package main

import "fmt"

func printSlice(arr []int) {
	fmt.Printf("arr=%v,len=%d,cap=%d\n", arr, len(arr), cap(arr))
}

//func main() {
//	fmt.Println("slice append start")
//	//添加元素如果超过cap，系统重新分配更大的数组
//	//由于值传递的关系，必须接收append的返回值，s=append(s,val)
//	arr := [...]int{0, 1, 2, 3, 4, 5}
//	s := arr[2:4]
//	s1 := append(s, 11)
//	s2 := append(s1, 12)
//	fmt.Printf("arr=%v,s=%v,s1=%v,s2=%v", arr, s, s1, s2)
//}

//func main() {
//	fmt.Println("Creating slice")
//	//var s []int //Zero value for slice is nil
//	//for i := 0; i < 100; i++ {
//	//	s = append(s, 2*i+1)
//	//}
//	//fmt.Println(s)
//
//	s1 := []int{2, 4, 6, 8}
//	printSlice(s1)
//	//make方法创建数组,可以指定len和cap
//	s2 := make([]int, 16)
//	s3 := make([]int, 10, 32)
//	printSlice(s2)
//	printSlice(s3)
//
//}

func main() {
	fmt.Println("Copying slice")
	s1 := []int{2, 4, 6, 8}
	s2 := []int{3, 5, 7, 9}
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s3 := make([]int, 10)
	s3 = append(s2[2:], s3[:]...)
	printSlice(s3)

	fmt.Println("Popping from front")
	front := s3[0]
	s3 = s3[1:]

	fmt.Println("Popping from tail")
	tail := s3[len(s3)-1]
	s3 = s3[:len(s3)-1]

	fmt.Println(front, tail)
	fmt.Println(s3)
}
