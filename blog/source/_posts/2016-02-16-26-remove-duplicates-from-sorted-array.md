---
title: 26. Remove Duplicates from Sorted Array
date: 2016-02-16 15:49:12
categories:
  - LeetCode
  - Java
tags:
  - Array
  - Two Pointers
---

Total Accepted: **112046**
Total Submissions: **342076**
Difficulty: **Easy**

Given a sorted array, remove the duplicates in place such that each element appear only _once_ and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

For example,
Given input array _nums_ = `[1,1,2]`,

Your function should return length = `2`, with the first two elements of _nums_ being `1` and `1` respectively. It doesn't matter what you leave beyond the new length.

<!-- more -->

Java:

``` java
public class Solution {
    public int removeDuplicates(int[] nums) {
        int dupes = 0;
        int len=nums.length;

        for (int i = 1; i < len; i++)
        {
            if (nums[i] == nums[i - 1])
                dupes++;
            nums[i - dupes] = nums[i];
        }

        return len - dupes;
    }
}
```