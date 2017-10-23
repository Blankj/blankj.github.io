---
title: 自定义View(三)(Android群英传)
date: 2016-07-20 09:51:45
categories:
  - Android群英传
  - 3.Android控件架构与自定义控件详解
tags:
  - 3.Android控件架构与自定义控件详解
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***

上一篇[自定义View(二)(Android群英传)][pre_passage_url]中说的是**创建复合控件**，这篇介绍第三种自定义View方法，**重写View来实现全新的控件**。

### 重写View来实现全新的控件

　　当Android系统原生的控件无法满足我们的需求时，我们就可以完全创建一个新的自定义View来实现需要的功能。创建一个自定义View，难点在于绘制控件和实现交互，这也是评价一个自定义View优劣的标准之一。通常需要继承View类，并重写它的onDraw()、onMeasure()等方法来实现绘制逻辑，同时通过重写onTouchEvent()等触控事件来实现交互逻辑。当然，我们还可以像实现组合控件方式那样，通过引入自定义属性，丰富自定义View的可定制性。

　　下面就通过几个实例，让大家了解如何创建一个自定义View，不过为了让程序尽可能简单，我们就不去自定义属性值了。
<!-- more -->
#### 弧线展示图

　　在PPT的很多模板中，都有如下图所示的这样一张比例图。

{% asset_img cpb.png [比例图] %}

　　这个比例图可以非常清楚地展示一个项目所占的比例，简介明了。因此，实现这样一个自定义View用在我们的程序中，可以让整个程序实现比较清晰地数据展示效果。那么该如何创建一个这样的自定义View呢？很明显，这个自定义View其实分为三个部分，分别是中间的圆形、中间显示的文字和外圈的弧线。既然有了这样的思路，只要在onDraw()方法中一个个去绘制就可以了。这里为了简单，我们把View的绘制长度直接设置为屏幕的宽度。首先，在初始化的时候，设置好绘制三种图形的参数。圆的代码如下所示。

``` java
mCircleXY = length / 2;
mRadius = (float) (length * 0.5 / 2);
```

　　绘制弧线，需要指定其椭圆的外接矩形，代码如下所示。

``` java
mArcRectF = new RectF(
        (float) (length * 0.1),
        (float) (length * 0.1),
        (float) (length * 0.9),
        (float) (length * 0.9));
```

　　绘制文字，只需要设置好文字的起始绘制位置即可。

　　接下来，我们就可以在onDraw()方法中进行绘制了，代码如下所示。

``` java
// 绘制圆
canvas.drawCircle(mCircleXY, mCircleXY, mRadius, mCirclePaint);
// 绘制弧线
canvas.drawArc(mArcRectF, 270, mSweepAngle, false, mArcPaint);
// 绘制文字
canvas.drawText(mShowText, 0, mShowText.length(), mCircleXY, mCircleXY + mShowTextSize / 4, mTextPaint);
```

　　相信这些图形如果单独让你去绘制，应该是非常容易的事情，只是这里进行了一下组合，就创建了一个新的View。其实，不论是多么复杂的图形、控件，它都是由这些最基本的图形绘制出来的，关键就在于你如何去分解、设计这些图形，当你脑海中有了一幅设计图之后，剩下的事情就只是对坐标的计算了。

　　当然，对于这个简单的View，有一些方法可以让调用者来设置不同的状态值，代码如下所示。

``` java
public void setSweepValue(float sweepValue) {
    if (sweepValue != 0) {
        mSweepValue = sweepValue;
    } else {
        mSweepValue = 25;
    }
    this.invalidate();
}
```
　　例如，当用户不指定具体的比例值时，可以默认设置为25，而调用者可以通过如下代码来设置相应的比例值。

``` java
CircleProgressView circle = (CircleProgressView) findViewById(R.id.circle);
circle.setSweepValue(70);
```

#### 音频条形图

　　以下这个问题来源于群里一位开发者的问题，他想实现类似在PC上某些音乐播放器上根据音频音量大小显示音频条形图，如下图所示。

{% asset_img yptxt.png [音频条形图] %}

　　如果要实现一个如下图所示的静态音频条，相信大家可以很快找到思路，也就是绘制一个个矩形，每个矩形之间稍微偏移一点距离即可。

{% asset_img jtyptxt.png [静态音频条形图] %}

　　如下代码就展示了一种计算坐标的方法。

``` java
@Override
protected void onDraw(Canvas canvas) {
    super.onDraw(canvas);
    for (int i = 0; i < mRectCount; ++i) {
        canvas.drawRect(mRectWidth * i + offset,
                currentHeight,
                mRectWidth * (i + 1),
                mRectHeight,
                mPaint);
    }
}
```

　　如上代码中，我们通过循环创建这些小的矩形，其中currnetHeight就是每个小矩形的高，通过横坐标的不断偏移，就绘制出了这些静态的小矩形。下面我们再让这些小矩形的高度进行随机变化，通过Math.random()方法来随机改变这些高度值，并赋值给currentHeight，代码如下所示。

``` java
mRandom = Math.random();
float currentHeight = (float) (mRectHeight * mRandom);
```

　　这样，我们就完成了静态效果的绘制，那么如何实现动态效果呢？其实非常简单，只要在onDraw()方法中再去调用invalidete()方法通知View进行重绘就可以了。不过，在这里不需要每次一绘制完新的矩形就通知View进行重绘，这样会因为刷新速度太快反而影响效果。因此，我们可以使用如下代码进行View的延迟重绘，代码如下所示。

``` java
postInvalidateDelayed(300);
```

　　这样每隔300ms通知View进行重绘，就可以得到一个比较好的视觉效果了。最后，为了让自定义View更加逼真，可以在绘制小矩形的时候，给绘制的Paint对象增加一个LinearGradient渐变效果，这样不同高度的矩形就会有不同颜色的渐变效果，更加能够模拟音频条形图的风格，代码如下所示。

``` java
@Override
protected void onSizeChanged(int w, int h, int oldw, int oldh) {
    super.onSizeChanged(w, h, oldw, oldh);
    mWidth = getWidth();
    mRectHeight = getHeight();
    mRectWidth = mWidth / mRectCount;
    mLinearGradient = new LinearGradient(0, 0,
            mRectWidth, mRectHeight,
            Color.YELLOW, Color.GREEN,
            Shader.TileMode.CLAMP);
    mPaint.setShader(mLinearGradient);
}
```

　　从这个例子中，我们可以知道，在创建自定义View的时候，需要一步步来，从一个基本的效果开始，慢慢地增加功能，绘制更复杂的效果。不论是多么复杂的自定义View，它一定是慢慢迭代起来的功能，所以不要觉得自定义View有多难。千里之行始于足下，只要开始做，慢慢地就能越来越熟练。

项目地址→[RewriteView](https://github.com/Blankj/RewriteView)

* * *

原文地址[自定义View(三)(Android群英传)][passage_url]

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。

[pre_passage_url]: http://blankj.com/2016/07/17/自定义View(二)(Android群英传)/
[passage_url]: http://blankj.com/2016/07/20/自定义View(三)(Android群英传)/