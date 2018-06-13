package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("[0-1]: %v\n", intergral(0, 1, 1, 0, 10000))
	fmt.Printf("[0-2]: %v\n", intergral(0, 1, 2, 0, 10000))
}

func intergral(x, y, max, min float64, t int) float64 {
	var instance, sum float64
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < t; i++ {
		instance = (max-min)*rand.Float64() + min
		if (x <= instance) && (instance <= y) {
			sum += instance * instance
		}
	}
	sum *= y - x
	return sum / float64(t)
}
