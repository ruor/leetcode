/*problems: https://leetcode.com/problems/binary-tree-preorder-traversal/
No.144
explore: https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/928/
Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Recursive solution

func preorderTraversal(root *TreeNode) []int {

	var res []int
	if root == nil {
		return res
	}

	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

    return res
}
