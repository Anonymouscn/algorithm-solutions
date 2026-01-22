class Solution {
    static final Map<Character, char[]> DIC = Map.of(
        '2', new char[]{'a', 'b', 'c'},
        '3', new char[]{'d', 'e', 'f'},
        '4', new char[]{'g', 'h', 'i'},
        '5', new char[]{'j', 'k', 'l'},
        '6', new char[]{'m', 'n', 'o'},
        '7', new char[]{'p', 'q', 'r', 's'},
        '8', new char[]{'t', 'u', 'v'},
        '9', new char[]{'w', 'x', 'y', 'z'}
    );

    List<String> result = new LinkedList<>();

    public List<String> letterCombinations(String digits) {
        List<char[]> seq = new LinkedList<>();
        for (int i = 0; i < digits.length(); i++) {
            seq.add(DIC.get(digits.charAt(i)));
        }
        backtrack(seq, "", 0);
        return result;
    }

    void backtrack(List<char[]> arr, String selected, int p) {
        if (selected.length() == arr.size()) {
            result.add(selected);
            return;
        }
        char[] ac = arr.get(p);
        for (int i = 0; i < ac.length; i++) {
            backtrack(arr, selected + ac[i], p+1);
        }
    }
}