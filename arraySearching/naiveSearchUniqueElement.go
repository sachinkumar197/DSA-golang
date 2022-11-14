//  https://www.geeksforgeeks.org/find-element-appears-array-every-element-appears-twice/?ref=lbp

package main

import "fmt"

func searchElem(xs []int) int {
	i := 0
	for i < len(xs) {
		count := 0
		j := 0
		for j < len(xs) {
			if i != j && xs[i] == xs[j] {
				count += 1
				break
			}
			j += 1
		}
		if count == 0 {
			return xs[i]
		}
		i += 1
	}
	return -1
}

func main() {
	x := []int{1, 1, 2, 2, 3, 4, 4, 5, 5}
	fmt.Println(searchElem(x))
}
