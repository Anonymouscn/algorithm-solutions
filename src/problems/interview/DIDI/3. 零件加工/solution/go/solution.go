package main

import (
	"fmt"
	"math"
)

func main() {
	T := 0
	v, _ := fmt.Scan(&T)
	if v != 1 {
		return
	}

	for r := 0; r < T; r++ {
		n, k := 0, 0
		v, _ = fmt.Scan(&n, &k)
		if v != 2 {
			break
		}
		a, b := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			v, _ = fmt.Scan(&a[i], &b[i])
			if v != 2 {
				break
			}
		}

		h := make([]int, n)
		minH := math.MaxInt
		for i := 0; i < n; i++ {
			h[i] = a[i] + b[i]
			if h[i] < minH {
				minH = h[i]
			}
		}

		cutH := make([]int, n)
		totalCut := 0
		for i := 0; i < n; i++ {
			cutH[i] = h[i] - minH
			totalCut += cutH[i]
		}
		fmt.Println(totalCut)
	}
}

/* 输入
1
4 3
3 6
5 2
4 7
8 9
*/
