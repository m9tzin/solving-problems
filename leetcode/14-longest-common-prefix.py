# 14 - Longest common prefix
# 19/04/2025
# Matheus Marinho

"Write a function to find the longest common prefix string amongst an array of strings. If there is no common prefix, return an empty string ""."

from typing import List

class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        # first case: empty
        if not strs:
            return ""
    
        # second case: not empty
        base = strs[0]
        prefix = "" 

        for i in range(len(base)):
            current_char = base[i]

            for s in strs[1:]: # except the first element
                if i >= len(s) or s[i] != current_char:
                    return prefix
            
            prefix += current_char # add the common char

        return prefix


# Simple test:
if __name__ == "__main__":
    testr = Solution()
    items = ["abobora", "abacate", "abracadabra", "abrir"]
    print(testr.longestCommonPrefix(items))
    