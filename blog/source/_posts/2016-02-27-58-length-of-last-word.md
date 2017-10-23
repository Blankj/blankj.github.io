---
title: 58. Length of Last Word
date: 2016-02-27 17:27:40
categories:
  - LeetCode
  - Java
tags:
  - String
---

Total Accepted: **83878**
Total Submissions: **291685**
Difficulty: **Easy**

Given a string _s_ consists of upper/lower-case alphabets and empty space characters `' '`, return the length of last word in the string.

If the last word does not exist, return 0.

**Note:** A word is defined as a character sequence consists of non-space characters only.

For example,
Given _s_ = `"Hello World"`, return `5`.

<!-- more -->

Java:

``` java
public class Solution {
    public int lengthOfLastWord(String s) {
        s = s.trim();
        if (s.equals(""))
            return 0;
        int res = 0;
        for (int i = s.length() - 1; i > -1 && s.charAt(i) != ' '; ++res, --i);
        return res;
    }
}
```