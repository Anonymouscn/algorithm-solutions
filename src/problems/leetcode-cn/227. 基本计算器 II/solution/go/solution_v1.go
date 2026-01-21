// 表达式计算：操作栈思路，可以拓展复杂运算

// ============= 栈结构定义 ============= //
type Stack[T any] struct {
    array []*T
}

func (s *Stack[T]) Reverse() {
    l, r := 0, s.Len()-1
    array := s.array
    var tmp *T
    for l < r {
        tmp = array[l]
        array[l] = array[r]
        array[r] = tmp
        l++
        r--
    }
}

func (s *Stack[T]) Len() int {
    return len(s.array)
}

func (s *Stack[T]) IsEmpty() bool {
    return s.Len() == 0
}

func (s *Stack[T]) IsNotEmpty() bool {
    return !s.IsEmpty()
}

func (s *Stack[T]) Push(item *T) {
    s.array = append(s.array, item)
}

func (s *Stack[T]) Pop() *T {
    if s.IsEmpty() {
        return nil
    }
    item := s.array[s.Len()-1]
    s.array = s.array[:s.Len()-1]
    return item
}

func (s *Stack[T]) Peek() *T {
    if s.IsEmpty() {
        return nil
    }
    return s.array[s.Len()-1]
}

func NewStack[T any]() *Stack[T] {
    array := make([]*T, 0)
    return &Stack[T]{
        array: array,
    }
}
// ============= 包装结构定义 ============= //
type Item[T any] struct {
    Val T
}
// ====================================== //

func calculate(s string) int {
    ns, ss := NewStack[Item[int]](), NewStack[Item[byte]]()
    var sb strings.Builder
    for i := 0; i < len(s); i++ {
        // 数字匹配
        j := i
        for  ; j < len(s) && ((s[j] >= '0' && s[j] <= '9') || s[j] == ' '); j++ {
            if s[j] == ' ' {
                continue
            }
            sb.WriteByte(s[j])
        }
        // 匹配到数字 (空值/非法值过滤)
        if j > i {
            // 存储匹配数字入栈
            num, _ := strconv.Atoi(sb.String());
            ns.Push(&Item[int]{
                Val: num,
            })
            sb.Reset()
            i = j
            // 乘除运算合并
            if ns.IsNotEmpty() && ss.IsNotEmpty() {
                sign := ss.Peek().Val
                if sign == '*' || sign == '/' {
                    ss.Pop()
                    v2, v1, result := ns.Pop().Val, ns.Pop().Val, 0
                    switch sign {
                        case '*':
                            result = v1 * v2
                        case '/':
                            result = v1 / v2
                    }
                    ns.Push(&Item[int]{
                        Val: result,
                    })
                }
            }
            // 运算符匹配
            if i < len(s) {
                ss.Push(&Item[byte]{
                    Val: s[i],
                })
            }
        }
    }
    // 栈方向反转
    ss.Reverse()
    ns.Reverse()
    // 尾部运算
    for ss.IsNotEmpty() {
        sign, v1, v2, result := ss.Pop().Val, ns.Pop().Val, ns.Pop().Val, 0
        switch sign {
            case '+':
                result = v1 + v2
            case '-':
                result = v1 - v2
            case '*':
                result = v1 * v2
            case '/':
                result = v1 / v2
        }
        ns.Push(&Item[int]{
            Val: result,
        })
    }
    return ns.Pop().Val
}