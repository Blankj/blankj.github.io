---
title: 冗余关系
date: 2016-02-05 16:17:55
categories:
  - 数据结构
tags:
  - 并查集
---

蒜头最近在沉迷小说，尤其是人物关系复杂的言情小说。它看到的人物关系描述得很的麻烦的时候觉得非常蒜疼，尤其是人物关系里有冗余的时候。什么是冗余关系呢？

这篇小说里有n句描述人物关系的句子，描述了n个人的关系。

每条句子的定义是这样的：

> X<->Y    它的意思是：X认识Y，Y也认识X

我们认为小说中的人物关系是具有传递性的，假如A认识B，B认识C，则A也认识C。

冗余关系的定义：就是即使没有这条人物关系，原来的人物之间的所有关系也照样成立。

比如：

小说中已经提到了A认识B，B也认识C。在此之后再讲A认识C就是一个冗余的关系。

小蒜头想求出一共有多少条冗余关系，你能帮帮它吗？也许**并查集**能帮上忙哦。

<!-- more -->
### 输入格式：

第一行两个整数，表示句子数量n(1<=n<=1000)，表示人数m(1<=m<=1000)。

接下来n行，每行两个数，表示一组人物关系。

### 输出格式：

一个整数，表示冗余关系的数目。

### 样例1

#### 输入：

<pre>3 3
1 2
1 3
2 3</pre>

#### 输出：

<pre>1</pre>

``` cpp
#include <bits/stdc++.h>
int fa[2000];
// 还记得之前阅读课里讲的并查集算法
// father函数返回的是节点x的祖先节点
int father(int x) {
    if (fa[x] != x) fa[x] = father(fa[x]);
    return fa[x];
}
// 合并两个节点所在集合，同时判断两个点之前是否在一个集合里
// 函数返回true则之前两个点不在一个集合中
bool join(int x, int y) {
    int fx = father(x), fy = father(y);
    if (fx != fy) {
        fa[fx] = fy;
        return true;
    } else {
        return false;
    }
}
// 初始化一个n个点的并查集
void init(int n) {
    for (int i = 1; i <= n; ++i) fa[i] = i;
}
int main(){
    int n,m;
    scanf("%d%d",&n,&m);
    init(m);
    int f,c,ans=0;
    for(int i=0;i<n;++i){
        scanf("%d%d",&f,&c);
        if(!join(c,f)){
            ++ans;
        }       
    }
    printf("%d",ans);
    return 0;
}
```