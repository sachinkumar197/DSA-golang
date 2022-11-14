// https://www.geeksforgeeks.org/find-element-appears-array-every-element-appears-twice/?ref=lbp

package main

import "fmt"

func searchElem(xs []int) int {
	_hash := make(map[int]int)
	for i := 0; i < len(xs); i++ {
		if val, ok := _hash[xs[i]]; ok {
			_hash[xs[i]] = val + 1
		} else {
			_hash[xs[i]] = 1
		}
	}
	for k, v := range _hash {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	x := []int{1, 1, 2, 2, 3, 4, 4, 5, 5}
	fmt.Println(searchElem(x))
}
