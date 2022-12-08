package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div(a, b)
		return q
	default:
		panic("unsupported operation:" + op)
	}
}

// 函数多个返回值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("Calling %s with %d %d \n", runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(), a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sumArgs(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}
func main() {
	//println(eval(1, 2, "/"))
	//println(eval(1, 2, "#"))
	//println(div(13, 4))
	//q, r := div(13, 4)
	//println(q, r)
	//fmt.Println(apply(pow, 3, 4))
	//fmt.Println(apply(func(i int, i2 int) int {
	//	return i + i2
	//}, 3, 4))
	fmt.Println(sumArgs(1, 2, 3, 4))
}
