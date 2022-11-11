// https://www.geeksforgeeks.org/search-insert-and-delete-in-an-unsorted-array/?ref=lbp

package main

import "fmt"

func insertElem(xs []int, n int, d int) []int {
	xs = append(xs[:n], append([]int{d}, xs[n:]...)...)
	return xs
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(insertElem(_list, 3, 40))
}
