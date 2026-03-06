/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    var ptr *TreeNode
    var dfs func(root *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }

        // 旋转交换子树
        tmp := node.Right
        node.Right = node.Left
        node.Left = tmp

        if node.Left != nil {
            dfs(node.Left)
        }
        if node.Right != nil {
            dfs(node.Right)
        }

        // 连接右子树链
        if ptr != nil {
            node.Right = ptr
        }
        ptr = node
        node.Left = nil
    }
    dfs(root)
}