/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    mask := math.MaxInt / 4
    var dfs func(root *TreeNode) (bool, int, int)
    dfs = func(root *TreeNode) (bool, int, int) {
        if root == nil {
            return true, mask, -mask
        }
        if root.Left == nil && root.Right == nil {
            return true, root.Val, root.Val
        }
        l_ok, l_min, l_max := dfs(root.Left)
        if !l_ok {
            return false, mask, -mask
        }
        r_ok, r_min, r_max := dfs(root.Right)
        if !r_ok {
            return false, mask, -mask
        }
        return l_max < root.Val && r_min > root.Val, min(l_min, r_min, root.Val), max(l_max, r_max, root.Val)
    }
    result, _, _ := dfs(root)
    return result
}