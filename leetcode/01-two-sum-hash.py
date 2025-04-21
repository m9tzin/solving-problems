# 01 - Two Sum - Hash approach
# 21/04/2025
# Matheus Marinho

"Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target."
"You may assume that each input would have exactly one solution, and you may not use the same element twice. You can return the answer in any order."

from typing import List

class Solution(object):
    def twoSum(self, nums, target):
        hash = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in hash:
                return [i, hash[complement]]
            hash[nums[i]] = i
        return None
    
if __name__ == "__main__":
    solution = Solution()
    nums = [1, 5, 9, 12]
    target = 17
    x = solution.twoSum(nums, target)
    print(f"Output: {x}")    