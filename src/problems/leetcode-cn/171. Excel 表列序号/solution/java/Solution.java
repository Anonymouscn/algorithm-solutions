class Solution {
    public int titleToNumber(String columnTitle) {
        byte[] arr = columnTitle.getBytes();
        int result = 0;
        for (int i = 0; i < arr.length; i++) {
            result = result * 26 + (arr[i] - 'A' + 1);
        }
        return result;
    }
}