// 计算性能较优解
func calculate(s string) int {
    stack, sign := make([]int, 0), byte('+')
    num := 0
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            continue
        }
        isNumber := s[i] >= '0' && s[i] <= '9'
        if isNumber {
            num *= 10
            num += int(s[i] - '0')
            continue
        }
        stack = cal(stack, sign, num)
        sign, num = s[i], 0
    }
    stack = cal(stack, sign, num)
    return sum(stack)
}

func cal(stack []int, sign byte, num int) []int {
    switch sign {
        case '+':
            stack = append(stack, num)
        case '-':
            stack = append(stack, -num)
        case '*':
            stack[len(stack)-1] *= num
        case '/':
            stack[len(stack)-1] /= num
    }
    return stack
}

func sum(stack []int) int {
    sum := 0
    for i := 0; i < len(stack); i++ {
        sum += stack[i]
    }
    return sum
}