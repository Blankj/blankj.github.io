---
title: 等差数列
date: 2015-03-22 03:07:23
categories:
  - OJ
tags:
  - 动态规划
  - 枚举
---

### 题目描述

给定n(1<=n<=400)个数,从中找出尽可能多的数使得他们能够组成一个等差数列.求最长的等差数列的长度.

<!-- more -->
### 输入要求

第一行是一个整数n,接下来一行包括了n个数,每个数的绝对值不超过10000000.

### 输出要求

对于每个输入数据,输出你所找出的最长等差数列的长度

### 假如输入

<pre>7
3
8
4
5
6
2
2</pre>

### 应当输出

<pre>5</pre>


分析：枚举法，n^3时间复杂度，另一种方法是dp，dp[j][i]表示以a[j]和a[i]结尾的最长等差数列的长度。枚举最后两个元素，对于每一个a[j]和a[i]，都要找到a[p]，p < j，满足a[p] + a[i] == 2 * a[j]。然后dp[p][j] + 1去更新dp[j][i]。看起来是三层循环，但其实对于同一个i，p的位置是随着j增大而增大的，所以最里面的while循环对于每个i值最多是O(n)的代价。总的代价还是O(n^2)。

枚举代码：

``` cpp
#include<iostream>
#include<algorithm>
#include<vector>
#include<cctype>
#include<cmath>
#include<cstring>
#include<cstdio>
using namespace std;
void solve();
int main()
{
	solve();
	return 0;
}
void solve()
{   
	int i,j,k,d,n,arr[401],ans=0,sum,la;
	cin>>n;
	if (n<=2)
		cout<<n<<endl;
	else
	{
		for(i=0;i<n;i++)
			cin>>arr[i];
		sort(arr,arr+n);
		for (i=0;i<n-1;++i)
		{
			for (j=i+1;j<n;++j)
			{
				d=arr[j]-arr[i];
				la=j;
				sum=2;
				for (k=j+1;k<n;++k)
				{
					if (arr[k]-arr[la]==d) 
					{
						la=k;
						++sum;
					}
					else if(arr[k]-arr[la]>d)
						break;
				}
				if (sum>ans)  
					ans=sum;
			}
			cout<<ans<<endl;
		}
	}
}
```

dp代码：

``` cpp
#include<iostream>
#include<algorithm>
#include<vector>
#include<cctype>
#include<cmath>
#include<cstring>
#include<cstdio>
using namespace std;
#define MAXL 410
int dp[MAXL][MAXL];
int find_(int a[], int n)
{        
        int i, j;
        for(i = 1; i < n; i++) {
                int p = 0;
                for(j = 0; j < i; j++) {
                        dp[j][i] = 2;
                        while(p < j && a[p] + a[i] < 2 * a[j])
                                p++;
                        if(p < j && a[p] + a[i] == 2 * a[j])
                                dp[j][i] = max(dp[j][i], dp[p][j] + 1);
                }
        }
        int ans = (n > 0);
        for(i = 1; i < n; i++)
                for(j = 0; j < i; j++)
                        ans = max(ans, dp[j][i]);
        return ans;

}
void solve();
int main()
{
    solve();
    return 0;
}
void solve()
{   
    int i,n,arr[401];
    cin>>n;
    for(i=0;i<n;i++)
        cin>>arr[i];
    sort(arr,arr+n);
    cout<<find_(arr,n);
}
```