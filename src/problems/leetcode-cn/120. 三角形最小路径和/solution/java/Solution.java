class Solution {
    public int minimumTotal(List<List<Integer>> triangle) {
        for (int i = triangle.size() - 2; i >= 0; i--) {
            List<Integer> row = triangle.get(i);
            for (int j = 0; j < row.size(); j++) {
                int n1 = triangle.get(i+1).get(j), n2 = triangle.get(i+1).get(j+1);
                row.set(j, row.get(j) + (n1 <= n2 ? n1 : n2));
            }
            triangle.set(i, row);
        }
        return triangle.get(0).get(0);
    }
}