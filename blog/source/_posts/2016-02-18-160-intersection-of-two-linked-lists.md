---
title: 160. Intersection of Two Linked Lists
date: 2016-02-18 23:09:49
categories:
  - LeetCode
  - Java
tags:
  - Linked List
---

Total Accepted: **62565**
Total Submissions: **208321**
Difficulty: **Easy**

Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:

<pre>
A:          a1 → a2
                   ↘
                     c1 → c2 → c3
                   ↗            
B:     b1 → b2 → b3
</pre>

begin to intersect at node c1.

**Notes:**

*   If the two linked lists have no intersection at all, return `null`.

*   The linked lists must retain their original structure after the function returns.

*   You may assume there are no cycles anywhere in the entire linked structure.

*   Your code should preferably run in O(n) time and use only O(1) memory.

**Credits:**
Special thanks to <span style="color: rgb(0, 136, 204);">@stellari</span> for adding this problem and creating all test cases.

<!-- more -->

Java:

``` java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        if(headA==null || headB==null)return null;
        ListNode a=headA;
        ListNode b=headB;
        while(a!=b){
            a=a==null?headB:a.next;
            b=b==null?headA:b.next;
        }
        return a;
    }
}
```