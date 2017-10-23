---
title: 并查集
date: 2016-02-05 15:48:31
categories:
  - 数据结构
tags:
  - 并查集
---

并查集是一个多棵树的数据结构(森林)，每个节点记录一个father[i]表示它的父节点。并查集可以解决集合的合并和查询操作。

<!-- more -->
并查集的初始化代码如下：

``` cpp
void make_set() {
    for (int i = 0; i < n; ++i)
        father[i] = i;
}
```

此时每个节点都属于不同的集合。查询操作的代码如下：

``` cpp
int get_father(int v) {
    return father[v] != v ? get_father(father[v]) : v;
}
```

上面代码的意思就是，如果当前节点不是这棵树的根节点，那么就不断回溯到根节点。

接下来是合并操作：

``` cpp
void merge(int x, int y) {
    int root_x = find(x), root_y = find(y);
    if (root_x != root_y) father[root_x] = root_y;
}
```

下图是合并操作的图解：

{% asset_img union_find.jpg %}

并查集的一个优化叫做【路径压缩】，是在并查集执行查询时对经过的点进行【扁平化】的方法。优化后的代码如下：

``` cpp
int get_father(int v) {
    if (father[v] != v) father[v] = get_father(father[v]);
    return father[v];
}
```