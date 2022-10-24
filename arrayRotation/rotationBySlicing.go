package main

import "fmt"

func rotateByOne(xs []int) []int {
	return append(xs[len(xs)-1:], xs[:len(xs)-1]...)
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateByOne((_list)))
}
