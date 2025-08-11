"Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle."

"In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:"

from typing import List

class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        if rowIndex == 0: return [1]
        triangle = []
        for i in range(0, rowIndex + 1):
            row = [1]
            for j in range(1, i):
                row.append(triangle[i-1][j-1] + triangle[i-1][j])
            row.append(1)
            triangle.append(row)
        return triangle[rowIndex]

# Test
if __name__ == "__main__":
    s = Solution()
    print(s.getRow(5))

