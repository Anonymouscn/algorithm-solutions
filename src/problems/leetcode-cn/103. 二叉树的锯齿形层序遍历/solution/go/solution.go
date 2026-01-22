/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    queue, result, tmp := make([]*TreeNode, 0), make([][]int, 0), make([]int, 0)

    if root == nil {
        return result
    }

    queue = append(queue, root)
	// size: 当前层元素总数, cnt:下一层元素总数计数, layer: 当前层序号
    size, cnt, layer := len(queue), 0, 0
    
    for len(queue) > 0 {
        node := queue[0]
        queue, tmp = queue[1:], append(tmp, node.Val)
        
        if node.Left != nil {
            queue = append(queue, node.Left)
            cnt++
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
            cnt++
        }

        if size--; size == 0 {
            if layer % 2 == 0 {
                result = append(result, tmp)
            } else {
                result = append(result, reverse(tmp))
            }
            tmp, size, cnt = make([]int, 0), cnt, 0
            layer++
        }
    }
    return result
}

func reverse(arr []int) []int {
    length := len(arr)
    if length < 2 {
        return arr
    }
    for i, j := 0, length-1; i < j; {
        arr[i] ^= arr[j]
        arr[j] ^= arr[i]
        arr[i] ^= arr[j]
        i++
        j--
    }
    return arr
}