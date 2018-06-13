package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for k, v := range rolling(10000) {
		fmt.Printf("[%2d], %4d times\n", k+6, v)
	}
}

func rolling(t int) (sum []int) {
	rand.Seed(time.Now().UnixNano())
	sum = make([]int, 31)

	for i := 0; i < t; i++ {
		var temp int
		for j := 0; j < 6; j++ {
			roll := rand.Intn(6) + 1
			temp += roll
		}
		sum[temp-6]++
	}
	return
}
