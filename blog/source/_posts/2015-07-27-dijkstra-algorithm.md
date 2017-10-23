---
title: 数据结构-图的邻接矩阵表示
date: 2015-07-27 23:42:26
categories:
  - 数据结构
tags:
  - 图
---

给出邻接矩阵描述的图，顶点数<26，求

1、边数

2、无向图各顶点度数

3、求两个顶点最短路径

<!-- more -->
### 输入要求

第一行一个数N，表示邻接矩阵的行数。

N*N的邻接矩阵，第一行第一列为顶点A。

求最短路径的两个顶点，用大写字母表示。
### 输出要求

第一行输出边的个数。

接下来N行输出各顶点度数。

最后输出给定两点间最短距离。
### 假如输入

<pre>4
0 25 40 100
25 0 45 -1
40 45 0 80
100 -1 80 0
B D</pre>

### 应当输出

<pre>5
3
2
3
2
125</pre>

### 提示

-1表示没有路径

用的是这几天学的迪杰斯特拉算法

Code：

``` cpp
#include<bits/stdc++.h>
using namespace std;
#define max 60
#define INF 0x7fff
typedef struct
{
    int no;
    int arcnum;
}VertexType;
typedef struct
{
    VertexType vex[max];
    int edge[max][max];
    int n,e;
}MGraph;
int dist[max],path[max];
void Dijkstra(MGraph g,int v)
{
    int visit[max];
    int min,i,j,u;
    for(i=0;i<g.n;++i)
    {
        dist[i]=g.edge[v][i];
        visit[i]=0;
        if(g.edge[v][i]<INF)
            path[i]=v;
        else
            path[i]=-1;
    }
    visit[v]=1;
    path[v]=-1;
    for(i=0;i<g.n;++i)
    {
        min=INF;
        for(j=0;j<g.n;++j)
        {
            if(visit[j]==0&&dist[j]<min)
            {
                u=j;
                min=dist[j];
            }
        }
        visit[u]=1;
        for(j=0;j<g.n;++j)
        {
            if(visit[j]==0&&dist[u]+g.edge[u][j]<dist[j])
            {
                dist[j]=dist[u]+g.edge[u][j];
                path[j]=u;
            }
        }
    }
}
int main()
{
    int i,j;
    MGraph g;
    char begin,end;
    g.e=0;
    cin>>g.n;
    for(i=0;i<g.n;++i)
    {
        g.vex[i].no=i;
        g.vex[i].arcnum=0;
        for(j=0;j<g.n;++j)
        {
            cin>>g.edge[i][j];
            if(g.edge[i][j]!=0&&g.edge[i][j]!=-1)
            {
                ++g.e;
                ++g.vex[i].arcnum;
            }
            else
                g.edge[i][j]=INF;//将自身和没路径的弧设置为无穷
        }
    }
    cin>>begin>>end;
    cout<<g.e/2<<endl;
    for(i=0;i<g.n;++i)
        cout<<g.vex[i].arcnum<<endl;    
    Dijkstra(g,begin-'A');
    cout<<dist[end-'A'];
}
```