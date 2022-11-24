// https://www.geeksforgeeks.org/given-an-array-a-and-a-number-x-check-for-pair-in-a-with-sum-as-x/?ref=lbp

package main

import (
	"fmt"
	"sort"
)

func searchElem(xs []int, left int, right int, val int) string {
	sort.Ints(xs)
	if left >= right {
		return "No"
	} else if xs[left]+xs[right] == val {
		return "Yes"
	} else if xs[left]+xs[right] > val {
		return searchElem(xs, left, right-1, val)
	} else if xs[left]+xs[right] < val {
		return searchElem(xs, left+1, right, val)
	}
	return "No"
}

func main() {
	xs := []int{0, -1, 2, -3, 1}
	val := -2
	fmt.Println(searchElem(xs, 0, len(xs)-1, val))
}
