class Solution {
public:
    vector<int> minBitwiseArray(vector<int>& nums) {
        for_each(nums.begin(), nums.end(), [](int& i) {            
            if ((i & 1) == 0) i = -1;
            else {
                int n = i + 1;
                i ^= ((n & -n) >> 1);
            }
        });
        return nums;
    }
};