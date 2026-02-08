/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
    // 1.判断是不是平衡二叉树
    // 2.如果是平衡二叉树直接返回根节点，否则重建二叉树

    treeList, head := make([]*TreeNode, 0), root

    var abs func(n int) int
    abs = func(n int) int {
        if n < 0 {
            return -n
        }
        return n
    }

    // 检查二叉树是否平衡并备份节点 (此处追求性能极致一遍遍历，工程化最好遵循单一职责)
    var check_and_backup func(root *TreeNode) (bool, int)
    check_and_backup = func(root *TreeNode) (bool, int) {
        if root == nil {
            return true, 0
        }
        l, r, l_ok, r_ok := 0, 0, true, true
        if root.Left != nil {
            l_ok, l = check_and_backup(root.Left)
        }
        treeList = append(treeList, root)
        if root.Right != nil {
            r_ok, r = check_and_backup(root.Right)
        }
        return l_ok && r_ok && abs(l-r) <= 1, max(l, r)+1
    }

    // 重建二叉树 (二分后按照中序遍历，保证二叉树平衡)
    var build func(treeList []*TreeNode) *TreeNode
    build = func(treeList []*TreeNode) *TreeNode {
        length := len(treeList)
        if length == 0 {
            return nil
        }
        mid := (length-1) / 2
        root := treeList[mid]
        root.Left, root.Right = nil, nil
        left, right := treeList[:mid], []*TreeNode{}
        if mid+1 < length {
            right = treeList[mid+1:]
        }
        
        if len(left) > 0 {
            root.Left = build(left)
        }
        if len(right) > 0 {
            root.Right = build(right)
        }

        return root
    }

    isBalanced, _ := check_and_backup(root)
    if !isBalanced {
        head = build(treeList)
    }

    return head
}