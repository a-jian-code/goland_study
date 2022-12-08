package tree

import "fmt"

// 重点：值接收者vs指针接收者
// 值接收者是go特有
// 要改变内容必须使用指针接收者
// 结构过大也考虑使用指针接收者
// 一致性：如果有执政接收者，最好都是指针接收者，指针接收者需要判读是否为nil
// 树形结构体
type Node struct {
	Value       int
	Left, Right *Node
}

// 工厂函数
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}

// 打印节点值
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 设置节点值,只有使用指针才可以改变结构内容
func (node *Node) SetValue(value int) {
	//空指针也可以调用函数
	if node == nil {
		fmt.Println("Setting vaule to nil Node. Ignored")
		return
	}
	node.Value = value
	fmt.Println(node.Value)
}
