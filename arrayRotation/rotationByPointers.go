// https://www.geeksforgeeks.org/c-program-cyclically-rotate-array-one/

package main

import "fmt"

func rotateByOne(xs []int) []int {
	i := 0
	j := len(xs) - 1
	for i != j {
		xs[i], xs[j] = xs[j], xs[i]
		i += 1
	}
	return xs
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateByOne(_list))
}
