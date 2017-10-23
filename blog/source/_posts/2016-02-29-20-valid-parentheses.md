---
title: 20. Valid Parentheses
date: 2016-02-29 12:49:43
categories:
  - LeetCode
  - Java
tags:
  - Stack
  - String
---

Total Accepted: **95596**
Total Submissions: **332615**
Difficulty: **Easy**

Given a string containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

The brackets must close in the correct order, `"()"` and `"()[]{}"` are all valid but `"(]"` and `"([)]"` are not.

<!-- more -->

Java:

``` java
public class Solution {
    public boolean isValid(String s) {
        int len = s.length();
        char[] stack = new char[len + 1];
        int top = 1;
        for (int i = 0; i < len; ++i) {
            char c = s.charAt(i);
            if (c == '(' || c == '[' || c == '{')
                stack[top++] = c;
            else if (c == ')' && stack[top - 1] != '(')
                return false;
            else if (c == ']' && stack[top - 1] != '[')
                return false;
            else if (c == '}' && stack[top - 1] != '{')
                return false;
            else
                --top;
        }
        if (top == 1)
            return true;
        return false;
    }
}
```