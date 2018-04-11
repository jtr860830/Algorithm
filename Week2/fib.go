package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Recursive")
	for i := 8; i <= 40; i += 8 {
		start := time.Now()
		fibRec(i)
		elapsed := time.Since(start)
		fmt.Printf("Fib(%d) ans: %d, time: %s\n", i, fibRec(i), elapsed)
	}

	fmt.Println("\nLoop")
	for i := 8; i <= 40; i += 8 {
		start := time.Now()
		fibLoop(i)
		elapsed := time.Since(start)
		fmt.Printf("Fib(%d) ans: %d, time: %s\n", i, fibLoop(i), elapsed)
	}
}

func fibRec(n int) int {
	if n <= 1 {
		return n
	}
	return fibRec(n-1) + fibRec(n-2)
}

func fibLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
