---
title: 101. Symmetric Tree
date: 2016-02-13 22:22:23
categories:
  - LeetCode
  - Java
tags:
  - Tree
  - Depth-first Search
  - Breadth-first Search
---

Total Accepted: **93600**
Total Submissions: **281557**
Difficulty: **Easy**

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree is symmetric:

<pre>
    1
   / \
  2   2
 / \ / \
3  4 4  3
</pre>

But the following is not:

<pre>
    1
   / \
  2   2
   \   \
   3    3
</pre>

**Note:**
Bonus points if you could solve it both recursively and iteratively.

confused what `"{1,#,2,3}"` means? > read more on how binary tree is serialized on OJ.

<!-- more -->

Java:

递归:
```java
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Solution {
    public boolean isSymmetric(TreeNode root) {
        if (root == null)
                return true;
        return isSymmetric(root.left ,root.right); 
    }
    public boolean isSymmetric(TreeNode node1 , TreeNode node2){
        if(node1 == null && node2 == null) return true;
        if(node1 == null || node2 == null) return false;

        return node1.val == node2.val
               &&isSymmetric(node1.left, node2.right)
               && isSymmetric(node1.right,node2.left);
    }
}
```

***

单队列：
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
    public boolean isSymmetric(TreeNode root) {
        if (root == null)
                return true;
        Queue<TreeNode> q = new LinkedList<TreeNode>();
        q.offer(root.left);
        q.offer(root.right);
        while (!q.isEmpty()) {
            TreeNode l = q.poll();
            TreeNode r = q.poll();
            if (l == null && r == null)
                continue;
            if (l == null || r == null)
                return false;
            if(l.val!=r.val)
                return false;
            q.offer(l.left);
            q.offer(r.right);
            q.offer(l.right);
            q.offer(r.left);
        }
        return true;
    }
}
```