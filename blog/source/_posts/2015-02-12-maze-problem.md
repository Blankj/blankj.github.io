---
title: Maze Problem
date: 2015-02-12 18:09:35
categories:
  - OJ
tags:
  - 广搜
---

### 题目描述

Given a maze, find a shortest path from start to goal.

<!--more-->
### 输入要求

Input consists serveral test cases.

First line of the input contains number of test case T.

For each test case the first line contains two integers N , M ( 1 <= N, M <= 100 ).

Each of the following N lines contain M characters. Each character means a cell of the map.

Here is the definition for chracter.


Constraint:

* For a character in the map:

 * 'S' : start cell

 * 'E' : goal cell

 * '-' : empty cell

 * '#' : obstacle cell

* no two start cell exists.

* no two goal cell exists.

### 输出要求

For each test case print one line containing shortest path. If there exists no path from start to goal, print -1.

### 假如输入

<pre>1
5 5
S-###
-----
##---
E#---
---##</pre>

### 应当输出

<pre>9</pre>

分析：迷宫最短路径问题，安安分分用bfs宽搜就好。

``` cpp
#include<stdio.h>
#include<iostream>
#include<algorithm>
#include<vector>
#include<cctype>
#include<cmath>
#include<cstring>
#include<queue>
#include<string>
using namespace std;
struct point 
{
    int x,y;
}r[10005];
int dis[4][2]={{1,0},{0,-1},{-1,0},{0,1}};
char map[105][105];
int ans[105][105];
int m,n,si,sj,ei,ej;
void bfs()
{
    int tail=1,head=0,i,x1,y1;  
    r[0].x=si;  
    r[0].y=sj;     
    while(tail != head)  
    {  
        x1=r[head].x;  
        y1=r[head].y;  
        for(i=0; i<4; i++)  
        {  
            x1+=dis[i][0], y1+=dis[i][1];  
            if(x1>=0&&y1>=0&&x1<n&&y1<m&&map[x1][y1]!='#'&&ans[x1][y1]==-1)  
            {  
                r[tail].x=x1;  
                r[tail].y=y1;  
                ans[x1][y1] = 1 + ans[x1-dis[i][0]][y1-dis[i][1]];  
                tail++;  
            }  
            x1-=dis[i][0], y1-=dis[i][1];  
        }  
        head++;   
        if(ans[ei][ej]!=-1)  
            break;  
    }
}
void solve();
int main()
{
    solve();
    return 0;
}
void solve()
{
    int t,i,j;
    cin>>t;
    while(t--)
    {
        cin>>n>>m;
        memset(map,'#',sizeof(map));
        memset(ans,-1,sizeof(ans));
        for(i=0;i<n;i++)
        {
            for(j=0;j<m;j++)
            {
                cin>>map[i][j];
                if(map[i][j]=='S')
                    si=i,sj=j,ans[i][j]=0;
                if(map[i][j]=='E')
                    ei=i,ej=j;
            }
            getchar();
        }
        bfs();
        if(ans[ei][ej]==-1)
            cout<<-1<<endl;
        else
            cout<<ans[ei][ej]<<endl;
    }
}
```