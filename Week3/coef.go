package main

import "fmt"

func binoCoef(n, k int) int {
	t := make([][]int, n+1)
	for r := 0; r < n+1; r++ {
		t[r] = make([]int, r+1)
		left, right := 1, r-1
		t[r][left-1], t[r][right+1] = 1, 1
		for c := left; c <= right; c++ {
			t[r][c] = t[r-1][c-1] + t[r-1][c]
		}
	}
	return t[n][k]
}

func main() {
	fmt.Printf("C(30, 12) = %d\n", binoCoef(30, 12))
	fmt.Printf("C(50, 12) = %d", binoCoef(50, 12))
}
