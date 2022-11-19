// https://www.geeksforgeeks.org/given-an-array-a-and-a-number-x-check-for-pair-in-a-with-sum-as-x/?ref=lbp

package main

import "fmt"

func searchElem(xs []int, val int) string {
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			if xs[i]+xs[j] == val {
				return "Yes"
			}
		}
	}
	return "No"
}

func main() {
	xs := []int{0, -1, 2, -3, 1}
	val := -2
	fmt.Println(searchElem(xs, val))
}
