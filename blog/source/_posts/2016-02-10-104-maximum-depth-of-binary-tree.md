---
title: 104. Maximum Depth of Binary Tree
date: 2016-02-10 14:42:42
categories:
  - LeetCode
  - C++
tags:
  - Tree
  - Depth-first Search
---

Total Accepted: **120354**
Total Submissions: **255664**
Difficulty: **Easy**

Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

<!-- more -->

一种是递归后序遍历树，返回深度，一种是层次遍历，有几层便是深度。

C++:

``` cpp
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
    int maxDepth(TreeNode* root) {
        /*
        if(root==NULL)
            return 0;
        int rd=maxDepth(root->right);
        int ld=maxDepth(root->left);
        return 1+max(rd,ld);
        */
        if(root == NULL)
        return 0;

        int res = 0;
        queue<TreeNode *> q;
        q.push(root);
        while(!q.empty())
        {
            ++ res;
            for(int i = 0, n = q.size(); i < n; ++ i)
            {
                TreeNode *p = q.front();
                q.pop();

                if(p -> left != NULL)
                    q.push(p -> left);
                if(p -> right != NULL)
                    q.push(p -> right);
            }
        }

        return res;
    }
};
```