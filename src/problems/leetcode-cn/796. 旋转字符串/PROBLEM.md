# 796. 旋转字符串

<p style="color: #1cb8b8; display: inline-flex; item-align: center; justify-content: center; padding: 2px 10px; border: 2px solid #666; border-radius: 20px">简单</p>

给定两个字符串, `s` 和 `goal`。如果在若干次旋转操作之后，`s` 能变成 `goal` ，那么返回 `true` 。

`s` 的 **旋转操作** 就是将 `s` 最左边的字符移动到最右边。 

 * 例如, 若 `s = 'abcde'`，在旋转一次之后结果就是 `'bcdea'` 。



**示例 1:**
> 输入: s = "abcde", goal = "cdeab" \
> 输出: true


**示例 2:**
> 输入: s = "abcde", goal = "abced" \
> 输出: false


**提示:**

 * `1 <= s.length, goal.length <= 100`
 * `s` 和 `goal` 由小写英文字母组成


## 参考引用
[题目链接 - LeetCode(CN) - 796. 旋转字符串](https://leetcode.cn/problems/rotate-string)