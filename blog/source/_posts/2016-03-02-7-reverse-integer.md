---
title: 7. Reverse Integer
date: 2016-03-02 23:10:46
categories:
  - LeetCode
  - Java
tags:
  - Math
---

Total Accepted: **125501**
Total Submissions: **532378**
Difficulty: **Easy**

Reverse digits of an integer.

**Example1:** x =  123, return  321
**Example2:** x = -123, return -321

click to show spoilers.

**Have you thought about this?**

Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!

If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.

Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer, then the reverse of 1000000003 overflows. How should you handle such cases?

For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

<!-- more -->

Java:

``` java
public class Solution {
    public int reverse(int x) {
        long res = 0;
        for (; x != 0; x /= 10)
            res = res * 10 + x % 10;
        return res > Integer.MAX_VALUE || res < Integer.MIN_VALUE ? 0: (int) res;
    }
}
```