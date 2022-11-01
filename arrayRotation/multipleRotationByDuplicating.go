// https://www.geeksforgeeks.org/quickly-find-multiple-left-rotations-of-an-array/

package main

import "fmt"

func rotateArray(xs []int, k int) []int {
	i := k % len(xs)
	j := i + len(xs)
	temp_list := append(xs, xs...)
	return temp_list[i:j]
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateArray(_list, 51))
}
