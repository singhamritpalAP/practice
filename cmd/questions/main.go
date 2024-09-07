package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 9
	rotate := 4
	result := Search(arr, target, 0, len(arr)-1)
	fmt.Println(result)
	rotatedArr := RotateArray(arr, rotate)
	fmt.Println(rotatedArr)
}
