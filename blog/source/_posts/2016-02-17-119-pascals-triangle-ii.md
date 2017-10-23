---
title: 119. Pascal's Triangle II
date: 2016-02-17 16:06:26
categories:
  - LeetCode
  - Java
tags:
  - Array
---

Total Accepted: **66540**
Total Submissions: **211819**
Difficulty: **Easy**

Given an index _k_, return the _k_<sup>th</sup> row of the Pascal's triangle.

For example, given _k_ = 3,
Return `[1,3,3,1]`.

**Note:**
Could you optimize your algorithm to use only _O_(_k_) extra space?

<!-- more -->

Java:

``` java
public class Solution {
    public List<Integer> getRow(int rowIndex) {
        Integer[] arr = new Integer[rowIndex + 1];
        int end = (rowIndex + 2) >> 1;
        arr[0] = arr[rowIndex] = 1;
        long j = rowIndex;
        for (int i = 1; i < end; ++i, --j)
            arr[i] = arr[rowIndex - i] = (int) (j * arr[i - 1] / i);
        return Arrays.asList(arr);
    }
}
```