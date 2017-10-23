---
title: 28. Implement strStr()
date: 2016-02-25 21:32:13
categories:
  - LeetCode
  - Java
tags:
  - Two Pointers
  - String
---

Total Accepted: **93775**
Total Submissions: **386039**
Difficulty: **Easy**

Implement strStr().

Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

<!-- more -->

Java:

``` java
public class Solution {
    public int strStr(String haystack, String needle) {
        int len1 = haystack.length();
        int len2 = needle.length();
        if (len2 == 0)
            return 0;
        int[] next = new int[len2];
        char[] substr = needle.toCharArray();
        char[] str = haystack.toCharArray();
        getNext(substr, next);
        int i = 0, j = 0;
        while (i < len1 && j < len2) {
            if (j == -1 || str[i] == substr[j]) {
                ++i;
                ++j;
            } else {
                j = next[j];
            }
        }
        if (len2 == j)
            return i - j;
        return -1;
    }

    private void getNext(char str[], int next[]) {
        int i = 0, j = -1;
        next[0] = -1;
        int len = str.length;
        while (i < len - 1) {
            if (j == -1 || str[i] == str[j]) {
                ++i;
                ++j;
                if (str[i] == str[j])
                    next[i] = next[j];
                else
                    next[i] = j;

            } else
                j = next[j];
        }
    }
}
```