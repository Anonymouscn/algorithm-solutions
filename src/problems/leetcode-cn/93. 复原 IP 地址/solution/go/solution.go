type Node struct {
    Val string // 当前值
    Selected []string // 已选字符组
    Layer int // 当前层数
    Children []*Node // 子节点
}

func restoreIpAddresses(s string) []string {
    result, selected := make([]string, 0), make([]string, 0)
    root := &Node{
        Val: s,
        Layer: 0,
        Selected: selected,
    }
    root = reverse(root, func(node *Node) {
        result = append(result, strings.Join(node.Selected, ".") + fmt.Sprintf(".%v", node.Val))
    })
    return result
}

func reverse(node *Node, dol func(node *Node)) (*Node) {
    if node.Layer > 3 {
        return node
    }
    children, length := make([]*Node, 0), len(node.Val)
    for i := 1; i < length && i <= 3; i++ {
        t := node.Val[:i]
        if len(t) > 1 && t[0] == '0' {
            break
        }
        if cal(t) <= 255 {
            selected := cloneArr(node.Selected)
            selected = append(selected, t)
            children = append(children, &Node{
                Val: node.Val[i:],
                Layer: node.Layer + 1,
                Selected: selected,
            })
        }
    }
    for _, c := range children {
        c = reverse(c, dol)
    }
    node.Children = children
    if node.Layer == 3 {
        if l := len(node.Val); l == 1 || (l > 1 && node.Val[0] != '0') && cal(node.Val) <= 255 {
            dol(node)
        }
    }
    return node
}

func cal(s string) int {
    length, result := len(s), 0
    for i, pow := length - 1, 0; i >= 0; i-- {
        result += (int(math.Pow(10, float64(pow))) * int(s[i] - '0'))
        pow++
    }
    return result
}

func cloneArr(arr []string) []string {
    length := len(arr)
    result := make([]string, length)
    for i := 0; i < length; i++ {
        result[i] = arr[i]
    }
    return result
}