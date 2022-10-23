// https://www.geeksforgeeks.org/block-swap-algorithm-for-array-rotation/

package main

import "fmt"

func splitArray(L []int, R []int) []int {
	if len(L) > len(R) {
		Ll := L[:len(R)]
		Lr := L[len(R):]
		return append(R, splitArray(Lr, Ll)...)
	} else if len(L) < len(R) {
		Rl := R[:len(R)-len(L)]
		Rr := R[len(R)-len(L):]
		return append(splitArray(Rr, Rl), L...)
	} else {
		return append(R, L...)
	}
}

func rotateArray(xs []int, d int) []int {
	return splitArray(xs[:d], xs[d:])
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateArray(_list, 3))
}
