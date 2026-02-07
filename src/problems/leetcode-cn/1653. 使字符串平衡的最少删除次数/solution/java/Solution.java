class Solution {
    public int minimumDeletions(String s) {
        int length = s.length();
        if (length < 2) return 0;

        int bc = 0, del = 0;

        for (byte b : s.getBytes()) {
            if (b == 'b') bc++;
            else
                if (++del > bc) del = bc;
        }

        return del;
    }
}