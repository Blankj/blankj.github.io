---
title: A sequence of numbers（快速求幂）
date: 2015-01-12 00:59:44
categories:
  - OJ
tags:
  - 数学
---

### 题目描述

Xinlv wrote some sequences on the paper a long time ago, they might be arithmetic or geometric sequences. The numbers are not very clear now, and only the first three numbers of each sequence are recognizable. Xinlv wants to know some numbers in these sequences, and he needs your help.

<!-- more -->
### 输入要求

The first line contains an integer N, indicting that there are N sequences. Each of the following N lines contain four integers. The first three indicating the first three numbers of the sequence, and the last one is K, indicating that we want to know the K-th numbers of the sequence.
You can assume 0 < K <= 10^9, and the other three numbers are in the range [0, 2^63). All the numbers of the sequences are integers. And the sequences are non-decreasing.

### 输出要求

Output one line for each test case, that is, the K-th number module (%) 200907.

### 假如输入

<pre>2
1 2 3 5
1 2 4 5</pre>

### 应当输出

<pre>5
16</pre>


题目大意很简单，就是给你一个序列的前三项，该序列不是等差就是等比，让你求第K项余200907，求等比时用到快速求幂，不用的话应该会TLE，下面介绍快速求幂：

快速求幂实现代码为

``` c
int fastpow(int a,int b) 
{
    int r=1,base=a; 
    while(b!=0)
    { 
        if(b&1) r*=base;
        base*=base;
        b>>=1;
    } 
    return r;
}
```

``` cpp
#include<iostream>
#include<algorithm>
#include <vector>
#include<string.h>
#include<string>
#include <cstring>
#include<ctype.h>
#include<math.h>
#include <queue>
#include<map>
using namespace std;
int MOD=200907;
long long fastpow(long long q,long long n)
{
    long long r=1;
    while(n)
    {
        if(n&amp;1)
            r=r%MOD*(q%MOD)%MOD;
        q=q%MOD*(q%MOD)%MOD;
        n>>=1;
    }
    return r;
}
void solve();
int main()
{
    solve();
    return 0;
}
void solve()
{   
    long long t,a1,a2,a3,n,d,ans,q;
    cin>>t;
    while(t--)
    {
        cin>>a1>>a2>>a3>>n;
        if(a2*2==a1+a3)
        {
            d=a2-a1;
            ans=(a1+(n-1)*d%MOD)%MOD;
        }
        else
        {
            q=a2/a1;
            ans=a1*fastpow(q,n-1)%MOD;
        }
        cout<<ans<<endl;
    }
}
```