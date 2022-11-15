// https://www.geeksforgeeks.org/find-repetitive-element-1-n-1/?ref=lbp

package main

import "fmt"

func searchElem(xs []int) int {
	xor1 := 1
	for i := 2; i < len(xs); i++ {
		xor1 ^= i
	}
	xor2 := xs[0]
	for i := 1; i < len(xs); i++ {
		xor2 ^= xs[i]
	}
	return xor1 ^ xor2
}

func main() {
	arr := []int{9, 8, 2, 6, 1, 8, 5, 3, 4, 7}
	fmt.Println(searchElem(arr))
}
