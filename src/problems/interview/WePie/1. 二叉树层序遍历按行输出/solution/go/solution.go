package main

import "fmt"

type TreeNode struct {
	Val int
	Layer int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	// 构造数据 (自动生成顺序完全二叉树)
	length := 7 // 生成树的节点长度
	nodes := make([]*TreeNode, length)
	for i := 0; i < length; i++ {
		nodes[i] = &TreeNode{
			Val: i + 1,
		}
		if i > 0 {
			parent := (i-1) / 2
			if nodes[parent] != nil {
				if (i - 1) % 2 == 0 {
					nodes[parent].Left = nodes[i]
				} else {
					nodes[parent].Right = nodes[i]
				}
			}
		}
	}
	// 执行逻辑代码
	nodes[0].Layer = 0
	queue, layer := []*TreeNode{nodes[0]}, 0
	for len(queue) > 0 {
		node := queue[0]
		nextLayer := node.Layer + 1
		if node.Layer != layer {
			layer++
			fmt.Println()
		}
		fmt.Printf("%v ", node.Val)
		if node.Left != nil {
			node.Left.Layer = nextLayer
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node.Right.Layer = nextLayer
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	// EOF 换行
	fmt.Println()
}