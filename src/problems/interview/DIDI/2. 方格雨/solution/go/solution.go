package main

import (
	"fmt"
	"math"
)

func main() {
	n, m := 0, 0
	v, _ := fmt.Scan(&n, &m)
	if v != 2 {
		return
	}

	grids := make([]int, n)

	for i := 0; i < m; i++ {
		l, r, d := 0, 0, 0
		v, _ = fmt.Scan(&l, &r, &d)
		for j := l - 1; j <= r-1; j++ {
			grids[j] += d
		}
		if v != 3 {
			break
		}
	}

	minV, maxV := math.MaxInt, math.MinInt
	vm := make(map[int]int)
	for i := 0; i < n; i++ {
		if grids[i] < minV {
			minV = grids[i]
		}
		if grids[i] > maxV {
			maxV = grids[i]
		}
		if _, ok := vm[grids[i]]; ok {
			vm[grids[i]]++
		} else {
			vm[grids[i]] = 1
		}
	}
	if minV == 0 || float64(maxV) < math.Pow(10, 100) {
		vm[0]++
	}
	fmt.Println(len(vm))
}
