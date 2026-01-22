/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> result = new LinkedList<>();

        if (root == null) {
            return result;
        }

        List<Integer> tmp = new ArrayList<>();
        List<TreeNode> queue = new LinkedList<>();

        queue.add(root);
        int size = queue.size(), cnt = 0, layer = 0;

        while (queue.size() > 0) {
            TreeNode node = queue.removeFirst();
            tmp.add(node.val);

            if (node.left != null) {
                queue.add(node.left);
                cnt++;
            }
            if (node.right != null) {
                queue.add(node.right);
                cnt++;
            }

            if (--size == 0) {
                if (layer++ % 2 == 1) {
                    Collections.reverse(tmp);
                }
                result.add(tmp);
                tmp = new ArrayList<>();
                size = cnt;
                cnt = 0;
            }
        }

        return result;
    }
}