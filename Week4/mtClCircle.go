package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Accuracy
const (
	N1 int = 512
	N2 int = 2048
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var cnt int
	for i := 0; i <= N1; i++ {
		x := rand.Float64() * 4.0
		y := rand.Float64() * 2.0
		if compare(x, y) {
			cnt++
		}
	}
	fmt.Printf("N = 512, result = %v\n", float64(cnt)/float64(N1)*8)

	cnt = 0

	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= N2; i++ {
		x := rand.Float64() * 4.0
		y := rand.Float64() * 2.0
		if compare(x, y) {
			cnt++
		}
	}
	fmt.Printf("N = 2048, result = %v\n", float64(cnt)/float64(N2)*8)
}

func compare(x, y float64) bool {
	return (x*x+y*y <= 2*2) && ((x-2)*(x-2)+y*y <= 2*2)
}
