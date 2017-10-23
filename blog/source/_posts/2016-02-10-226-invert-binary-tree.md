---
title: 226. Invert Binary Tree
date: 2016-02-10 15:12:44
categories:
  - LeetCode
  - C++
tags:
  - Tree
---

Total Accepted: **67715**
Total Submissions: **155630**
Difficulty: **Easy**

Invert a binary tree.

<pre>     4
   /   \
  2     7
 / \   / \
1   3 6   9</pre>

to

<pre>     4
   /   \
  7     2
 / \   / \
9   6 3   1</pre>

<!-- more -->

一种是用递归，一种是用栈，代码如下

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
    TreeNode* invertTree(TreeNode* root) {
        /*递归
        if (root == NULL) {
            return NULL;
        }

        TreeNode* left = root->left;
        TreeNode* right = root->right;
        root->left = invertTree(right);
        root->right = invertTree(left);
        return root;
        /*
        if(root==NULL)
            return root;
        queue<TreeNode *>q;
        q.push(root);
        while(!q.empty()){
            auto node=q.front();
            q.pop();
            auto temp=node->left;
            node->left=node->right;
            node->right=temp;
            if(node->left!=NULL)
                q.push(node->left);
            if(node->right!=NULL)
                q.push(node->right);
        }
        return root;
    }
};
```