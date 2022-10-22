// https://www.geeksforgeeks.org/array-rotation/
// Approach 2 (Rotate one by one):

package main

import "fmt"

func rotateByOne(x []int) []int {
	firstElem := x[0]
	for i := 0; i < len(x)-1; i++ {
		x[i] = x[i+1]
	}
	x[len(x)-1] = firstElem
	return x
}

func rotate(x []int, d int) []int {
	for i := 1; i <= d; i++ {
		x = rotateByOne(x)
	}
	return x
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotate(_list, 2))
}
