---
title: 9. Palindrome Number
date: 2016-02-17 22:44:25
categories:
  - LeetCode
  - Java
tags:
  - Math
---

Total Accepted: **106331**
Total Submissions: **345763**
Difficulty: **Easy**

Determine whether an integer is a palindrome. Do this without extra space.

click to show spoilers.

**Some hints:**

Could negative integers be palindromes? (ie, -1)

If you are thinking of converting the integer to string, note the restriction of using extra space.

You could also try reversing an integer. However, if you have solved the problem "Reverse Integer", you know that the reversed integer might overflow. How would you handle such case?

There is a more generic way of solving this problem.

<!-- more -->

Java:

``` java
public class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0)
            return false;
        int y = 0, x1 = 1;
        while (x / x1 != 0) {
            y = y * 10 + x % 10;
            x /= 10;
            x1 *= 10;
        }
        return y < 10 * x ? y == x : y / 10 == x;
    }
}
```