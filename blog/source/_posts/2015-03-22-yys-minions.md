---
title: YY's Minions
date: 2015-03-22 14:34:22
categories:
  - OJ
tags:
  - 模拟
---

Despite YY's so much homework, she would like to take some time to play with her minions first.

YY lines her minions up to an N*M matrix. Every minion has two statuses: awake or asleep. We use 0(the digit) to represent that it is asleep, and 1 for awake. Also, we define the minions who are around a minion closest in one of the eight directions its neighbors. And every minute every minion will change its status by the following specific rules:
If this minion is awake, and the number of its neighbors who are awake is less than 2, this minion will feel lonely and turn to asleep.
If this minion is awake, and the number of its neighbors who are awake is more than 3, this minion will turn to asleep for it will feel too crowded.
If this minion is awake, and the number of its neighbors who are awake is exactly 2 or 3, this minion will keep being awake and feel very happy.
If this minion is asleep, and the number of its neighbors who are awake is exactly 3, this minion will wake up because of the noise.
Note that all changes take place at the same time at the beginning of a specific minute.
Also, some minions will get bored and leave this silly game. We use 'X's to describe them. We suppose that a minion would leave after T minutes. It will leave at the end of the Tth minute. Its status is considered during the change at the beginning of the Tth minute, and should be ignored after that. Of course, one minion will not leave twice!

YY is a girl full of curiosity and wants to know every minion's status after F minutes. But you know she is weak and lazy! Please help this cute girl to solve this problem :)

<!-- more -->
### Input
There are multiple test cases.

The first line contains the number of test cases Q. 1<=Q<=100.
For each case, there are several lines:
The first line contains four integers N, M, F, K. K means the number of leaving messages. 1<=N, M<=50, 1<=F<=1000, 1<=K<=N*M. 
Next N lines are the matrix which shows the initial status of each minion. Each line contains M chars. We guarantee that 'X' wouldn't appear in initial status matrix. 
And next K lines are the leaving messages. Each line contains three integers Ti, Xi, Yi, They mean the minion who is located in (Xi, Yi) will leave the game at the end of the Tith minutes. 1<=Ti<= F, 1<=Xi<=N, 1<=Yi<=M.

### Output
For each case, output N lines as a matrix which shows the status of each minion after F minutes.

### Sample Input
<pre>2
3 3 2 1
101
110
001
1 2 2
5 5 6 3
10111
01000
00000
01100
10000
2 3 3
2 4 1
5 1 5</pre>

### Sample Output
<pre>010
1X0
010
0000X
11000
00X00
X0000
00000</pre>

### Hint

For case 1:

T=0, the game starts
<pre>101
110
001</pre>

---------------
at the beginning of T=1, a change took place
<pre>100
101
010</pre>

---------------
at the end of T=1 (the minion in (2,2) left)
<pre>100
1X1
010</pre>

---------------
at the beginning of T=2, a change took place
<pre>010
1X0
010</pre>

---------------
at the end of T=2 (nothing changed for no minion left at T=2)
<pre>010
1X0
010</pre>


**分析：题目虽然长了点，但很好理解，注意那四个条件即可模拟求得。**

``` cpp
#include<iostream>
#include<algorithm>
#include<vector>
#include<cctype>
#include<cmath>
#include<cstring>
#include<cstdio>
using namespace std;
int cnt[55][55],lea[55][55];//cnt周围醒着的数量，lea将要离开的时间
char sta[55][55];//当前状态
int dir[8][2]={{1, 0}, {0, 1}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}};//遍历的八个方向
void solve();
int main()
{
	solve();
	return 0;
}
void solve()
{   
	int t,i,j,q,n,m,f,k;
	cin>>q;	
	while(q--)
	{		
		memset(sta,0,sizeof(sta));
		memset(cnt,0,sizeof(cnt));
		memset(lea,0,sizeof(cnt));
		cin>>n>>m>>f>>k;
		getchar();//去除回车
		for(i=1;i<=n;i++)//得到初始状态图
		{
			for(j=1;j<=m;j++)
				sta[i][j]=getchar();
			getchar();//去除回车
		}
		while(k--)
		{
			int t,x,y;
			cin>>t>>x>>y;
			lea[x][y]=t;//x,y处要离开的时间
		}
		for(t=1;t<=f;t++)
		{
			for(i=1;i<=n;i++)
			{
				for(j=1;j<=m;j++)
				{
					for(k=0;k<8;k++)//遍历八个方向
					{
						int x=i+dir[k][1];
	                    int y=j+dir[k][0];
						if(x<1 || y<1 || x>n || y>m)
							continue;
						if(sta[x][y]=='1')
							++cnt[i][j];
					}
				}
			}
			for(i=1;i<=n;i++)
			{
				for(j=1;j<=m;j++)
				{
					if(cnt[i][j]==3 && sta[i][j]=='0')//自己睡着，周围醒着的数量恰好为3将醒来
	                     sta[i][j]='1';
					else if((cnt[i][j]<2 || cnt[i][j]>3) && sta[i][j]=='1')//自己醒着，并且如果周围醒着的数量小于2或大于3将入睡
	                     sta[i][j]='0';
					//其余状态就不用改变
					if(lea[i][j]==t)  
						sta[i][j]='X';//到了t时间离开
		            cnt[i][j]=0;//重置醒着的数量
				}
			}
		}
		for(i=1;i<=n;i++)
		{
			for(j=1;j<=m;j++)
				putchar(sta[i][j]);
			puts("");
		}
	}
}
```