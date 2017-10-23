---
title: 13. Roman to Integer
date: 2016-02-10 22:42:41
categories:
  - LeetCode
  - Java
tags:
  - Math
  - String
---

Total Accepted: **71392**
Total Submissions: **188974**
Difficulty: **Easy**

Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.

<!-- more -->

Java:

``` java
public class Solution {
    public int romanToInt(String s) {
         int len=s.length()-1;
         int[] hash=new int[26];
         hash['I'-'A']=1;
         hash['V'-'A']=5;
         hash['X'-'A']=10;
         hash['L'-'A']=50;
         hash['C'-'A']=100;
         hash['D'-'A']=500;
         hash['M'-'A']=1000;
         int i=0,ans=0;
         for(;i<len;++i){
             if(hash[s.charAt(i)-'A']<hash[s.charAt(i+1)-'A'])
                 ans-=hash[s.charAt(i)-'A'];
             else
                 ans+=hash[s.charAt(i)-'A'];
         }
         return ans+hash[s.charAt(i)-'A'];

    }
}
```