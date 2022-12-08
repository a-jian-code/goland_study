package main

import (
	"fmt"
)

// 闭包累加器，返回值是一个函数
func adder() func(int) int {
	sum := 0 //自由变量
	return func(v int) int {
		sum += v
		return sum
	}
}

// 定义一个返回类型，返回值是此次的求和以及下一次的函数
type addr func(int) (int, addr)

// 定义求和函数，参数为基数
func add2(base int) addr {
	return func(v int) (int, addr) {
		fmt.Println("11111", base+v, add2(base+v))
		return base + v, add2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d=%d\n", i, a(i))
	}
	fmt.Println("----------------------------------------")
	a2 := add2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Printf("0+1+...+%d=%d\n", i, s)
	}
}
