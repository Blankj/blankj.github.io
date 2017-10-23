---
title: QuickSort（快速排序）
date: 2016-02-04 22:08:28
id: 252
categories:
  - 数据结构
tags:
  - 排序
---

``` cpp
#include <cstdio>
#include <iostream>
using namespace std;
void quick_sort(int dat[], int l, int r) {
    // 首先请填写下面三个变量的初值
    int i = l, j =r , mid = dat[r];
    do {
        while (dat[i] < mid) ++i;
        while (dat[j] > mid) --j;
        if (i <= j) {
            swap(dat[i], dat[j]);
            ++i; --j;
        }
    } while (i < j);
    // 接下来请填写第一个递归调用的参数，仔细回顾一下刚刚讲的快速排序算法的思想哈。
    if (l < j) quick_sort(dat,l,j);
    // 接下来请填写第二个递归调用的参数。
    if (i < r) quick_sort(dat,i,r);
}
int main() {
    int dat[10] = {1, 4, 3, 2, 5, 3, 2, 5, 10, 9};
    quick_sort(dat, 0, 9);
    for (int i = 0; i < 10; ++i)
        printf("%d ", dat[i]);
    return 0;
}
```