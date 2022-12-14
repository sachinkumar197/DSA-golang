// https://www.geeksforgeeks.org/search-insert-and-delete-in-a-sorted-array/?ref=lbp

package main

import "fmt"

func binarySearch(xs []int, start int, end int, key int) int {
	mid := (start + end) / 2
	if xs[mid] == key {
		return mid
	} else if end <= start || start >= end {
		return -1
	} else if key > xs[mid] {
		return binarySearch(xs, mid+1, end, key)
	} else if key < xs[key] {
		return binarySearch(xs, start, mid, key)
	}
	return -1
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(binarySearch(_list, 0, len(_list), 6))
}
