---
title: 328. Odd Even Linked List
date: 2016-02-10 20:38:41
categories:
  - LeetCode
  - C++
tags:
  - Linked List
---

Total Accepted: **12675**
Total Submissions: **33458**
Difficulty: **Easy**

Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

**Example:**
Given `1->2->3->4->5->NULL`,
return `1->3->5->2->4->NULL`.

**Note:**
The relative order inside both the even and odd groups should remain as it was in the input. 
The first node is considered odd, the second node even and so on ...

<!-- more -->

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
    ListNode* oddEvenList(ListNode* head) {
        if(head==NULL)
            return NULL;
        ListNode *odd=head,*even=head->next,*te=even;
        while(even&&even->next){
            odd->next=even->next;
            odd=odd->next;
            even->next=odd->next;
            even=even->next;
        }
        odd->next=te;
        return head;
    }
};
```