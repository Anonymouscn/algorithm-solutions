# 109. 有序链表转换二叉搜索树

<p style="color: #ffb800; display: inline-flex; item-align: center; justify-content: center; padding: 2px 10px; border: 2px solid #666; border-radius: 20px">中等</p>

给定一个单链表的头节点 `head` ，其中的元素 **按升序排序** ，将其转换为 **`平衡` 二叉搜索树** 。



**示例 1:**

![https://assets.leetcode.com/uploads/2020/08/17/linked.jpg](https://assets.leetcode.com/uploads/2020/08/17/linked.jpg)

> **输入:** head = [-10,-3,0,5,9] \
> **输出:** [0,-3,9,-10,null,5] \
> **解释:** 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。


**示例 2:**

> **输入:** head = [] \
> **输出:** []



**提示:**

* `head` 中的节点数在 [0, 2 * $10^4$] 范围内
* $-10^5$ <= Node.val <= $10^5$


## 参考引用
[题目链接 - LeetCode(CN) - 109. 有序链表转换二叉搜索树](https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree)