# 回溯

## 什么是回溯法?
[回溯法 - Wiki](https://zh.wikipedia.org/wiki/%E5%9B%9E%E6%BA%AF%E6%B3%95) \
[回溯法 - OI Wiki](https://oi-wiki.org/search/backtracking)

> 回溯法采用试错的思想，它尝试分步的去解决一个问题。在分步解决问题的过程中，当它通过尝试发现，现有的分步答案不能得到有效的正确的解答的时候，它将取消上一步甚至是上几步的计算，再通过其它的可能的分步解答再次尝试寻找问题的答案。

## 应用场景
* 排列
* 组合
* 穷举

## 解题模版
1. go 伪代码
```go
func solution(args ...any) any {
    result, path := any, make([]any, 0)

    var backtrack func(%args%) 
    backtrack = func(%args%) {
        if (%end_condition%) { // end condition
            %collect_result%
            return
        }
        for %range_case% { // dfs
            %cut_out_if_necessary%
            path = append(path, %current_element%)
            backtrack(%next_args%)
            path = path[:len(path)-1] // backtracking point
        }
    }

    backtrack(%start_args%)

    return result
}
```

## 算法题集
* [题目链接 - LeetCode(CN) - 39. 组合总和](https://leetcode.cn/problems/combination-sum)
* [题目链接 - LeetCode(CN) - 77. 组合](https://leetcode.cn/problems/combinations)
* [题目链接 - LeetCode(CN) - 93. 复原 IP 地址](https://leetcode.cn/problems/restore-ip-addresses)

## 更新信息
`2026.01.20`, `anonymous`