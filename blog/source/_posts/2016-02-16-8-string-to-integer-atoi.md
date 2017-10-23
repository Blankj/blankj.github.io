---
title: 8. String to Integer (atoi)
date: 2016-02-16 18:16:30
categories:
  - LeetCode
  - Java
tags:
  - Math
  - String
---

Total Accepted: **87527**
Total Submissions: **658165**
Difficulty: **Easy**      

Implement <span style="font-family:monospace">atoi</span> to convert a string to an integer.

**Hint:** Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

**Notes:** It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

**<span style="color:red">Update (2015-02-10):</span>**
The signature of the `C++` function had been updated. If you still see your function signature accepts a `const char *` argument, please click the reload button to reset your code definition.

spoilers alert... click to show requirements for atoi.

**Requirements for atoi:**

<span style="color: rgb(255, 0, 0);">The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.</span>

<span style="color: rgb(255, 0, 0);">The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.</span>

<span style="color: rgb(255, 0, 0);">If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.</span>

<span style="color: rgb(255, 0, 0);">If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.</span>

<!-- more -->

因为没有看到红字部分，都是一直提交来查看错误，哈哈哈，所以略坑

Java:

``` java
public class Solution {
    public int myAtoi(String str) {
        if(str.isEmpty())
            return 0;
        str=str.trim();
        int i=0,ans=0,sign=1,len=str.length();
        if(str.charAt(i)=='-'||str.charAt(i)=='+')
            sign=str.charAt(i++)=='+'?1:-1;
        for(;i<len;++i){
            int tmp=str.charAt(i)-'0';
            if(tmp<0||tmp>9)
                break;
            if(ans>Integer.MAX_VALUE/10||ans==Integer.MAX_VALUE/10&&Integer.MAX_VALUE %10 < tmp)
                return sign==1?Integer.MAX_VALUE:Integer.MIN_VALUE;
            else
                ans=ans*10+tmp;
        }
        return sign*ans;
    }
}
```