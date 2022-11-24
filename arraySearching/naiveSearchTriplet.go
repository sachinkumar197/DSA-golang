package main

import "fmt"

func searchElem(xs []int, val int) bool {
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			for k := j + 1; k < len(xs); k++ {
				if xs[i]+xs[j]+xs[k] == val {
					fmt.Println(xs[i], xs[j], xs[k])
					return true
				}
			}
		}
	}
	return false
}

func main() {
	A := []int{1, 4, 45, 6, 10, 8}
	val := 22
	searchElem(A, val)
}
