// https://www.geeksforgeeks.org/given-an-array-a-and-a-number-x-check-for-pair-in-a-with-sum-as-x/?ref=lbp

package main

import "fmt"

func searchElem(xs []int, val int) string {
	rem := make([]int, val)
	for i := 0; i < len(xs); i++ {
		if xs[i] < val {
			rem[xs[i]] += 1
		}
	}
	for i := 1; i < val; i++ {
		if rem[i] > 0 && rem[val-i] > 0 {
			return "Yes"
		}
	}
	return "No"
}

func main() {
	xs := []int{1, 4, 45, 6, 10, 8}
	val := 16
	fmt.Println(searchElem(xs, val))
}
