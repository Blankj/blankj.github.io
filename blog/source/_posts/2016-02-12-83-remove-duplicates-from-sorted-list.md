---
title: 83. Remove Duplicates from Sorted List
date: 2016-02-12 18:01:34
categories:
  - LeetCode
  - Java
tags:
  - Linked List
---

Total Accepted: **100371**
Total Submissions: **278401**
Difficulty: **Easy**

Given a sorted linked list, delete all duplicates such that each element appear only _once_.

For example,
Given `1->1->2`, return `1->2`.
Given `1->1->2->3->3`, return `1->2->3`.

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
    public ListNode deleteDuplicates(ListNode head) {
        if(head == null||head.next==null)
            return head;
        ListNode curr = head;
        ListNode itr = head;
        while(curr.next != null) {
            while(itr.val == curr.val && itr.next != null) {
                itr = itr.next;
            }
            if(curr.val != itr.val) {
                curr.next = itr;
                curr = curr.next;
            }
            else {
                curr.next = null;
            }
        }
        return head;
    }
}
```