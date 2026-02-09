/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    length := len(postorder)
    if length == 0 {
        return nil
    }

    root := &TreeNode {
        Val: postorder[length-1],
    }

    ptr := 0
    for ; ptr < length && inorder[ptr] != root.Val; ptr++ {}

    if ptr > 0 {
        root.Left = buildTree(inorder[:ptr], postorder[:ptr])
    }
    if s := ptr+1; s < length {
        root.Right = buildTree(inorder[s:], postorder[ptr:length-1])
    }

    return root
}