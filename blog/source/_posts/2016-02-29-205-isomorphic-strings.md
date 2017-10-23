---
title: 205. Isomorphic Strings
date: 2016-02-29 12:15:39
categories:
  - LeetCode
  - Java
tags:
  - Hash Table
---

Total Accepted: **47679**
Total Submissions: **165936**
Difficulty: **Easy**

Given two strings **_s_** and **_t_**, determine if they are isomorphic.

Two strings are isomorphic if the characters in **_s_** can be replaced to get **_t_**.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

For example,
Given `"egg"`, `"add"`, return true.

Given `"foo"`, `"bar"`, return false.

Given `"paper"`, `"title"`, return true.

**Note:**
You may assume both **_s_** and **_t_** have the same length.

<!-- more -->

Java:

``` java
public class Solution {
    public boolean isIsomorphic(String s, String t) {
        int[] times1 = new int[128], times2 = new int[128];
        int len = s.length();
        for (int i = 0, smax = 0, tmax = 0; i < len; ++i) {
            char sc = s.charAt(i);
            char tc = t.charAt(i);
            if (times1[sc] == 0) {
                times1[sc] = ++smax;
            }
            if (times2[tc] == 0) {
                times2[tc] = ++tmax;
            }
            if (times1[sc] != times2[tc])
                return false;
        }
        return true;
    }
}
```