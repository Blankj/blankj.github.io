---
title: 171. Excel Sheet Column Number
date: 2016-02-10 19:07:45
categories:
  - LeetCode
  - C++
tags:
  - Math
---

Total Accepted: **64100**
Total Submissions: **159751**
Difficulty: **Easy**

Related to question Excel Sheet Column Title

Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

<pre>
    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
</pre>

<!-- more -->

我是把它看成是类似是26进制

C++:

``` cpp
class Solution {
public:
    int titleToNumber(string s) {
        int ans = 0;
        for(int i = 0; i < s.length(); ++i)
            ans = ans * 26 + s[i] - 'A' + 1;
        return ans;
    }
};
```