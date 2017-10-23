---
title: 344. Reverse String
date: 2016-07-22 09:54:02
categories:
  - LeetCode
  - JS
tags:
  - Two Pointers
  - String
---

Total Accepted: **71392**
Total Submissions: **188974**
Difficulty: **Easy**

Write a function that takes a string as input and returns the string reversed.

**Example:**

Given s = "hello", return "olleh".

<!-- more -->

js字符串是引用类型，要转为值类型来处理。

JS:
``` js
/**
 * @param {string} s
 * @return {string}
 */
var reverseString = function(s) {
    var arr = s.split('');
    var len = arr.length;
    for(var i = 0;i < len / 2; i++){
        var temp = arr[i];
        arr[i] = arr[len - 1 - i];
        arr[len - i - 1] = temp;
    }
    return arr.join('');
};
```