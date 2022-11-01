// https://www.geeksforgeeks.org/program-for-array-rotation-continued-reversal-algorithm/

package main

import "fmt"

func reverse(xs []int) []int {
	for i := 0; i < len(xs)/2; i++ {
		temp1 := xs[len(xs)-i-1]
		xs[len(xs)-i-1] = xs[i]
		xs[i] = temp1
	}
	return xs
}

func rotateArray(xs []int, d int) []int {
	k := d % len(xs)
	x := reverse(xs)
	ind := len(xs) - k
	leftArray := reverse(x[:ind])
	rightArray := reverse(x[ind:])
	return append(leftArray, rightArray...)
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateArray(_list, 22))
}
