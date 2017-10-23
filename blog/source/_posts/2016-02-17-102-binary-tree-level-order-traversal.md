---
title: 102. Binary Tree Level Order Traversal
date: 2016-02-17 15:13:43
categories:
  - LeetCode
  - Java
tags:
  - Tree
  - Breadth-first Search
---

Total Accepted: **89586**
Total Submissions: **282417**
Difficulty: **Easy**

Given a binary tree, return the _level order_ traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree `{3,9,20,#,#,15,7}`,

<pre>
    3
   / \
  9  20
    /  \
   15   7
</pre>

return its level order traversal as:

<pre>
[
  [3],
  [9,20],
  [15,7]
]
</pre>

confused what` "{1,#,2,3}"` means? > read more on how binary tree is serialized on OJ.

<!-- more -->

Java:

宽搜:
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
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<List<Integer>>();
        if (root == null)
            return res;
        Queue<TreeNode> q = new LinkedList<TreeNode>();
        q.offer(root);
        while (!q.isEmpty()) {
            List<Integer> tmp = new ArrayList<Integer>();
            int size = q.size();
            for (int i = 0; i < size; ++i) {
                TreeNode t = q.poll();
                if (t.left != null)
                    q.offer(t.left);
                if (t.right != null)
                    q.offer(t.right);
                tmp.add(t.val);
            }
            res.add(tmp);
        }
        return res;
    }
}
```

***

深搜:
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
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<List<Integer>>();
        dfs(res,root,0);
        return res;
    }
    public void dfs(List<List<Integer>> list,TreeNode node,int deep){
        if(node==null)return;
        if(list.size()==deep)
            list.add(new ArrayList<Integer>());
        list.get(deep).add(node.val);
        dfs(list, node.left, deep+1);
        dfs(list, node.right, deep+1);    
    }
}
```