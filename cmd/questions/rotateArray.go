package main

func RotateArray(arr []int, rotate int) []int {
	resp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		resp[(i+rotate)%len(arr)] = arr[i]
	}
	return resp
}
