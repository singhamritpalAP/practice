package main

import (
	"awesomeProject/solutions"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	nums := []int{3, 4, 5, 1, 2}
	arr2 := []int{1, 2, 3, 4, 4, 4, 7, 8, 9}
	mountainArray := []int{0, 10, 5, 2}
	duplicateArray := []int{4, 3, 2, 7, 8, 2, 3, 1}
	target := 9
	rotate := 4

	fmt.Println("Search result: ", solutions.Search(arr, target, 0, len(arr)-1))

	fmt.Println("original array: ", arr, "\nrotated array: ", solutions.RotateArray(arr, rotate))

	fmt.Println("Is array rotated and sorted: ", solutions.IsRotateSorted(nums))

	fmt.Printf("The next greatest letter after '%c' is '%c'\n", byte('c'), solutions.NextGreatestLetter([]byte{'c', 'j', 'j'}, byte('c')))

	fmt.Println("Search result: ", solutions.SearchRange(arr2, 4))

	fmt.Println("Peak index in mountain array is: ", solutions.PeakIndexInMountainArray(mountainArray))

	fmt.Println("Duplicate elements are: ", solutions.FindDuplicates(duplicateArray))
}
