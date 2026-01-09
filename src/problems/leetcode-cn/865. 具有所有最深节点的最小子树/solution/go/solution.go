/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNodeData struct {
    Node *TreeNode
    Layer int
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    queue, stack, pm := make([]*TreeNodeData, 0), make([]*TreeNodeData, 0), make(map[*TreeNode]*TreeNode)
    queue = append(queue, &TreeNodeData{
        Node: root,
        Layer: 1,
    })
    // 层序遍历树
    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]
        if left := item.Node.Left; left != nil {
            pm[left] = item.Node
            queue = append(queue, &TreeNodeData{
                Node: left,
                Layer: item.Layer + 1,
            })
        }
        if right := item.Node.Right; right != nil {
            pm[right] = item.Node
            queue = append(queue, &TreeNodeData{
                Node: right,
                Layer: item.Layer + 1,
            })
        }
        if length := len(stack); length > 0 && stack[length-1].Layer < item.Layer {
            stack = make([]*TreeNodeData, 0)
        }
        stack = append(stack, item)
    }
    // 反向搜索最近公共父节点
    for len(stack) > 1 {
        race := make(map[*TreeNode]int)
        for i := 0; i < len(stack); i++ {
            race[pm[stack[i].Node]] = stack[i].Layer - 1
        }
        stack = make([]*TreeNodeData, 0)
        for n, l := range race {
            stack = append(stack, &TreeNodeData{
                Node: n,
                Layer: l,
            })
        }
    }
    return stack[0].Node
}