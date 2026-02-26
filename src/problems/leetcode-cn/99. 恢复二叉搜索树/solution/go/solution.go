/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    v1, v2, reverse := root, root, false
    q := make([]*TreeNode, 0)

    var swap func(v1, v2 *TreeNode)
    swap = func(v1, v2 *TreeNode) {
        tmp := v1.Val
        v1.Val = v2.Val
        v2.Val = tmp
    }

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node.Left != nil {
            dfs(node.Left)
        }
        
        if len(q) < 2 {
            q = append(q, node)
        } else {
            q[0] = q[1]
            q[1] = node
        }
        if len(q) == 2 && q[0].Val > q[1].Val {
            if reverse {
                v2 = q[1]
            } else {
                v1, v2 = q[0], q[1]
                reverse = true
            }
        }

        if node.Right != nil {
            dfs(node.Right)
        }
    }

    dfs(root)
    swap(v1, v2)
}