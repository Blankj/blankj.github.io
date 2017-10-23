---
title: 217. Contains Duplicate
date: 2016-02-10 18:46:02
categories:
  - LeetCode
  - C++
tags:
  - Array
  - Hash Table
---

Total Accepted: **67962**
Total Submissions: **169476**
Difficulty: **Easy**

Given an array of integers, find if the array contains any duplicates. Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

<!-- more -->

C++:

``` cpp
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        int n=nums.size();
        unordered_map<int,int> h;
        for(int i=0;i<n;++i){
            if(h.count(nums[i])){
                return true;
            }
            h[nums[i]]=1;
        }
        return false;
    }
};
```

看到了使用set的，表示佩服，代码只需一行

``` cpp
class Solution {  public:    bool containsDuplicate(vector<int>& nums) { 
       return set<int>(nums.begin(), nums.end()).size() < nums.size();
    }
};
```