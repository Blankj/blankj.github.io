---
title: 165. Compare Version Numbers
date: 2016-02-16 18:54:07
categories:
  - LeetCode
  - Java
tags:
  - String
---

Total Accepted: **45172**
Total Submissions: **268283**
Difficulty: **Easy**

Compare two version numbers _version1_ and _version2_.
If _version1_ > _version2_ return 1, if _version1_ < _version2_ return -1, otherwise return 0.

You may assume that the version strings are non-empty and contain only digits and the `.` character.

The `.` character does not represent a decimal point and is used to separate number sequences.

For instance, `2.5` is not "two and a half" or "half way to version three", it is the fifth second-level revision of the second first-level revision.

Here is an example of version numbers ordering:

<pre>0.1 < 1.1 < 1.2 < 13.37</pre>

<!-- more -->

Java:

``` java
public class Solution {
    public int compareVersion(String version1, String version2) {
        int len1=version1.length(),len2=version2.length(),i1=-1,i2=-1;
        while(i1<len1||i2<len2){
            ++i1;++i2;
            int v1=0,v2=0;
            while(i1<len1&&version1.charAt(i1)!='.')
                v1=v1*10+version1.charAt(i1++)-'0';
            while(i2<len2&&version2.charAt(i2)!='.')
                v2=v2*10+version2.charAt(i2++)-'0';
            if(v1!=v2)
                return v1>v2?1:-1;
        }
        return 0;
    }
}
```