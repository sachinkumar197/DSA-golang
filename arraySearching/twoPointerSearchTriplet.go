// https://www.geeksforgeeks.org/find-a-triplet-that-sum-to-a-given-value/?ref=lbp

package main

import (
	"fmt"
	"sort"
)

func twoPointerSearch(xs []int, left int, right int, val int) string {
	sort.Ints(xs)
	if left >= right {
		return "No"
	} else if xs[left]+xs[right] == val {
		fmt.Print(xs[left], " ")
		fmt.Print(xs[right], " ")
		return "Yes"
	} else if xs[left]+xs[right] > val {
		return twoPointerSearch(xs, left, right-1, val)
	} else if xs[left]+xs[right] < val {
		return twoPointerSearch(xs, left+1, right, val)
	}
	return "No"
}

func searchElem(xs []int, val int) bool {
	sort.Ints(xs)
	for i := 0; i < len(xs)-2; i++ {
		result := twoPointerSearch(xs, i+1, len(xs)-1, val-xs[i])
		if result == "Yes" {
			fmt.Print(xs[i], "\n")
			return true
		}
	}
	return false
}

func main() {
	xs := []int{1, 4, 45, 6, 10, 8}
	val := 22
	searchElem(xs, val)
}
