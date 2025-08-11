# 26 - Remove duplicates from sorted array
# 19/04/2025
# Matheus Marinho

"Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once."
"The relative order of the elements should be kept the same. Then return the number of unique elements in nums. Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things: Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially."

"The remaining elements of nums are not important as well as the size of nums."

"Return k."

from typing import List
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if nums is None or len(nums) == 0:
            return None
        j = 0
        for i in range(len(nums)):
            if nums[i] != nums[j]:
                j += 1
                nums[j] = nums[i]
        return j+1

# Simple test:
if __name__ == "__main__":
    testring = Solution()
    items = [1,1,1,2,3,4,5,5,6] # k = 6
    print(testring.removeDuplicates(items))
