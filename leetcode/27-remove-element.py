# 27 - Remove element
# 20/04/2025
# Matheus Marinho

"Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val."

"Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:"
"Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums."

"Return k."

from typing import List

class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
            j = 0
            for i in range(len(nums)):
                if nums[i] != val:
                     nums[j] = nums[i]
                     j += 1
            return j
    
# Input: nums = [3,2,2,3], val = 3
# Output: 2, nums = [2,2]

# Output
if __name__ == "__main__":
    nums = [3, 2, 2, 3]
    val = 3
    solution = Solution()
    k = solution.removeElement(nums, val)
    print(f"Output: {k}, nums = {nums[:k]}")