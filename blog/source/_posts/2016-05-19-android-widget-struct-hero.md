---
title: Android控件架构(Android群英传)
date: 2016-05-19 20:38:56
categories:
  - Android群英传
  - 3.Android控件架构与自定义控件详解
tags:
  - 3.Android控件架构与自定义控件详解
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***

　　控件大致被分为两类，ViewGroup控件和View控件。ViewGroup可以包含多个View并管理它们。通过ViewGroup，整个界面上的控件形成一个树形结构，也就是我们常说的控件树，上层控件负责下层子控件的测量与绘制，并传递交互事件。通常在Activity中使用findViewById()方法，就是在控件树中以树的深度优先遍历来查找对应元素。在每棵控件树的顶部，都有一个ViewParent对象，这就是整棵树的控制核心，所有的交互管理事件都由它统一调度和分配，从而可以对整个视图进行整体控制。View视图树如下图所示。
<!--more-->
{% asset_img azkjjg1.png [View树结构] %}

　　通常情况下，在Activity中使用setContentView()方法来设置一个布局，在调用该方法后，布局内容才真正显示出来。下面来看一下Android界面的架构图，如下图所示。

{% asset_img azkjjg2.png [UI界面架构图] %}

　　每个Activity都包含一个Window对象，在Android中Window对象通常由PhoneWindow来实现。PhoneWindow将一个DecorView设置为整个应用窗口的根View。DecorView作为窗口界面的顶层视图，封装了一些窗口操作的通用方法。可以说，DecorView将要显示的具体内容呈现在了PhoneWindow上，这里面的所有View的监听事件都通过WindowManagerService来进行接收，并通过Activity对象来回调相应的onClickListener。在显示上，他将屏幕分为两部分，一个是TitleView，另一个是ContentView。看到这里，大家一定看见了一个非常熟悉得布局----ContentView。它是一个ID为content的FrameLayout，activity_main.xml就是设置在这样一个Framelayout里。通过以上过程，我们可以建立起这样一个标准视图树，如下图所示。

{% asset_img azkjjg3.png [标准视图树] %}

　　上图所示的视图树的第二层装在了一个LinearLayout作为ViewGroup，这一层的布局结构会根据对应的参数设置不同的布局，如最常用的布局----上面显示TitleBar，下面是Content这样的布局，也就是图3.3中所设置的布局。而如果用户通过设置requestWindowFeature(Window.FEATURE_NO_TITLE)来设置显示全屏，视图树中的布局就只有Content了，这就解释了为什么调用requestWindowFeature()方法一定要在setContentView()方法之前才能生效的原因。不过这里要注意的是，由于每个Android版本对UI的修改都比较多，上图只是比较粗略地显示了视图树的结构。

　　而在代码中，当程序在onCreat()方法中调用setContentView()方法后，ActivityManagerService会回调onResume()方法，此时系统才会把整个DecorView添加到PhoneWindow中，并让其显示出来，从而最终完成界面的绘制。

* * *

原文地址[Android控件架构(Android群英传)][passage_url]

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。

[passage_url]: http://blankj.com/2016/05/19/Android控件架构(Android群英传)/