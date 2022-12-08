package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 欧拉公式  e**i*Pi+1=0
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量定义
func consts() {
	const (
		//不指定常量类型就会当作纯文本，类型可以自选
		a, b     = 3, 4
		filename = "text.txt"
	)
	var c int = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

// 枚举类型
func enums() {
	const (
		cpp = iota
		python
		goland
		_
		java
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb)
	fmt.Println(cpp, java)
}

func main() {
	//euler()
	//triangle()
	consts()
}
