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
	x := reverse(xs)
	fmt.Println(x)
	ind := len(xs) - d
	leftArray := reverse(x[:ind])
	fmt.Println(leftArray)
	rightArray := reverse(x[ind:])
	fmt.Println(rightArray)
	return append(leftArray, rightArray...)
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateArray(_list, 2))
}
