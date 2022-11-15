// https://www.geeksforgeeks.org/find-repetitive-element-1-n-1/?ref=lbp

package main

import "fmt"

func searchElem(xs []int) int {
	tempMap := make(map[int]int)
	for i := 0; i < len(xs); i++ {
		if _, ok := tempMap[xs[i]]; ok {
			return xs[i]
		} else {
			tempMap[xs[i]] = 1
		}
	}
	return -1
}

func main() {
	x := []int{9, 8, 2, 6, 1, 8, 5, 3, 4, 7}
	fmt.Println(searchElem(x))
}
