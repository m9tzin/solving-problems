# 66 - Plus one
# 20/04/2025
# Matheus Marinho

"You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's. Increment the large integer by one and return the resulting array of digits."

from typing import List

class Solution:
    def plusOne(self, digits):
        number = 0
        for i in range(len(digits)):
            number = number * 10 + digits[i]
        
        number += 1
        num_str = str(number) # 124 -> "124"
        new_array = []
        for c in num_str:
            new_array.append(int(c))
        return new_array
        
# Input: digits = [1,2,3]
# Output: [1,2,4]
if __name__ == "__main__":
    solution = Solution()
    digits = [1,2,3]
    x = solution.plusOne(digits)
    print(f"Output: {x}")
