---
title: 172. Factorial Trailing Zeroes
date: 2016-02-17 14:26:04
categories:
  - LeetCode
  - Java
tags:
  - Math
---

Total Accepted: **50404**
Total Submissions: **158693**
Difficulty: **Easy**

Given an integer _n_, return the number of trailing zeroes in _n_!.

**Note:** Your solution should be in logarithmic time complexity.

**Credits:**
Special thanks to @ts for adding this problem and creating all test cases.

<!-- more -->

Java

``` java
public class Solution {
    public int trailingZeroes(int n) {
        int res = 0;
        while(n > 4)
        {
            res += n / 5;
            n /= 5;
        }
        return res;
    }
}
```