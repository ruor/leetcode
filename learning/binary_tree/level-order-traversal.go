/* 
problems: https://leetcode.com/problems/binary-tree-level-order-traversal/
# 102

explore: https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/931/

Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type LevelNode struct {
    Node *TreeNode
    level int
}

func levelOrder(root *TreeNode) [][]int {
    var res [][]int
    var tra []*LevelNode
    if root == nil {
        return res
    }

    // first level
    var tmp []int
    tmp = append(tmp, root.Val)
    res = append(res, tmp)

    // next level
    level := 1
    curNode := root

    for {
        if level >= len(res) && (curNode.Left != nil || curNode.Right != nil){
            res = append(res, []int{})
        }

        if curNode.Left != nil {
            tra = append(tra, &LevelNode{curNode.Left, level})
            res[level] = append(res[level], curNode.Left.Val)
        }

        if curNode.Right != nil {
            tra = append(tra, &LevelNode{curNode.Right, level})
            res[level] = append(res[level], curNode.Right.Val)
        }

        if len(tra) == 0 {
            break
        } else {
            curNode = tra[0].Node
            level = tra[0].level + 1
            tra = tra[1:len(tra)]
        }
    }

    return res
}
