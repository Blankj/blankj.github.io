---
title: 88. Merge Sorted Array
date: 2016-02-21 13:28:32
categories:
  - LeetCode
  - Java
tags:
  - Array
  - Two Pointers
---

Total Accepted: **88608**
Total Submissions: **297379**
Difficulty: **Easy**

Given two sorted integer arrays _nums1_ and _nums2_, merge _nums2_ into _nums1_ as one sorted array.

**Note:**
You may assume that _nums1_ has enough space (size that is greater or equal to _m_ + _n_) to hold additional elements from _nums2_. The number of elements initialized in _nums1_ and _nums2_ are _m_ and _n_ respectively.

<!-- more -->

Java:

``` java
public class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int k = m-- + n-- - 1;
        while (m >= 0 && n >= 0)
            nums1[k--] = nums1[m] > nums2[n] ? nums1[m--] : nums2[n--];
        while (n >= 0)
            nums1[k--] = nums2[n--];
    }
}
```