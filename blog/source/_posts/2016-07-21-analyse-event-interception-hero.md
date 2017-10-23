---
title: 事件拦截机制分析(Android群英传)
date: 2016-07-21 14:30:09
categories:
  - Android群英传
  - 3.Android控件架构与自定义控件详解
tags:
  - 3.Android控件架构与自定义控件详解
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***

　　当Android系统捕获到用户的各种输入事件后，如何准确地传递给真正需要这个事件的控件呢？Android给我们提供了一整套完善的事件传递、处理机制，来帮助开发者完成准确的事件分配与处理。

　　要了解触摸事件的拦截机制，首先要了解什么是触摸事件？顾名思义，触摸事件就是捕获触摸屏幕后产生的事件。当点击一个按钮时，通常就会产生两个或者三个事件——按钮按下，这是事件一；如果不小心滑动一点，这就是事件二；当手抬起，这是事件三。Android为触摸事件封装了一个类——MotionEvent，如果重写onTouchEvent()方法，那就会发现给方法的参数就是这样一个MotionEvent。其实，只要是重写触摸相关的方法，参数一般都含有MotionEvent，可见它的重要性。

　　在MotionEvent里面封装了不少好东西，比如触摸点的坐标，可以通过event.getX()方法和event.getRawX()方法取出坐标点；再比如获得点击的事件类型，可以通过不同的Action(如MotionEvent.ACTION_DOWN、MotionEvent.ACTION_MOVE)来进行区分，并实现不同的逻辑。
<!--more-->
　　如此看来，触摸事件还是比较简单的，其实就是一个动作类型加坐标而已。但是我们知道，Android的View结构是树形结构，也就是说，View可以放在ViewGroup里面，通过不同的组合来实现不同的样式。那么问题来了，View放在一个ViewGroup里面，这个ViewGroup又放在另一个ViewGroup里面，甚至还有可能继续嵌套，一层层地叠起来。可我们的触摸事件就一个，到底该分给谁呢？同一个事件，子View和父ViewGroup都有可能想要进行处理。因此，这就产生了“事件拦截”这个“霸气”的称呼。

　　当然，事件拦截可以很复杂，也可以很简单。但是初学者却经常“卡”在这里不知道如何继续进行，所以我们不想通过过多的源代码让大家不知所措。我们通过最直观的Log信息，让大家先有一个大概的了解，知道事件拦截的本质，然后大家在结合源代码学习时，就可以有方向、有目的性去理解了。

　　首先，请想象一下生活中非常常见的场景：假设你所在的公司，有一个总经理，级别最高；他下面有一个部长，级别次之；最低层，就是干活的你，没有级别。现在董事会交给总经理一项任务，总经理将这项任务布置给了部长，部长又把任务安排给了你。而当你好不容易干完活了，你就把任务交给部长，部长觉得任务完成得不错，于是就签了他的名字交给总经理，总经理看了也觉得不错，就业签了名字交给董事会。这样，一个任务就顺利完成了。如果大家能非常清楚地理解这样一个场景，那么对于事件拦截机制，你就超过了40%的开发者了。下面，我们再来超越剩下的开发者。为了能过方便地了解整个事件的流程，我们设计了这样一个实例，如下图所示。

{% asset_img sjljsl.png [事件拦截实例] %}

　　一个总经理——MyViewGroupA，最外层的ViewGroup（红色）。

　　一个部长——MyViewGroupB，中间的ViewGroup（绿色）。

　　一个干活的你——MyView，在最底层（蓝色）。

　　本实例的整个布局结构如下图所示。

{% asset_img slbjjg.png [实例布局结构] %}

　　代码非常简单，只是重写了事件拦截和处理的几个方法，并给它加上了一些Log而已。

　　对于ViewGroup来说，重写了如下所示的三个方法。

``` java
@Override
public boolean dispatchTouchEvent(MotionEvent ev) {
    Log.d("blankj", "ViewGroupA dispatchTouchEvent" + ev.getAction());
    return super.dispatchTouchEvent(ev);
}
@Override
public boolean onInterceptTouchEvent(MotionEvent ev) {
    Log.d("blankj", "ViewGroupA onInterceptTouchEvent" + ev.getAction());
    return super.onInterceptTouchEvent(ev);
}
@Override
public boolean onTouchEvent(MotionEvent event) {
    Log.d("blankj", "ViewGroupA onTouchEvent" + event.getAction());
    return super.onTouchEvent(event);
}
```

　　而对于View来说，重写了如下所示的两个方法。

``` java
@Override
public boolean onTouchEvent(MotionEvent event) {
    Log.d("blankj", "View onTouchEvent" + event.getAction());
    return super.onTouchEvent(event);
}
@Override
public boolean dispatchTouchEvent(MotionEvent event) {
    Log.d("blankj", "View dispatchTouchEvent" + event.getAction());
    return super.dispatchTouchEvent(event);
}
```

　　从上面的代码中可以看到，ViewGroup级别比较高，比View多了一个方法——onInterceptTouchEvent()。这个方法看名字就能猜到是事件拦截的核心方法。我们先不修改任何返回值，只是点击一下View，然后看Log会怎样记录我们的操作和程序响应。点击View后的Log如下所示。

> D/blankj: ViewGroupA dispatchTouchEvent0
> D/blankj: ViewGroupA onInterceptTouchEvent0
> D/blankj: ViewGroupB dispatchTouchEvent0
> D/blankj: ViewGroupB onInterceptTouchEvent0
> D/blankj: View dispatchTouchEvent0
> D/blankj: View onTouchEvent0
> D/blankj: ViewGroupB onTouchEvent0
> D/blankj: ViewGroupA onTouchEvent0

　　可以看见，正常情况下，时间的传递顺序是：

　　总经理(MyViewGroupA)→部长(MyViewGroupB)→你(View)。事件传递的时候，先执行dispatchTouchEvent()方法，再执行onInterceptTouchEvent()方法。

　　事件的处理顺序是：

　　你(View)→部长(MyViewGroupB)→总经理(MyViewGroupA)。事件处理都是执行onTouchEvent()方法。

　　事件传递的返回值非常容易理解：true，拦截，不继续；false，不拦截，继续流程。

　　事件处理的返回值也类似：true，处理了，不用审核了；false，给上级处理。

　　初始情况下，返回值都是false。

　　这里为了能够方便大家理解事件拦截的过程，在事件传递中，我们只关心onInterceptTouchEvent()，而dispatchTouchEvent()方法虽然是事件分发的第一步，但一般情况下，我们不太会去改写这个方法，所以暂时不管这个方法。可以把上面的整个事件过程整理成如下图所示的一张图。

{% asset_img sjclgc.png [事件处理过程] %}

　　相信大家只要把MyView想成自己，就能充分理解事件分发、拦截、处理的整个流程了。

　　下面我们稍微改动一下，假设总经理(MyViewGroupA)发现这个任务太简单了，觉得自己完成就可以了，完全没必要再找下属。因此时间就被总经理(MyViewGroupA)使用onInterceptTouchEvent()方法把事件给拦截了，即让MyViewGroupA的onInterceptTouchEvent()方法返回true，我们再来看一下Log。

> D/blankj: ViewGroupA dispatchTouchEvent0
> D/blankj: ViewGroupA onInterceptTouchEvent0
> D/blankj: ViewGroupA onTouchEvent0

　　跟我们设想的一样，总经理(MyViewGroupA)把所有事情都干了，没后面人的事了。同理，我们让部长(MyViewGroupB)也来当一次好人，即让部长(MyViewGroupB)使用onInterceptTouchEvent()方法返回true，把事件拦截下来，Log就会使以下这样。

> D/blankj: ViewGroupA dispatchTouchEvent0
> D/blankj: ViewGroupA onInterceptTouchEvent0
> D/blankj: ViewGroupB dispatchTouchEvent0
> D/blankj: ViewGroupB onInterceptTouchEvent0
> D/blankj: ViewGroupB onTouchEvent0
> D/blankj: ViewGroupA onTouchEvent0

　　可以看到，这次部长(MyViewGroupB)当了好人，你(MyView)就不用干活了。

　　那么这两种情况，也可以整理成类似如上图所示的图。

　　总经理(MyViewGroupA)拦截事件，如下图所示。

{% asset_img sjljA.png [事件拦截A] %}

　　部长(MyViewGroupB)拦截事件，如下图所示。

{% asset_img sjljB.png [事件拦截B] %}

　　对事件的分发、拦截，现在大家应该比较清楚了，下面我们再看看事件的处理。先来看看底层人民——你(MyView)。最开始的时候讲了，当你处理完任务后会向上级报告，需要上级的确认，所以你的事件处理返回false。那么你突然有一天受不了老板的压迫了，罢工不干了，那么你的任务就没人做了，也就不用报告上机了，所以就直接返回true。现在再来看看Log，如下所示。

> D/blankj: ViewGroupA dispatchTouchEvent0
> D/blankj: ViewGroupA onInterceptTouchEvent0
> D/blankj: ViewGroupB dispatchTouchEvent0
> D/blankj: ViewGroupB onInterceptTouchEvent0
> D/blankj: View dispatchTouchEvent0
> D/blankj: View onTouchEvent0

　　可以看见，事件传递跟以前一样，但是事件处理，到你(MyView)这就结束了，因为你返回true，表示不用向上级汇报了。这时，我们同样来整理下关系图，如下所示。

{% asset_img sjclA.png [事件处理A] %}

　　你(MyView)终于翻身做了主，决定了自己的命运。但是，如果部长(MyViewGroupB)看到了你的报告，觉得太丢人，不敢给经理看，所以他就偷偷地返回true，整个事件也就到此为止了，即部长(MyViewGroupB)将自己的onTouchEvent返回true，Log如下所示。

> D/blankj: ViewGroupA dispatchTouchEvent0
> D/blankj: ViewGroupA onInterceptTouchEvent0
> D/blankj: ViewGroupB dispatchTouchEvent0
> D/blankj: ViewGroupB onInterceptTouchEvent0
> D/blankj: View dispatchTouchEvent0
> D/blankj: View onTouchEvent0
> D/blankj: ViewGroupB onTouchEvent0

　　他们之间的关系图如下图所示。

{% asset_img sjclB.png [事件处理B] %}

　　通过对前面几种情况的分析，相信大家能比较容易地了解事件的分发、拦截、处理事件的流程了。在后面的学习中，结合源码，你会更加深入地理解，为什么流程会是这样的？初学者在学习的时候，最好先对流程有一个大致的认识之后，再去接触源码，这样就不会一头雾水，摸不着头脑，从而丧失学习的兴趣。

项目地址→[EventIntercept](https://github.com/Blankj/EventIntercept)

* * *

原文地址[事件拦截机制分析(Android群英传)][passage_url]

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。

[passage_url]: http://blankj.com/2016/07/21/事件拦截机制分析(Android群英传)/