// https://www.geeksforgeeks.org/given-an-array-a-and-a-number-x-check-for-pair-in-a-with-sum-as-x/?ref=lbp

package main

import (
	"fmt"
	"sort"
)

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

func searchElem(xs []int, val int) string {
	sort.Ints(xs)
	for i := 0; i < len(xs)-1; i++ {
		searchKey := val - xs[i]
		j := binarySearch(xs, i+1, len(xs)-1, searchKey)
		if j != -1 {
			return "True"
		}
	}
	return "False"
}

func main() {
	xs := []int{0, -1, 2, -3, 1}
	val := -2
	fmt.Println(searchElem(xs, val))
}
