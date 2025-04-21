# 704 - Binary Search
# 21/04/2025
# Matheus Marinho

"Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1."
"You must write an algorithm with O(log n) runtime complexity."

from typing import List

class Solution(object):
    def search(self, nums, target):
        left, right = 0, len(nums)-1
        while left <= right:
            mid = (left+right) // 2
            if nums[mid] < target:
                left = mid + 1 
            elif nums[mid] > target:
                right = mid - 1
            else: 
                return mid
                
        return -1


# Output
if __name__ == "__main__":
    solution = Solution()
    nums = [-1,0,3,5,9,12]
    target = 9
    x = solution.search(nums, target)
    print(f"Output: index = [{x}]")