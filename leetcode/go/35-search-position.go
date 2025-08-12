// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order. You must write an algorithm with O(log n) runtime complexity.

package main

func searchInsert(nums []int, target int) int {

	// nums = [1,3,5,6]

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return i
		}
	}
	return len(nums)
}
