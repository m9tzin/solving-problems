# 226 - Invert Binary Tree
# 21/04/2025
# Matheus Marinho

"Given the root of a binary tree, invert the tree, and return its root."

from typing import Optional
#Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
    
    def __repr__(self):
        return f"TreeNode({self.val}, {repr(self.left)}, {repr(self.right)})"

class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        def invert(root):
            if not root: return None
            if root:
                root.left, root.right = root.right, root.left
                invert(root.left)
                invert(root.right)
        
        invert(root)
        
        return root

# Test
if __name__ == "__main__":
    solution = Solution()
    root = TreeNode(4,
        TreeNode(2, TreeNode(1), TreeNode(3)),
        TreeNode(7, TreeNode(6), TreeNode(9))
    )
    invert = solution.invertTree(root)
    print("Inorder da árvore invertida:", invert)
