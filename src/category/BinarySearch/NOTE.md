# Binary Search - 二分查找


## 什么是二分查找?
[二分查找 - Wiki](https://zh.wikipedia.org/wiki/%E4%BA%8C%E5%88%86%E6%90%9C%E5%B0%8B) \
[二分查找 - OI Wiki](https://oi-wiki.org/basic/binary)

> 二分查找（Binary Search）是一种在 **有序** 数据中快速找目标值的方法：每次都取 **中间** 位置的元素和目标比较，把不可能包含目标的那一半直接丢掉，然后在剩下的一半里继续找，直到找到或范围为空。


## 应用场景
* 在 **有序** 数组 / 列表 中查某个值 ( **是否存在** / **位置下标** )

    如：在 **排序后** 的数据表里找某个数据。

* **有序** 数组 / 列表 中处理重复元素 (找区间范围)

    如：在 **有序** 数组里找某个值出现的 起始和结束位置 / 出现次数。

* **有序** 列表中插入位置 (保持有序)

    如：把新元素插入到 **有序** 列表中。

* 数值计算 / 逼近 ( **连续** 区间二分)

    如：求函数的零点、求平方根、在区间内 **逼近** 一个满足误差要求的值。

* 系统 / 工程中的阈值定位 (单调变化)

    如：压测时并发越大越容易超时，想找 **刚好** 不超时的最大并发。


## 常用模版
1. 扫描窗口 `[start, end)` - Rust 代码

    ```rust
    fn binary_search(list : &Vec<i32>, target: i32) -> Option<usize> {
        if list.is_empty() {
            return None;
        }
        let (mut left, mut right) = (0, list.len());
        while left < right { // 注意点1: '<' 扫描边界值
            let mid = (left + right) / 2;
            match list[mid].cmp(&target) {
                Ordering::Equal => return Some(mid),
                Ordering::Greater => {
                    right = mid; // 注意点2: 闭区间
                }
                Ordering::Less => left = mid + 1,
            }
        }
        None
    }
    ```

2. 扫描窗口 `[start, end]` - Rust 代码

    ```rust
    fn binary_search(list : &Vec<i32>, target: i32) -> Option<usize> {
        if list.is_empty() {
            return None;
        }
        let (mut left, mut right) = (0, list.len()-1);
        while left <= right { // 注意点1: '<=' 扫描边界值
            let mid = (left + right) / 2;
            match list[mid].cmp(&target) {
                Ordering::Equal => return Some(mid),
                Ordering::Greater => {
                    if mid == 0 { break; } // 注意点2: 防止扫描下标下溢 (0 - 1 = -1)
                    right = mid - 1; // 注意点3: 开区间
                }
                Ordering::Less => left = mid + 1,
            }
        }
        None
    }
    ```

模版均使用迭代法实现，同理也可递归实现。


## 算法题集
* [题目链接 - LeetCode(CN) - 74. 搜索二维矩阵](https://leetcode.cn/problems/search-a-2d-matrix)


## 更新信息
`2026.02.10`, `anonymous`