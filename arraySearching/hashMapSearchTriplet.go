// https://www.geeksforgeeks.org/find-a-triplet-that-sum-to-a-given-value/?ref=lbp

package main

import "fmt"

func searchElem(xs []int, val int) bool {
	_hash := make(map[int]int)
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			tempVal := val - xs[i] - xs[j]
			if _, ok := _hash[tempVal]; ok {
				fmt.Print(xs[i], xs[j], tempVal, "\n")
				return true
			} else {
				_hash[xs[j]] = 1
			}
		}
	}
	return false
}

func main() {
	xs := []int{1, 4, 45, 6, 10, 8}
	val := 22
	searchElem(xs, val)
}
