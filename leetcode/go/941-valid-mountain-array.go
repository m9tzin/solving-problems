package main

import "fmt"

func validMountainArray(arr []int) bool {
	n := len(arr)
	if len(arr) < 3 {
		return false
	}

	// up case:
	idx := 0
	// up case:
	for idx+1 < n && arr[idx] < arr[idx+1] {
		idx += 1
	}
	// peak case:
	if idx == 0 || idx == n-1 {
		return false
	}
	// down case:
	for idx+1 < n && arr[idx] > arr[idx+1] {
		idx += 1
	}
	return idx == n-1
}

func main() {
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}
