// https://www.geeksforgeeks.org/find-repetitive-element-1-n-1/?ref=lbp

package main

import "fmt"

func sumOfArray(xs []int) int {
	result := 0
	for _, val := range xs {
		result += val
	}
	return result
}

func searchElem(xs []int) int {
	sumOfNElems := (len(xs) * (len(xs) - 1)) / 2
	sumOfArray := sumOfArray(xs)
	return sumOfArray - sumOfNElems
}

func main() {
	arr := []int{9, 8, 2, 6, 1, 8, 5, 3, 4, 7}
	fmt.Println(searchElem(arr))
}
