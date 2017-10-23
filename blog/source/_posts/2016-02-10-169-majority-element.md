---
title: 169. Majority Element
date: 2016-02-10 19:32:09
categories:
  - LeetCode
  - C++
tags:
  - Array
  - Divide and Conquer
  - Bit Manipulation
---

Total Accepted: **92197**
Total Submissions: **232301**
Difficulty: **Easy**

Given an array of size _n_, find the majority element. The majority element is the element that appears more than `⌊ n/2 ⌋` times.

You may assume that the array is non-empty and the majority element always exist in the array.

<!-- more -->

还是hash，hash中过半即答案

C++:

``` cpp
class Solution {
public:
    int majorityElement(vector<int>& nums) {
        unordered_map<int,int> h;
        for(int i=0,n=nums.size();i<n;++i)
            if(++h[nums[i]]>n/2)return nums[i];
    }
};
```