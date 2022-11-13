// https://www.geeksforgeeks.org/search-insert-and-delete-in-an-unsorted-array/?ref=lbp

package main

import "fmt"

func findElem(xs []int, d int) int {
	for i := 0; i < len(xs); i++ {
		if xs[i] == d {
			return i
		}
	}
	return -1
}

func deleteElem(xs []int, d int) []int {
	_index := findElem(xs, d)
	for i := _index; i < len(xs)-1; i++ {
		xs[i] = xs[i+1]
	}
	return xs[:len(xs)-1]
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(deleteElem(_list, 3))
}
