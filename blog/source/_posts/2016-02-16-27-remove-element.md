---
title: 27. Remove Element
date: 2016-02-16 15:10:14
categories:
  - LeetCode
  - Java
tags:
  - Array
  - Two Pointers
---

Total Accepted: **101318**
Total Submissions: **307664**
Difficulty: **Easy**

Given an array and a value, remove all instances of that value in place and return the new length.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

<!-- more -->

Java:

``` java
public class Solution {
    public int removeElement(int[] nums, int val) {
        int len=nums.length;
        int res=0;
        for(int i=0;i<len;++i)
            if(nums[i]!=val)
                nums[res++]=nums[i];
        return res;
    }
}
```