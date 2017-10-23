---
title: 14. Longest Common Prefix
date: 2016-03-01 19:25:56
categories:
  - LeetCode
  - Java
tags:
  - String
---

Total Accepted: **88409**
Total Submissions: **321579**
Difficulty: **Easy**

Write a function to find the longest common prefix string amongst an array of strings.

<!-- more -->

Java:

``` java
public class Solution {
    public String longestCommonPrefix(String[] strs) {
        int len = strs.length;
        if (len == 0)
            return "";
        int minlen = 0x7fffffff;
        for (int i = 0; i < len; ++i) 
            minlen = Math.min(minlen, strs[i].length());
        for (int j = 0; j < minlen; ++j) 
            for (int i = 1; i < len; ++i) 
                if (strs[0].charAt(j) != strs[i].charAt(j)) 
                    return strs[0].substring(0, j);
        return strs[0].substring(0, minlen);
    }
}
```