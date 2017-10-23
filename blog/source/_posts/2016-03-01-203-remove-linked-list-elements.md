---
title: 203. Remove Linked List Elements
date: 2016-03-01 09:38:03
categories:
  - LeetCode
  - Java
tags:
  - Linked List
---

Total Accepted: **53495**
Total Submissions: **191288**
Difficulty: **Easy**

Remove all elements from a linked list of integers that have value **_val_**.

**Example**
_**Given:**_ 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6,  **_val_** = 6
_**Return:**_ 1 --> 2 --> 3 --> 4 --> 5

**Credits:**
Special thanks to @mithmatt for adding this problem and creating all test cases.

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
    public ListNode removeElements(ListNode head, int val) {
        while (head != null && head.val == val)
            head = head.next;
        if (head == null)
            return null;
        ListNode p = head;
        while (p.next != null) {
            if (p.next.val == val) {
                p.next = p.next.next;
            } else {
                p = p.next;
            }
        }
        return head;
    }
}
```