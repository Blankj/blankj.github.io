---
title: 278. First Bad Version
date: 2016-02-16 23:17:59
categories:
  - LeetCode
  - Java
tags:
  - Binary Search
---

Total Accepted: **31954**
Total Submissions: **146259**
Difficulty: **Easy**

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have `n` versions `[1, 2, ..., n]` and you want to find out the first bad one, which causes all the following ones to be bad. You are given an API `bool isBadVersion(version)` which will return whether `version` is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

<!-- more -->

Java:

``` java
/* The isBadVersion API is defined in the parent class VersionControl.
      boolean isBadVersion(int version); */

public class Solution extends VersionControl {
    public int firstBadVersion(int n) {
        int l=1,r=n;
        while(l<r){
            int mid=l+((r-l)>>1);
            if (!isBadVersion(mid)) l = mid + 1;
            else r = mid;   
        }
        return l;
    }
}
```