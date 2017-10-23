---
title: 326. Power of Three
date: 2016-02-13 15:11:28
categories:
  - LeetCode
  - Java
tags:
  - Math
---

Total Accepted: **18416**
Total Submissions: **52117**
Difficulty: **Easy**

Given an integer, write a function to determine if it is a power of three.

**Follow up:**

Could you do it without using any loop / recursion?

<!-- more -->

Java:

``` java
public class Solution {
    public boolean isPowerOfThree(int n) {            
        return n==Math.pow(3, (int)Math.round(Math.log10(n)/Math.log10(3)));
    }
}
```

更好的当然是这个了

``` java
public class Solution {
    public boolean isPowerOfThree(int n) {
        return n>0&&1162261467%n==0;
    }
}
```