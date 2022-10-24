// https://www.geeksforgeeks.org/c-program-cyclically-rotate-array-one/

package main

import "fmt"

func rotateByOne(xs []int) []int {
	temp := xs[0]
	for i := 0; i < len(xs)-1; i++ {
		xs[i] = xs[i+1]
	}
	xs[len(xs)-1] = temp
	return xs
}

func main() {
	_list := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateByOne(_list))

}
