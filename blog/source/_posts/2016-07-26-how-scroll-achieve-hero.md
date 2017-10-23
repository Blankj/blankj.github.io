---
title: 滑动效果是如何产生的(Android群英传)
date: 2016-07-26 01:19:25
categories:
  - Android群英传
  - 5.Android Scroll分析
tags:
  - 5.Android Scroll分析
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***  

前言：相对于Android2.X版本中常见的长按、点击操作，滑动的操作方式具有更好的用户体验性。因此，从Android4.X版本开始，滑动操作就大量出现在了Android系统中，各种第三方应用也竞相模仿这种效果，来改善自己的应用，下面就将向大家展示如何在应用中添加滑动效果。  
　　滑动一个View，本质上来说就是移动一个View。改变其当前所处的位置，它的原理与动画效果的实现非常相似，都是通过不断地改变View的坐标来实现这一效果。所以，要实现View的滑动，就必须监听用户触摸的事件，并根据事件传入的坐标，动态且不断地改变View的坐标，从而实现View跟随用户触摸的滑动而滑动。  

　　在讲解如何实现滑动效果之前，需要先来了解一下Android中的窗口坐标体系和屏幕的触控事件——MotionEvent。  
<!-- more -->
### Android坐标系

　　在物理学中，要描述一个物体的运动，就必须选定一个参考系。所谓滑动，正是相对于参考系的运动。在Android中，将屏幕最左上角的顶点作为Android坐标系的原点，从这个点向右是X轴正方向，从这个点向下是Y轴的正方向，如下图所示。  

{% asset_img android_coordinate.png [Android坐标系] %}

　　系统提供了getLocationOnScreen(int location[])这样的方法来获取Android坐标系中点的位置，即该视图左上角在Android坐标系的坐标。另外，在触控事件中使用getRawX()、getRawY()方法所获得的坐标同样是Android坐标系中的坐标。  

### 视图坐标系

　　Android中除了上面所说的这种坐标系之外，还有一个视图坐标系，它描述了子视图在父视图中的位置关系。这两种坐标系并不矛盾也不复杂，他们的作用是相辅相成的。与Android坐标系类似，视图坐标系同样是以原点向右为X轴正方向，以原点向下为Y轴正方向，只不过在视图坐标系中，原点不再是Android坐标系中的屏幕最左上角，而是以**父视图左上角**为坐标原点，如下图所示。  

{% asset_img view_coordinate.png [视图坐标系] %}

　　在触控事件中，通过getX()、getY()所获得的坐标就是视图坐标系中的坐标。  

### 触控事件——MotionEvent

　　触控事件MotionEvent在用户交互中，站着举足轻重的地位，学好触控事件是掌握后序内容的基础。首先，来看看MotionEvent中封装的一些常用的事件常量，它定义了触控事件的不同类型。  

``` java
// 单点触摸按下动作
public static final int ACTION_DOWN             = 0;
// 单点触摸离开动作
public static final int ACTION_UP               = 1;
// 触摸点移动动作
public static final int ACTION_MOVE             = 2;
// 触摸动作取消
public static final int ACTION_CANCEL           = 3;
// 触摸动作超出边界
public static final int ACTION_OUTSIDE          = 4;
// 多点触摸按下动作
public static final int ACTION_POINTER_DOWN     = 5;
// 多点离开动作
public static final int ACTION_POINTER_UP       = 6;
```

　　通常情况下，我们会在onTouchEvent(MotionEvent event)方法中通过event.getAction()方法来获取触控事件的类型，并使用switch-case方法来进行筛选，这个代码的模式基本固定，如下所示。  

``` java
@Override
public boolean onTouchEvent(MotionEvent event) {
    // 获取当前输入点的X、Y坐标(视图坐标)
    int x = (int) event.getX();
    int y = (int) event.getY();
    switch (event.getAction()) {
        case MotionEvent.ACTION_DOWN:
            // 处理输入的按下事件
            break;
        case MotionEvent.ACTION_MOVE:
            // 处理输入的移动事件
            break;
        case MotionEvent.ACTION_UP:
            // 处理输入的离开事件
            break;
    }
    return true;
}
```

　　在不涉及多点操作的情况下，通常可以使用以上代码来完成触控事件的监听，不过这里只是一个代码模板，后面我们会在触控事件中完成具体的逻辑。  

　　在Android中，系统提供了非常多的方法来获取坐标值、相对距离等。方法丰富固然好，但也给初学者带来了很多困惑，不知道在什么情况下使用什么方法，下面总结了一些API，结合Android坐标系来看看该如何使用它们，如下图所示。  

{% asset_img get_xy_method.png [获取坐标值的各种方法] %}

　　这些方法可以分成如下两个类别：  

- View提供的获取坐标方法  
getTop()：获取到的是View自身的顶边到其父布局顶边的距离。  
getLeft()：获取到的是View自身的左边到其父布局左边的距离。  
getRight()：获取到的是View自身的右边到其父布局左边的距离。  
getBottom()：获取到的是View自身的底边到其父布局顶边的距离。  

- MotionEvent提供的方法  
getX()：获取点击事件距离控件左边的距离，即视图坐标。  
getY()：获取点击事件距离控件顶边的距离，即视图坐标。  
getRawX()：获取点击事件距离整个屏幕左边的距离，即绝对坐标。  
getRawY()：获取点击事件距离整个屏幕顶边的距离，即绝对坐标。  

　　相信通过上图，读者们应该对MotionEvent和Android坐标系有了一个比较清楚的认识。 

* * *

原文地址[滑动效果是如何产生的(Android群英传)][passage_url]  

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。  

[passage_url]: http://blankj.com/2016/07/26/滑动效果是如何产生的(Android群英传)/