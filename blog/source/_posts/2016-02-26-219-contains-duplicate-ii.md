---
title: 219. Contains Duplicate II
date: 2016-02-26 16:52:07
categories:
  - LeetCode
  - Java
tags:
  - Array
  - Hash Table
---

Total Accepted: **46925**
Total Submissions: **160698**
Difficulty: **Easy**

Given an array of integers and an integer _k_, find out whether there are two distinct indices _i_ and _j_ in the array such that **nums[i] = nums[j]** and the difference between _i_ and _j_ is at most _k_.

<!-- more -->

Java:

``` java
public class Solution {
    public class MyArr {
        int data;
        int index;
    }

    class mycmp implements Comparator<MyArr> {
        @Override
        public int compare(MyArr o1, MyArr o2) {
            return o1.data - o2.data;
        }
    }

    public boolean containsNearbyDuplicate(int[] nums, int k) {
        int len = nums.length;
        MyArr[] myArrs = new MyArr[len];
        for (int i = 0; i < len; ++i) {
            myArrs[i] = new MyArr();
            myArrs[i].data = nums[i];
            myArrs[i].index = i;
        }
        Arrays.sort(myArrs, 0, len, new mycmp());
        for (int i = 0; i < len - 1; i++) {
            if (myArrs[i].data == myArrs[i + 1].data
                    && Math.abs((myArrs[i].index - myArrs[i + 1].index)) <= k)
                return true;
        }
        return false;
    }
}
```