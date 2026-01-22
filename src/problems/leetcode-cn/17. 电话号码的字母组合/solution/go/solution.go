var dic = map[byte][]byte {
    '2': []byte{'a', 'b', 'c'},
    '3': []byte{'d', 'e', 'f'},
    '4': []byte{'g', 'h', 'i'},
    '5': []byte{'j', 'k', 'l'},
    '6': []byte{'m', 'n', 'o'},
    '7': []byte{'p', 'q', 'r', 's'},
    '8': []byte{'t', 'u', 'v'},
    '9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
    result, collection := []string{}, [][]byte{}

    var backtrack func(arr [][]byte, selected []byte, p int)
    backtrack = func(arr [][]byte, selected []byte, p int) {
        if p == len(arr) {
            result = append(result, string(selected))
            return
        }
        for i := 0; i < len(arr[p]); i++ {
            backtrack(arr, append(selected, arr[p][i]), p+1)
        }
    }

    for i := 0; i < len(digits); i++ {
        collection = append(collection, dic[digits[i]])
    }

    backtrack(collection, []byte{}, 0)

    return result
}