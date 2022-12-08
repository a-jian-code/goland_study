package main

// 改变指针
func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	a, b = swap1(a, b)
	//fmt.Println(a, b)

}
