---
title: 237. Delete Node in a Linked List
date: 2016-02-10 14:52:43
categories:
  - LeetCode
  - C++
tags:
  - Linked List
---

Total Accepted: **61181**
Total Submissions: **139905**
Difficulty: **Easy**

Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.

Supposed the linked list is 1 -> 2 -> 3 -> 4 and you are given the third node with value 3, the linked list should become 1 -> 2 -> 4 after calling your function.

<!-- more -->

标准写当然需要delete了，不标准的话就直接*node=node->next；即可

C++:

``` cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    void deleteNode(ListNode* node) {
        ListNode *next = node->next;
        *node = *next;
        delete next;
    }
};
```