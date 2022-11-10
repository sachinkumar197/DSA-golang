// https://www.geeksforgeeks.org/search-insert-and-delete-in-an-unsorted-array/?ref=lbp

package main

import "fmt"

func naiveInsert(xs []int, d int) []int {
	xs = append(xs, d)
	return xs
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(naiveInsert(_list, 5))
}
