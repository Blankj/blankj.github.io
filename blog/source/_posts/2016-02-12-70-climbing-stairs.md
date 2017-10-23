---
title: 70. Climbing Stairs
date: 2016-02-12 21:40:32
categories:
  - LeetCode
  - Java
tags:
  - Dynamic Programming
---

Total Accepted: **93127**
Total Submissions: **258282**
Difficulty: **Easy**

You are climbing a stair case. It takes _n_ steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

<!-- more -->

Java:

``` java
public class Solution {
    public int climbStairs(int n) {
        int[] ans=new int[n+1];
        ans[0]=ans[1]=1;
        for(int i=2;i<=n;++i)
            ans[i]=ans[i-1]+ans[i-2];
        return ans[n];
    }
}
```