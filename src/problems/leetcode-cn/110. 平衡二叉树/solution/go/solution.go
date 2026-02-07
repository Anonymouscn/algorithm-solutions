/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    result, _ := cnt(root)
    return result
}

func cnt(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }

    l_ok, l := cnt(root.Left)
    if !l_ok {
        return false, -1
    }
    r_ok, r := cnt(root.Right)
    if !r_ok {
        return false, -1
    }

    if abs(l - r) > 1 {
        return false, -1
    }

    return true, max(l, r) + 1
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}