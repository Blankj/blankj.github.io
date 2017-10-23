---
title: 242. Valid Anagram
date: 2016-02-10 17:14:31
categories:
  - LeetCode
  - C++
tags:
  - Hash Table 
  - Sort
---

Total Accepted: **58209**
Total Submissions: **143425**
Difficulty: **Easy**

Given two strings _s_ and _t_, write a function to determine if _t_ is an anagram of _s_.

For example,
_s_ = "anagram", _t_ = "nagaram", return true.
_s_ = "rat", _t_ = "car", return false.

**Note:**
You may assume the string contains only lowercase alphabets.

**Follow up:**
What if the inputs contain unicode characters? How would you adapt your solution to such case?

<!-- more -->

求两个字符串中字母出现频率是否一致，hash最快了当然

C++:

``` cpp
class Solution {
public:
    bool isAnagram(string s, string t) {
        int l1=s.length();
        int l2=t.length();
        if(l1!=l2)
            return false;
        int hash[26]={0};
        for(int i=0;i<l1;++i){
            ++hash[s[i]-'a'];
            --hash[t[i]-'a'];
        }
        for(int i=0;i<26;++i){
            if(hash[i])
                return false;
        }
        return true;
    }
};
```