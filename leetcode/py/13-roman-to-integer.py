# 13 - Roman to Integer problem from Leetcode
# 17/04/2025
# Matheus Marinho

"Convert a Roman numeral string into an integer, following the standard rules of Roman numeral representation, including subtraction for specific cases (e.g., IV = 4)."

class Solution:
    def romanToInt(self, s:str) -> int:
        romans = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000
        }
        total = 0
        prev_value = 0

        for char in reversed(s):
            value = romans[char]
            if value < prev_value:
                total -= value
            else:
                total += value
            prev_value = value
        return total