---
title: 303. Range Sum Query - Immutable
date: 2016-03-02 22:52:03
categories:
  - LeetCode
  - Java
tags:
  - Dynamic Programming
---

Total Accepted: **20875**
Total Submissions: **85065**
Difficulty: **Easy**

Given an integer array _nums_, find the sum of the elements between indices _i_ and _j_ (_i_ ≤ _j_), inclusive.

**Example:**

<pre>Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3</pre>

**Note:**

1.  You may assume that the array does not change.

2.  There are many calls to _sumRange_ function.

<!-- more -->

Java:

``` java
public class NumArray {

    private static int[] sum;

    public NumArray(int[] nums) {
        for (int i = 1; i < nums.length; ++i)
            nums[i] += nums[i - 1];
        this.sum = nums;
    }

    public int sumRange(int i, int j) {
        return sum[j] - (i == 0 ? 0 : sum[i - 1]);
    }
}
```