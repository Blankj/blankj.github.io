---
title: 118. Pascal's Triangle
date: 2016-02-17 14:05:40
categories:
  - LeetCode
  - Java
tags:
  - Array
---

Total Accepted: **73821**
Total Submissions: **227589**
Difficulty: **Easy**

Given _numRows_, generate the first _numRows_ of Pascal's triangle.

For example, given _numRows_ = 5,

Return

<pre>
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
</pre>

<!-- more -->

Java:

``` java
public class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> res = new ArrayList<List<Integer>>();
        List<Integer> row, pre = null;
        for (int i = 0; i < numRows; ++i) {
            row = new ArrayList<Integer>();
            for (int j = 0; j <= i; ++j)
                if (j == 0 || j == i)
                    row.add(1);
                else
                    row.add(pre.get(j - 1) + pre.get(j));
            pre = row;
            res.add(row);
        }
        return res;
    }
}
```