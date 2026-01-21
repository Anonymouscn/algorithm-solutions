class Solution {
    public int[] plusOne(int[] digits) {
        int[] result = new int[digits.length+1];
        digits[digits.length-1]++;
        int r = 0;
        for (int i = digits.length-1; i >= 0; i--) {
            digits[i] += r;
            if (digits[i] > 9) {
                r = digits[i] / 10;
                digits[i] %= 10;
            } else {
                r = 0;
            }
            result[i+1] = digits[i];
        }
        if (r > 0) {
            result[0] = r;
        } else {
            result = digits;
        }
        return result;
    }
}