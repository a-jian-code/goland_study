package main

import "fmt"

// 包内变量,可以使用括号
var (
	aa = 3
	ss = "kkk"
	bb = true
)

// 空值
func variableZeroValue() {
	var a int
	var s string
	fmt.Println("%d %q\n", a, s)

}

// 连续赋值，:=缩写使用只能在函数内使用
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 缩写
func variableShorter() {
	a, b, c, s := 2, 3, false, "short"
	b = 10
	fmt.Println(a, b, c, s)

}

func main() {
	variableZeroValue()
	variableTypeDeduction()
	variableShorter()
	println(aa, ss, bb)
}
