/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    result, path := make([]string, 0), []string{}

    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        path = append(path, strconv.Itoa(node.Val))

        if node.Left == nil && node.Right == nil {
            result = append(result, strings.Join(path, "->"))
        }

        if node.Left != nil {
            dfs(node.Left)
        }
        if node.Right != nil {
            dfs(node.Right)
        }

        path = path[:len(path)-1]
    }

    dfs(root)

    return result
}