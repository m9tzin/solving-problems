//go:build bubble
// +build bubble

package main

import "fmt"

func bubbleSort(arr []int32) {
	// Bubble sort algorithm
	for range arr {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func main() {
	arr := []int32{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr)
	fmt.Println(arr)
}
