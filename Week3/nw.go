package main

import "fmt"

var (
	Up   byte = 1
	Left byte = 2
	NW   byte = 3
	None byte = 4
) // constant for direction

func main() {
	fmt.Println(Align("GTACCCCAT", "TGACCGCA", 0, -1, -2))
}

func idx(i, j, bLen int) int {
	return (i * bLen) + j
} // To make 1D array to 2D array

func Align(a, b string, match, mismatch, gap int) (alignA, alignB string, score int) {

	aLen := len(a) + 1
	bLen := len(b) + 1

	maxLen := aLen
	if maxLen < bLen {
		maxLen = bLen
	} // Find size that provides to make arrays

	aBytes := make([]byte, 0, maxLen)
	bBytes := make([]byte, 0, maxLen)

	f := make([]int, aLen*bLen)        // Array that provides scores
	pointer := make([]byte, aLen*bLen) // Array that provides direction

	for i := 1; i < aLen; i++ {
		f[idx(i, 0, bLen)] = gap * i
		pointer[idx(i, 0, bLen)] = Up
	}
	for j := 1; j < bLen; j++ {
		f[idx(0, j, bLen)] = gap * j
		pointer[idx(0, j, bLen)] = Left
	}

	pointer[0] = None // No direction

	for i := 1; i < aLen; i++ {
		for j := 1; j < bLen; j++ {
			matchMismatch := mismatch
			if a[i-1] == b[j-1] {
				matchMismatch = match
			}

			max := f[idx(i-1, j-1, bLen)] + matchMismatch
			hgap := f[idx(i-1, j, bLen)] + gap
			vgap := f[idx(i, j-1, bLen)] + gap

			if hgap > max {
				max = hgap
			}
			if vgap > max {
				max = vgap
			}

			p := NW
			if max == hgap {
				p = Up
			} else if max == vgap {
				p = Left
			}

			pointer[idx(i, j, bLen)] = p
			f[idx(i, j, bLen)] = max
		}
	} // Calculate scores and right direction

	i := aLen - 1
	j := bLen - 1

	score = f[idx(i, j, bLen)] // Find right score

	for p := pointer[idx(i, j, bLen)]; p != None; p = pointer[idx(i, j, bLen)] {
		if p == NW {
			aBytes = append(aBytes, a[i-1])
			bBytes = append(bBytes, b[j-1])
			i--
			j--
		} else if p == Up {
			aBytes = append(aBytes, a[i-1])
			bBytes = append(bBytes, '-')
			i--
		} else if p == Left {
			aBytes = append(aBytes, '-')
			bBytes = append(bBytes, b[j-1])
			j--
		}
	} // Put the char into result

	reverse(aBytes)
	reverse(bBytes) // Reverse to right order

	return string(aBytes), string(bBytes), score
}

func reverse(a []byte) {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
}
