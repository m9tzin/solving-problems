# 09 - Palindrome Number problem from Leetcode
# Given an integer x, return true if x is a palindrome, and false otherwise.
# 1221 is a palindrome
# 1212 isnt a palindrome
class Solution:
        def isPalindrome(self, x):
                if x < 0:
                        return False
                num = x
                reversed_num = 0       
                while x > 0:
                        last_digit = x % 10
                        reversed_num = reversed_num * 10 + last_digit
                        x = x // 10
                return num == reversed_num

print("")
solution = Solution() # instance
print(solution.isPalindrome(1221))
        