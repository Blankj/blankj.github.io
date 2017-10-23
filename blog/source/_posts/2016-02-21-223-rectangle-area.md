---
title: 223. Rectangle Area
date: 2016-02-21 14:24:38
categories:
  - LeetCode
  - Java
tags:
  - Math
---

Total Accepted: **30441**
Total Submissions: **104452**
Difficulty: **Easy**

Find the total area covered by two **rectilinear** rectangles in a **2D** plane.

Each rectangle is defined by its bottom left corner and top right corner as shown in the figure.

{% asset_img rectangle_area.png %}

Assume that the total area is never beyond the maximum possible value of **int**.

<!-- more -->

Java:

``` java
public class Solution {
    public int computeArea(int A, int B, int C, int D, int E, int F, int G, int H) {
        int r = (C - A) * (D - B) + (G - E) * (H - F);
        if (!(A >= G || B >= H || C <= E || D <= F))
            r -= (Math.min(C, G) - Math.max(A, E)) * (Math.min(D, H) - Math.max(B, F));
        return r;
    }
}
```