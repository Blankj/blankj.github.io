---
title: 21. Merge Two Sorted Lists
date: 2016-02-13 16:40:38
categories:
  - LeetCode
  - Java
tags:
  - Linked List
---

Total Accepted: **107818**
Total Submissions: **312053**
Difficulty: **Easy**

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

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
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode ans=new ListNode(0);
        ListNode temp=ans;
        while(l1!=null&&l2!=null){
            if(l1.val<l2.val){
                temp.next=l1;
                l1=l1.next;
            }else{
                temp.next=l2;
                l2=l2.next;
            }
            temp=temp.next;
        }
        if(l1!=null)
            temp.next=l1;
        else
            temp.next=l2;
        return ans.next;
    }
}
```