---
title: 257. Binary Tree Paths
date: 2016-03-02 10:02:04
categories:
  - LeetCode
  - Java
tags:
  - Tree
  - Depth-first Search
---

Total Accepted: **36923**
Total Submissions: **135278**
Difficulty: **Easy**

Given a binary tree, return all root-to-leaf paths.

For example, given the following binary tree:

<pre>   1
 /   \
2     3
 \
  5</pre>

All root-to-leaf paths are:
<pre>["1->2->5", "1->3"]</pre>

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
    public void help(List<String> list, TreeNode node, StringBuilder sb) {
        if (node == null)
            return;
        int len=sb.length();
        sb.append(node.val);
        if (node.left == null && node.right == null) {                
            list.add(sb.toString());
            sb.setLength(len);
            return;
        }
        sb.append("->");
        help(list, node.left, sb);
        help(list, node.right, sb);
        sb.setLength(len);
    }

    public List<String> binaryTreePaths(TreeNode root) {
        List<String> res = new ArrayList<String>();
        help(res, root, new StringBuilder());
        return res;
    }
}
```