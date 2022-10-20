//https://www.geeksforgeeks.org/array-rotation/
// Problem : Given an array of integers arr[] of size N and an integer, the task is to rotate the array elements to the left by d positions.

// Using Temp array

package main

import "fmt"

func rotateArray(_list []int, d int) []int {
	var tempArray []int
	tempArray = _list[d:]
	tempArray = append(tempArray, _list[:d]...)
	_list = tempArray
	return _list
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateArray(input, 2))
}
