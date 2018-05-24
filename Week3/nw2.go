package main

import "fmt"

var (
	Up   rune = 1
	Left rune = 2
	NW   rune = 3
	None rune = 4
)

func main() {
	str := "學問是資管系"
	str1 := "資管系有帥哥"
	str2 := "資管界美女"
	str3 := "淡江大學有很好的資管系"

	r := []rune(str)
	r1 := []rune(str1)
	r2 := []rune(str2)
	r3 := []rune(str3)

	fmt.Println(Align(r, r1, 0, -3, -2))
	fmt.Println(Align(r, r2, 0, -3, -2))
	fmt.Println(Align(r, r3, 0, -3, -2))
}

func idx(i, j, bLen int) int {
	return (i * bLen) + j
}

func Align(a, b []rune, match, mismatch, gap int) (alignA, alignB string, score int) {

	aLen := len(a) + 1
	bLen := len(b) + 1

	maxLen := aLen
	if maxLen < bLen {
		maxLen = bLen
	}

	aBytes := make([]rune, 0, maxLen)
	bBytes := make([]rune, 0, maxLen)

	f := make([]int, aLen*bLen)
	pointer := make([]rune, aLen*bLen)

	for i := 1; i < aLen; i++ {
		f[idx(i, 0, bLen)] = gap * i
		pointer[idx(i, 0, bLen)] = Up
	}
	for j := 1; j < bLen; j++ {
		f[idx(0, j, bLen)] = gap * j
		pointer[idx(0, j, bLen)] = Left
	}

	pointer[0] = None

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
	}

	i := aLen - 1
	j := bLen - 1

	score = f[idx(i, j, bLen)]

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
	}

	reverse(aBytes)
	reverse(bBytes)

	return string(aBytes), string(bBytes), score
}

func reverse(a []rune) {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
}
