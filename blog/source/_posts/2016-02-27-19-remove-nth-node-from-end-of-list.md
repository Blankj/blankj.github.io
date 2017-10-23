---
title: 19. Remove Nth Node From End of List
date: 2016-02-27 17:48:10
categories:
  - LeetCode
  - Java
tags:
  - Linked List
  - Two Pointers
---

Total Accepted: **96063**
Total Submissions: **333812**
Difficulty: **Easy**

Given a linked list, remove the _n_<sup>th</sup> node from the end of list and return its head.

For example,

<pre>   Given linked list: 1->2->3->4->5, and n = 2.

   After removing the second node from the end, the linked list becomes 1->2->3->5.</pre>

**Note:**
Given _n_ will always be valid.
Try to do this in one pass.

<!-- more -->

Java:

``` java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
public class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode pre = head, after = head;
        while (n-- != 0) {
            after = after.next;
        }
        if (after != null) {
            while (after.next != null) {
                pre = pre.next;
                after = after.next;
            }
            pre.next = pre.next.next;
        } else {
            return head.next;
        }
        return head;
    }
}
```