/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }

    tree := []*TreeNode{}
    for ptr := head; ptr != nil; ptr = ptr.Next {
        v := ptr.Val
        tree = append(tree, &TreeNode{
            Val: v,
        })
    }

    var buildBalanceTree func(tree []*TreeNode) *TreeNode
    buildBalanceTree = func(tree []*TreeNode) *TreeNode {
        length := len(tree)
        mid := (length-1) / 2
        root := tree[mid]
        if mid > 0 {
            root.Left = buildBalanceTree(tree[:mid])
        }
        if mid < len(tree)-1 {
            root.Right = buildBalanceTree(tree[mid+1:])
        }
        return root
    }

    return buildBalanceTree(tree)
}