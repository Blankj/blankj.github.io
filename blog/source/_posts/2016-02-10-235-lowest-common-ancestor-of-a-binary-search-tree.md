---
title: 235. Lowest Common Ancestor of a Binary Search Tree
date: 2016-02-10 20:00:44
categories:
  - LeetCode
  - C++
tags:
  - Tree
---

Total Accepted: **52403**
Total Submissions: **137993**
Difficulty: **Easy**

Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia:
 “The lowest common ancestor is defined between two nodes v and w as the
 lowest node in T that has both v and w as descendants (where we allow **a node to be a descendant of itself**).”</p>

<pre>        _______6______
       /              \
    ___2__          ___8__
   /      \        /      \
   0      _4       7       9
         /  \
         3   5</pre>

For example, the lowest common ancestor (LCA) of nodes `2` and `8` is `6`. Another example is LCA of nodes `2` and `4` is `2`, since a node can be a descendant of itself according to the LCA definition.

<!-- more -->

二叉搜索树的找到两个节点的最近双亲

C++:

``` cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        while((root->val-p->val)*(root->val-q->val)>0)
            root=p->val<root->val?root->left:root->right;
        return root;
    }
};
```