class SolutionV2 {
    public int rangeBitwiseAnd(int left, int right) {
        while (left < right) {
            right &= right - 1;
        }
        return right;
    }
}