---
title: ViewGroup的测量与绘制(Android群英传)
date: 2016-05-28 08:47:55
categories:
  - Android群英传
  - 3.Android控件架构与自定义控件详解
tags:
  - 3.Android控件架构与自定义控件详解
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***

### ViewGroup的测量

　　之前分析中说了，ViewGroup会去管理其子View，其中一个管理项目就是负责子View的显示大小。当ViewGroup的大小为wrap_content时，ViewGroup就需要对子View进行遍历，以便获得所有子View的大小，从而来决定自己的大小。而在其他模式下则会通过具体的指定值来设置自身的大小。

　　ViewGroup在测量时通过遍历所有子View，从而调用子View的Measure方法来获得每一个子View的测量结果，前面所说的对View的测量，就是在这里进行的。
<!--more-->
　　当子View测量完毕后，就需要将子View放到合适的位置，这个过程就是View的Layout过程，同样是使用遍历来调用子View的Layout方法，并制定其具体显示的位置，从而来决定其布局位置。

　　在自定义ViewGroup时，通常会去重写onLayout()方法来控制其子View显示位置的逻辑。同样，如果需要支持wrap_content属性，那么它还需要重写onMeasure()方法，这点与View是相同的。

### ViewGroup的绘制

　　ViewGroup通常情况下不需要绘制，因为它本身就没有需要绘制的东西，如果不是指定了ViewGroup的背景颜色，那么ViewGroup的onDraw()方法都不会被调用。但是，ViewGroup会调用dispatchDraw()方法来绘制其子View，其过程同样是通过遍历所有子View，并调用子View的绘制方法来完成绘制工作。

* * *

原文地址[ViewGroup的测量与绘制(Android群英传)][passage_url]

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。

[passage_url]: http://blankj.com/2016/05/28/ViewGroup的测量与绘制(Android群英传)/