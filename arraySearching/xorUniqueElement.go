// https://www.geeksforgeeks.org/find-element-appears-array-every-element-appears-twice/?ref=lbp

package main

import "fmt"

func searchElem(xs []int) int {
	temp := xs[0]
	for i := 1; i < len(xs); i++ {
		temp = temp ^ xs[i]
	}
	return temp
}

func main() {
	x := []int{1, 1, 2, 2, 3, 4, 4, 5, 5}
	fmt.Println(searchElem(x))
}
