---
title: 234. Palindrome Linked List
date: 2016-03-02 13:45:21
categories:
  - LeetCode
  - Java
tags:
  - Linked List
  - Two Pointers
---

Total Accepted: **39482**
Total Submissions: **145556**
Difficulty: **Easy**

Given a singly linked list, determine if it is a palindrome.

**Follow up:**
Could you do it in O(n) time and O(1) space?

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
    public boolean isPalindrome(ListNode head) {
        int len = 0;
        ListNode p = head, tmp, newHead = null;
        while (p != null) {
            p = p.next;
            ++len;
        }
        p = head;
        int halfLen = len >>> 1;
        for (int i = 0; i < halfLen; ++i) {
            tmp = p.next;
            p.next = newHead;
            newHead = p;
            p = tmp;
        }
        if (len % 2 == 1) {
            p = p.next;
        }
        for (int i = 0; i < halfLen; ++i) {
            if (newHead.val != p.val)
                return false;
            newHead = newHead.next;
            p = p.next;
        }
        return true;
    }
}
```