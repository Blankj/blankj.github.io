---
title: Additive equations
date: 2015-01-04 00:54:32
categories:
  - OJ
tags:
  - 深搜
---

### 题目描述

We all understand that an integer set is a collection of distinct integers. Now the question is: given an integer set, can you find all its addtive equations? To explain what an additive equation is, let's look at the following examples:
1+2=3 is an additive equation of the set {1,2,3}, since all the numbers that are summed up in the left-hand-side of the equation, namely 1 and 2, belong to the same set as their sum 3 does. We consider 1+2=3 and 2+1=3 the same equation, and will always output the numbers on the left-hand-side of the equation in ascending order. Therefore in this example, it is claimed that the set {1,2,3} has an unique additive equation 1+2=3.
It is not guaranteed that any integer set has its only additive equation. For example, the set {1,2,5} has no addtive equation and the set {1,2,3,5,6} has more than one additive equations such as 1+2=3, 1+2+3=6, etc. When the number of integers in a set gets large, it will eventually become impossible to find all the additive equations from the top of our minds -- unless you are John von Neumann maybe. So we need you to program the computer to solve this problem.

<!-- more -->
### 输入要求

The input data consists of several test cases.
The first line of the input will contain an integer N, which is the number of test cases.
Each test case will first contain an integer M (1<=M<=30), which is the number of integers in the set, and then is followed by M distinct positive integers in the same line.

### 输出要求

For each test case, you are supposed to output all the additive equations of the set. These equations will be sorted according to their lengths first( i.e, the number of integer being summed), and then the equations with the same length will be sorted according to the numbers from left to right, just like the sample output shows. When there is no such equation, simply output "Can't find any equations." in a line. Print a blank line after each test case.

### 假如输入

<pre>3
3 1 2 3
3 1 2 5
6 1 2 3 5 4 6</pre>

### 应当输出

<pre>1+2=3

Can't find any equations.

1+2=3
1+3=4
1+4=5
1+5=6
2+3=5
2+4=6
1+2+3=6</pre>


题目意思很简单，就是输入一串长度为M的整数，求这组数据中的等式并输出，等式长度从短到长，且等式中自左向右递增。

思路就是先对这组数据进行排序，然后从2——（M-1）长度的等式进行深搜，深搜注意剪枝，否则会超时，没搜到一个等式就存入表中，最后对表进行排序输出。

``` cpp
#include<iostream>
#include<algorithm>
#include <vector>
#include<string.h>
#include<string>
#include<ctype.h>
#include<cmath>
#include <queue>
#define MAXN 30000
using namespace std;
int a[50],flag[50];
int N;
struct s
{
    int a[50];
    int lenth;
} str[MAXN];
int k=0;
void input()
{
    k++;
    str[k].lenth=0;
    int t=1;
    for (int i=1; i<=N; i++)
        if (flag[i]!=0) str[k].a[t++]=a[i],str[k].lenth++;
}
void dfs(int i,int sum)
{                
    sum+=a[i];
    if (sum>a[N]) 
        return ;
    for (int j=i+1; j<=N; j++)
        if (sum==a[j])
        {
            flag[j]=1;
            input();
            flag[j]=0;
            break;
        }
        for (int j=i+1; j<N; j++)
        {
            flag[j]=1;
            dfs(j,sum);
            flag[j]=0;
        }
}
bool cmp(struct s a,struct s b)
{
    if(a.lenth!=b.lenth)
        return a.lenth<b.lenth;
    for(int i=1; i<=a.lenth; i++)
        if (a.a[i]!=b.a[i])
            return a.a[i]<b.a[i];
}
void solve();
int main()
{
    solve();
    return 0;
}
void solve()
{   
    int T;
    cin>>T;
    while (T--)
    {
        memset(flag,0,sizeof(flag));
        memset(a,0,sizeof(a));
        k=0;
        cin>>N;
        for (int i=1; i<=N; i++)
            cin>>a[i];
        sort(a+1,a+1+N);
        for (int i=1; i<N-1; i++)
        {
            flag[i]=1;
            dfs(i,0);
            flag[i]=0;
        }
        sort(str+1,str+k+1,cmp);
        if (k==0)
            printf("Can't find any equations.\n\n");
        else
            for (int i=1; i<=k; i++)
            {
                int j;
                printf("%d",str[i].a[1]);
                for (j=2; j<str[i].lenth; j++)
                    printf("+%d",str[i].a[j]);
                printf("=%d\n",str[i].a[j]);
            }
            printf("\n");
    }
}
```