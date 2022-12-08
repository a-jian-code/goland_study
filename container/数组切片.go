package main

import "fmt"

// 切片,slice本身没有数据，是对底层array的一个view。可以通过slice来修改数组的值
func slice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])
}

// 修改数组
func updateArray(arr []int) {
	s := arr[2:6]
	s[0] = 10
}

// 反复修改,slice会拓展，如果重复切片时范围大于第一次切片长度且小于数组范围，则可以拓展未参与切片部分
func reslice(arr []int) []int {
	s := arr[2:3]
	s = s[:3]
	s = s[1:]
	return s
}
func main() {
	//slice()
	arr := [...]int{0, 1, 1, 1, 1, 1}
	fmt.Println(reslice(arr[:]))
	//updateArray(arr[:])
	//fmt.Println(arr)

	fmt.Println("Extending slice")
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr1[2:6]
	s2 := s1[3:5]
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))
}
