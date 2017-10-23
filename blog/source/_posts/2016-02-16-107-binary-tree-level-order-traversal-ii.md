---
title: 107. Binary Tree Level Order Traversal II
date: 2016-02-16 13:58:02
categories:
  - LeetCode
  - Java
tags:
  - Tree
  - Breadth-first Search
---

Total Accepted: **70095**
Total Submissions: **211952**
Difficulty: **Easy**

Given a binary tree, return the _bottom-up level order_ traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree `{3,9,20,#,#,15,7}`,

<pre>
    3
   / \
  9  20
    /  \
   15   7
</pre>

return its bottom-up level order traversal as:

<pre>
[
  [15,7],
  [9,20],
  [3]
]
</pre>

confused what `"{1,#,2,3}"` means? > read more on how binary tree is serialized on OJ.

<!-- more -->

Java:

``` java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Solution {
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        List<List<Integer>> res=new LinkedList<List<Integer>>();            
        if(root==null)
            return res;
        Queue<TreeNode> q =new LinkedList<TreeNode>();
        q.offer(root);
        while(!q.isEmpty()){
            int size=q.size();
            List<Integer> tmp=new LinkedList<Integer>();                
            for(int i=0;i<size;++i){
                TreeNode temp=q.poll();
                if(temp.left!=null)
                    q.offer(temp.left);
                if(temp.right!=null)
                    q.offer(temp.right);
                tmp.add(temp.val);
            }
            res.add(0, tmp);
        }
        return res;
    }
}
```