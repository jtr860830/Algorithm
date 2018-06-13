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
			switch rand.Intn(10) {
			case 0:
				temp += 1
			case 1:
				fallthrough
			case 2:
				fallthrough
			case 3:
				temp += 2
			case 4:
				fallthrough
			case 5:
				fallthrough
			case 6:
				temp += 3
			case 7:
				temp += 4
			case 8:
				temp += 5
			case 9:
				temp += 6
			}
		}
		sum[temp-6]++
	}
	return
}
