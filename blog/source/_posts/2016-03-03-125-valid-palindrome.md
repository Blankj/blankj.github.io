---
title: 125. Valid Palindrome
date: 2016-03-03 23:56:23
categories:
  - LeetCode
  - Java
tags:
  - Two Pointers
  - String
---

Total Accepted: **91860**
Total Submissions: **392114**
Difficulty: **Easy**

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

For example, `"A man, a plan, a canal: Panama"` is a palindrome.

`"race a car"` is _not_ a palindrome.

**Note:**
Have you consider that the string might be empty? This is a good question to ask during an interview.

For the purpose of this problem, we define empty string as valid palindrome.

<!-- more -->

Java:

``` java
public class Solution {
    public boolean isPalindrome(String s) {
        if ("".equals(s))
            return true;
        for (int i = 0, j = s.length() - 1; i < j; ++i, --j) {
            char ci = s.charAt(i);
            while (i < j
                    && !(ci >= 'a' && ci <= 'z' || ci >= 'A' && ci <= 'Z' || ci >= '0' && ci <= '9'))
                ci = s.charAt(++i);

            char cj = s.charAt(j);
            while (i < j
                    && !(cj >= 'a' && cj <= 'z' || cj >= 'A' && cj <= 'Z' || cj >= '0' && cj <= '9'))
                cj = s.charAt(--j);

            if (i < j) {
                if (ci >= 'A' && ci <= 'Z')
                    ci += 32;
                if (cj >= 'A' && cj <= 'Z')
                    cj += 32;
                if (ci != cj)
                    return false;
            }
        }
        return true;
    }
}
```