/* Problems: https://leetcode.com/problems/binary-tree-postorder-traversal/
# 145. Binary Tree Postorder Traversal

Explore: https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/930/

Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
Follow up: Recursive solution is trivial, could you do it iteratively?

*/

// Recursive
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var res []int
	if root == nil {
		return res
	}

	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
    res = append(res, root.Val)

    return res
}
