// https://www.geeksforgeeks.org/search-insert-and-delete-in-a-sorted-array/?ref=lbp

package main

import "fmt"

func binarySearch(xs []int, start int, end int, key int) int {
	mid := (start + end) / 2
	_var := fmt.Sprintf("start: %d, end: %d, mid: %d", start, end, mid)
	fmt.Println(_var)
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

func deleteElem(xs []int, d int) []int {
	_index := binarySearch(xs, 0, len(xs)-1, d)
	if _index != -1 {
		for i := _index; i < len(xs)-1; i++ {
			xs[i] = xs[i+1]
		}
		return xs[:len(xs)-1]
	}
	return xs
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(deleteElem(_list, 5))
}
