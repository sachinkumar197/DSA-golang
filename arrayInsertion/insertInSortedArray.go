// https://www.geeksforgeeks.org/search-insert-and-delete-in-a-sorted-array/?ref=lbp

package main

import "fmt"

func insertElem(xs []int, key int) []int {
	xs = append(xs, 0)
	for i := 0; i < len(xs)-1; i++ {
		if key <= xs[i] {
			for j := len(xs) - 1; j > i; j-- {
				xs[j] = xs[j-1]
			}
			xs[i] = key
			return xs
		}
	}
	xs[len(xs)-1] = key
	return xs
}

func main() {
	_list := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(insertElem(_list, 60))
}
