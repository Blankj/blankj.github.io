---
title: 206. Reverse Linked List
date: 2016-02-10 21:21:54
categories:
  - LeetCode
  - Java
tags:
  - Linked List
---

Total Accepted: **81448**
Total Submissions: **214931**
Difficulty: **Easy**

Reverse a singly linked list.

click to show more hints.

**Hint:**

A linked list can be reversed either iteratively or recursively. Could you implement both?

<!-- more -->

发现java大部分是最快的，难道是easy题目的原因？

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
    public ListNode reverseList(ListNode head) {
        ListNode re=null,t;
        while(head!=null){
            t=head.next;
            head.next=re;
            re=head;
            head=t;
        }
        return re;
    }
}
```