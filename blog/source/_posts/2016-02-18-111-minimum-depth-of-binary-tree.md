---
title: 111. Minimum Depth of Binary Tree
date: 2016-02-18 16:52:03
categories:
  - LeetCode
  - Java
tags:
  - Tree
  - Depth-first Search 
  - Breadth-first Search
---

Total Accepted: **93588**
Total Submissions: **309967**
Difficulty: **Easy**

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

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
    public int minDepth(TreeNode root) {
        if(root==null)return 0;
        Queue<TreeNode> comeIn=new LinkedList<TreeNode>();
        int res=0;
        comeIn.offer(root);
        while(!comeIn.isEmpty()){
            int size=comeIn.size();
            res++;
            for(int i = 0;i < size; ++ i){
                TreeNode tmp=comeIn.poll();
                if(tmp.left == null && tmp.right == null)
                    return res;
                if(tmp.left != null)
                    comeIn.offer(tmp.left);
                if(tmp.right != null)
                    comeIn.offer(tmp.right);    
            }
        }
        return 520;
    }
}
```