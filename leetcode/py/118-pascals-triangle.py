# 118 - Pascal's Triangle
# 21/04/2025
# Matheus Marinho

"Given an integer numRows, return the first numRows of Pascal's triangle."
"In Pascal's triangle, each number is the sum of the two numbers directly above"

from typing import List

class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        if numRows <=0: return None
        triangle = [[1]]

        for i in range(1, numRows):
            rows = [1]
            
            for j in range(1, i):
                rows.append(triangle[i-1][j-1] + triangle[i-1][j])
            
            rows.append(1)
            triangle.append(rows)
        
        return triangle

# Test
if __name__ == "__main__":
    solution = Solution()
    numRows = 5
    print(solution.generate(numRows))
