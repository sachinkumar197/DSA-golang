// https://www.geeksforgeeks.org/given-an-array-a-and-a-number-x-check-for-pair-in-a-with-sum-as-x/?ref=lbp

package main

import "fmt"

func searchElem(xs []int, val int) string {
	_hash := make(map[int]int)
	for i := 0; i < len(xs); i++ {
		searchKey := val - xs[i]
		if _, ok := _hash[searchKey]; ok {
			return "Yes"
		} else {
			_hash[xs[i]] = 0
		}
	}
	return "No"
}

func main() {
	xs := []int{0, -1, 2, -3, 1}
	val := -2
	fmt.Println(searchElem(xs, val))
}
