// https://www.geeksforgeeks.org/find-element-appears-array-every-element-appears-twice/?ref=lbp

package main

import "fmt"

func dedupList(xs []int) []int {
	s := make(map[int]bool)
	for _, v := range xs {
		if _, ok := s[v]; !ok {
			s[v] = true
		}
	}
	keys := make([]int, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}

func sumArray(xs []int) int {
	result := 0
	for _, v := range xs {
		result += v
	}
	return result
}

func searchElem(xs []int) int {
	dList := dedupList(xs)
	sum1 := sumArray(dList) * 2
	sum2 := sumArray(xs)
	return (sum1 - sum2)
}

func main() {
	x := []int{1, 1, 2, 2, 3, 4, 4, 5, 5}
	fmt.Println(searchElem(x))
}
