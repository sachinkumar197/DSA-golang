// https://www.geeksforgeeks.org/find-repetitive-element-1-n-1/?ref=lbp

package main

import (
	"fmt"
	"sort"
)

func searchElem(xs []int) int {
	sort.Ints(xs)
	for i := 0; i < len(xs); i++ {
		if xs[i] != i+1 {
			return xs[i]
		}
	}
	return -1
}

func main() {
	x := []int{9, 8, 2, 6, 1, 8, 5, 3, 4, 7}
	fmt.Println(searchElem(x))
}
