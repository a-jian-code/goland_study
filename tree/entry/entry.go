package main

import (
	"fmt"
	"goland/tree"
)

// 封装
// 首字母大写:public
// 首字母小写:private

// 每个目录一个包
// 为结构定义的方法必须放在同一个包内
// 可以是不同文件，但包名要一致

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}
func main() {
	//创建树
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)
	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
	//root.print()
	//root.setValue(30)
	////定义空指针
	//var proot *treeNode
	//proot.setValue(1)
	//root.print()
	//myroot := myTreeNode{&root}
	//myroot.postOrder()
}
