---
title: View的绘制(Android群英传)
date: 2016-05-20 19:28:47
categories:
  - Android群英传
  - 3.Android控件架构与自定义控件详解
tags:
  - 3.Android控件架构与自定义控件详解
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***

　　当测量好了一个View之后，我们就可以简单地重写onDraw()方法，并在Canvas对象上来绘制所需要的图形。首先我们来了解一下利用系统2D绘图API所必须要使用到的Canvas对象。

　　要想在Android的界面中绘制相应的图像，就必须在Canvas上进行绘制。Canvas就像是一个画板，使用Paint就可以在上面作画了。通常需要通过继承View并重写它的onDraw()方法来完成绘图。
<!--more-->
　　那什么是Canvas呢？一般情况下，可以使用重写View类中的onDraw()方法来绘图，onDraw()中有一个参数，就是Canvas canvas对象。使用这个Canvas对象就可以进行绘图了，而在其他地方，通常需要使用代码创建一个Canvas对象，代码如下所示。

```java
Canvas canvas = new Canvas(bitmap);
```

　　当创建一个Canvas对象时，为什么要传进去一个bitmap对象呢？如果不传进入一个bitmap对象，IDE编译虽然不会报错，但是一般我们不会这样做。这是因为传进去的bitmap与通过这个bitmap创建的Canvas画布是紧紧联系在一起的，这个过程我们称之为装载画布。这个bitmap用来存储所有绘制在Canvas上的像素信息。所以当你通过这种方式创建了Canvas对象后，后面调用所有的Canvas.drawXXX方法都发生在这个bitmap上。如果在View类的onQraw()方法中，通过下面这段代码，我们可以了解到canvas与bitmap直接的关系。首先在onDraw方法中绘制两个bitmap，代码如下所示。

```java
canvas.drawBiimap(bitmap1, 0, 0, null);
canvas.drawBitmap(bitmap2, 0, 0, null);
```
　　而对于bitmap2，我们将它装载到另一个Canvas对象中，代码如下所示。

```java
Canvas mCanvas=new Canvas(bitmap2);
```

　　在其他地方使用Canvas对象的绘图方法在装载bitmap2的Canvas对象上进行绘图，代码如下所示。

```java
mCanvas.drawXXX
```

　　通过mCanvas将绘制效果作用在了bitmap2上，再刷新View的时候，就会发现通过onDraw()方法画出来的bitmap2已经发生变化，这就是因为bitmap2承载了在mCanvas上所进行的绘图操作。虽然我们也使用了Canvas的绘制API，但其实并没有将图形直接绘制在onDraw()方法指定的那块画布上，而是通过改变bitmap，然后让View重绘，从而显示改变之后的bitmap。这一过程对初学者来说可能非常难以理解，但是却非常重要，这对后续进行深入地学习和提升绘图技巧非常有帮助。

　　在理解了Canvas对象后，我们就可以调用Canvas所提供的绘图方法，来绘制自己想要的图形了。不管是多么复杂、精美的控件，它都可以被拆分成一个个小的图形单元，我们要做的正是找到这些小的绘图单元并将它们绘制出来。

* * *

原文地址[View的绘制(Android群英传)][passage_url]

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。

[passage_url]: http://blankj.com/2016/05/20/View的绘制(Android群英传)/