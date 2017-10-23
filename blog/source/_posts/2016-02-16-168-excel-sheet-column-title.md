---
title: 168. Excel Sheet Column Title
date: 2016-02-16 21:54:59
categories:
  - LeetCode
  - Java
tags:
  - Math
---

Total Accepted: **51862**
Total Submissions: **249733**
Difficulty: **Easy**

Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

<pre>
    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
</pre>

<!-- more -->

Java:

``` java
public class Solution {
    public String convertToTitle(int n) {
        StringBuilder sb = new StringBuilder();
        while(n!=0){    
            sb.insert(0,(char)('A' + --n % 26));
            n/=26;
        }
        return sb.toString();
    }
}
```