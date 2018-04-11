package main

import "fmt"

func main() {
	arr := []int{2, 3, 45, 11, 1, 6, 67, 9}
	qsort(arr)
	fmt.Println("Sorted array is: ", arr)
}

func qsort(target []int) {
	if len(target) <= 1 {
		return
	}
	mid := target[0]
	fmt.Println("Pivot: ", mid)
	fmt.Println(target)
	first, last := 0, len(target)-1
	for i := 1; i <= last; {
		if target[i] > mid {
			target[i], target[last] = target[last], target[i]
			last--
		} else {
			target[i], target[first] = target[first], target[i]
			first++
			i++
		}
	}
	target[first] = mid
	qsort(target[:first])
	qsort(target[first+1:])
}
