---
title: 110. Balanced Binary Tree
date: 2016-02-13 21:57:09
categories:
  - LeetCode
  - Java
tags:
  - Tree 
  - Depth-first Search
---

Total Accepted: **96062**
Total Submissions: **287739**
Difficulty: **Easy**

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of _every_ node never differ by more than 1.

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
    public boolean isBalanced(TreeNode root) {
        return check(root)!=-1;
    }

    public int check(TreeNode root){
        if(root==null)
            return 0;
        int lh=check(root.left);
        if(lh==-1)
            return -1;
        int rh=check(root.right);
        if(rh==-1)
            return -1;
        if(Math.abs(lh-rh)>1)
            return -1;
        return 1+Math.max(lh, rh);
    }
}
```