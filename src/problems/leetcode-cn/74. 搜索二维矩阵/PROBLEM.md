# 74. 搜索二维矩阵

<p style="color: #ffb800; display: inline-flex; item-align: center; justify-content: center; padding: 2px 10px; border: 2px solid #666; border-radius: 20px">中等</p>

给你一个满足下述两条属性的 `m x n` 整数矩阵：

* 每行中的整数从左到右按非严格递增顺序排列。
* 每行的第一个整数大于前一行的最后一个整数。

给你一个整数 `target` ，如果 `target` 在矩阵中，返回 `true` ；否则，返回 `false` 。



**示例 1：**

![https://assets.leetcode.com/uploads/2020/10/05/mat.jpg](https://assets.leetcode.com/uploads/2020/10/05/mat.jpg)

> **输入：** matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3 \
> **输出：** true


**示例 2：**

![https://assets.leetcode.cn/aliyun-lc-upload/uploads/2020/11/25/mat2.jpg](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2020/11/25/mat2.jpg)

> **输入：** matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13 \
> **输出：** false



**提示：**

* `m == matrix.length`
* `n == matrix[i].length`
* `1 <= m, n <= 100`
* $-10^4$ <= matrix[i][j], target <= $10^4$


## 参考引用
[题目链接 - LeetCode(CN) - 74. 搜索二维矩阵](https://leetcode.cn/problems/search-a-2d-matrix)