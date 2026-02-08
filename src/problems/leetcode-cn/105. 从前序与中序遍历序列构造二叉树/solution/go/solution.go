/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    root := &TreeNode{
        Val: preorder[0],
    }
    riptr := 0
    // 找根节点在中序遍历的位置
    for length := len(inorder); riptr < length && inorder[riptr] != preorder[0]; riptr++ {}

    // 存在左子树
    if riptr > 0 {
        // nextLeftInorder, nextLeftPreorder := inorder[:riptr], preorder[1:epptr+1]
        root.Left = buildTree(preorder[1:riptr+1], inorder[:riptr])
    }

    // 存在右子树
    if length := len(inorder); riptr+1 < length {
        // nextRightInorder, nextRightPreorder := inorder[riptr+1:], preorder[epptr+1:]
        root.Right = buildTree(preorder[riptr+1:], inorder[riptr+1:])
    }

    return root
}