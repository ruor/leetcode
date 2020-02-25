/* Problems: https://leetcode.com/problems/binary-tree-inorder-traversal/
# 94

Explore: https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/929/

Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?

*/

// Recursively:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, inorderTraversal(root.Left)...)
    res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

    return res
}


