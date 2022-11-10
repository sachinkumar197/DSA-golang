// https://www.geeksforgeeks.org/search-insert-and-delete-in-an-unsorted-array/?ref=lbp

package main

import "fmt"

func searchElem(xs []int, x int) int {
	for i := 0; i < len(xs); i++ {
		if xs[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(searchElem(_list, 40))
}
