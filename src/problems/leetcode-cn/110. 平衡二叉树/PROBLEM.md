# 110. 平衡二叉树

<p style="color: #1cb8b8; display: inline-flex; item-align: center; justify-content: center; padding: 2px 10px; border: 2px solid #666; border-radius: 20px">简单</p>

给定一个二叉树，判断它是否是 **平衡二叉树**

> 平衡二叉树 是指该树所有节点的左右子树的高度相差不超过 `1`。



**示例 1：**

![https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)

> **输入：** root = [3,9,20,null,null,15,7] \
> **输出：** true


**示例 2：**

![https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)

> **输入：** root = [1,2,2,3,3,null,null,4,4] \
> **输出：** false


**示例 3：**

> **输入：** root = [] \
> **输出：** true



**提示：**

* 树中的节点数在范围 `[0, 5000]` 内
* $-10^4$ <= Node.val <= $10^4$


## 参考引用
[题目链接 - LeetCode(CN) - 110. 平衡二叉树](https://leetcode.cn/problems/balanced-binary-tree)