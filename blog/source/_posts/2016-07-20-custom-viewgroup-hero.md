---
title: 自定义ViewGroup(Android群英传)
date: 2016-07-20 18:00:35
categories:
  - Android群英传
  - 3.Android控件架构与自定义控件详解
tags:
  - 3.Android控件架构与自定义控件详解
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***

　　前面我们分析了如何自定义View，下面我们继续来分析如何创建自定义ViewGroup。ViewGroup存在的目的就是为了对其子View进行管理，为其子View添加显示、响应的规则。因此，自定义ViewGroup通常需要重写onMeasure()方法来对子View进行测量，重写onLayout()方法来确定子View的位置，重写onTouchEvent()方法增加响应事件。下面通过一个实例，来看看如何自定义ViewGroup。

　　本例准备实现一个类似Android原生控件ScrollView的自定义ViewGroup，自定义ViewGroup可以实现ScrollView所具有的上下滑动功能，但是在滑动的过程中，增加一个黏性的效果，即当一个子View向上滑动大于一定的距离后，松开手指，它将自动活动到开始的位置，相信大家在很多App应用中都看见过这样的效果。
<!-- more -->
　　首先让自定义ViewGroup能够实现类似ScrollView的功能。

　　当然，在ViewGroup能够滚动之前，需要先放置好它的子View。使用遍历的方式来通知子View对自身进行测量，代码如下所示。

``` java
@Override
protected void onMeasure(int widthMeasureSpec, int heightMeasureSpec) {
    super.onMeasure(widthMeasureSpec, heightMeasureSpec);
    int count = getChildCount();
    for (int i = 0; i < count; ++i) {
        View childView = getChildAt(i);
        measureChild(childView, widthMeasureSpec, heightMeasureSpec);
    }
}
```

　　接下来，就要对子View进行放置位置的设定。让每个子View都显示完整的一屏，这样在滑动的时候，可以比较好地实现后面的效果。在放置子View前，需要确定整个ViewGroup的高度。在本例中，由于让每个子View占一屏的高度，因此整个ViewGroup的高度即子View的个数乘以屏幕的高度，我们通过如下代码来确定整个ViewGroup的高度。

``` java
// 设置ViewGroup的高度
MarginLayoutParams mlp = (MarginLayoutParams) getLayoutParams();
mlp.height = mScreenHeight * childCount;
setLayoutParams(mlp);
```

　　在获取了整个ViewGroup的高度之后，就可以通过遍历来设定每个子View需要放置的位置了，直接通过调用子View的layout()方法，并将具体的位置作为参数传递进去即可，代码如下所示。

``` java
@Override
protected void onLayout(boolean changed, int l, int t, int r, int b) {
    int childCount = getChildCount();
    // 设置ViewGroup的高度
    MarginLayoutParams mlp = (MarginLayoutParams) getLayoutParams();
    mlp.height = mScreenHeight * childCount;
    setLayoutParams(mlp);
    for (int i = 0; i < childCount; ++i) {
        View child = getChildAt(i);
        if (child.getVisibility() != View.GONE) {
            child.layout(l, i * mScreenHeight, r, (i + 1) * mScreenHeight);
        }
    }
}
```

　　在代码中主要是去修改每个子View的top和bottom这两个属性，让它们能依次排列下来。

　　通过上面的步骤，就可以将子View放置到ViewGroup中了。但此时的ViewGroup还不能响应任何触控事件，自然也不能滑动，因此我们需要重写onTouchEvent()方法，为ViewGroup添加响应事件。在ViewGroup中添加滑动事件，通常可以使用scrollBy()方法来辅助滑动。在onTouchEvent()的ACTION_MOVE事件中，只要使用scrollBy(0,dy)方法，让手指滑动的时候让ViewGroup中的所有子View也跟着滚动dy即可，计算dy的方法有很多，如下代码就提供了一种思路。

``` java
case MotionEvent.ACTION_DOWN:
    mLastY = y;
    break;
case MotionEvent.ACTION_MOVE:
    if (!mScroller.isFinished()) {
        mScroller.abortAnimation();
    }
    int dy = mLastY - y;
    if (getScrollY() < 0 || getScrollY() > getHeight() - mScreenHeight) {
        dy = 0;
    }
    scrollBy(0, dy);
    mLastY = y;
    break;
```

　　按如上方法操作就可以实现类似ScrollView的滚动效果了。当然，系统的原生ScrollView有更大的功能，比如滑动的惯性效果等，这些功能可以在后面慢慢添加，这也是一个控件的迭代过程。

　　最后，我们来实现这个自定义ViewGroup的黏性效果。要实现手指离开后ViewGroup的黏性效果，我们很自然地想到onTouchEvent()的ACTION_UP事件和Scroller类。在ACTION_UP事件中判断手指滑动的距离，如果超过一定距离，则使用Scroller类来平滑移动到下一个子View；如果小于一定距离，则回滚到原来的位置，代码如下所示。

``` java
case MotionEvent.ACTION_DOWN:
    // 记录触摸起点
    mStart = getScrollY();
    break;
case MotionEvent.ACTION_UP:
    // 记录触摸终点
    mEnd = getScrollY();
    int dScrollY = mEnd - mStart;
    if (dScrollY > 0) {
        if (dScrollY < mScreenHeight / 3) {
            mScroller.startScroll(
                    0, getScrollY(),
                    0, -dScrollY);
        } else {
            mScroller.startScroll(
                    0, getScrollY(),
                    0, mScreenHeight - dScrollY);
        }
    } else {
        if (-dScrollY < mScreenHeight / 3) {
            mScroller.startScroll(
                    0, getScrollY(),
                    0, -dScrollY);
        } else {
            mScroller.startScroll(
                    0, getScrollY(),
                    0, -mScreenHeight - dScrollY);
        }
    }
    break;
```

　　通过以上操作，我们就能在onTouchEvent()中实现滚动的逻辑和“黏性”的逻辑，整个onTouchEvent()的代码如下所示。

``` java
@Override
public boolean onTouchEvent(MotionEvent event) {
    int y = (int) event.getY();
    switch (event.getAction()) {
        case MotionEvent.ACTION_DOWN:
            // 记录触摸起点
            mStart = getScrollY();
            mLastY = y;
            break;
        case MotionEvent.ACTION_MOVE:
            if (!mScroller.isFinished()) {
                mScroller.abortAnimation();
            }
            int dy = mLastY - y;
            if (getScrollY() < 0 || getScrollY() > getHeight() - mScreenHeight) {
                dy = 0;
            }
            scrollBy(0, dy);
            mLastY = y;
            break;
        case MotionEvent.ACTION_UP:
            // 记录触摸终点
            mEnd = getScrollY();
            int dScrollY = mEnd - mStart;
            if (dScrollY > 0) {
                if (dScrollY < mScreenHeight / 3) {
                    mScroller.startScroll(
                            0, getScrollY(),
                            0, -dScrollY);
                } else {
                    mScroller.startScroll(
                            0, getScrollY(),
                            0, mScreenHeight - dScrollY);
                }
            } else {
                if (-dScrollY < mScreenHeight / 3) {
                    mScroller.startScroll(
                            0, getScrollY(),
                            0, -dScrollY);
                } else {
                    mScroller.startScroll(
                            0, getScrollY(),
                            0, -mScreenHeight - dScrollY);
                }
            }
            break;
    }
    postInvalidate();
    return true;
}
```

　　当然，最后不要忘记加上computeScroll()的代码，如下所示。

``` java
@Override
public void computeScroll() {
    super.computeScroll();
    if(mScroller.computeScrollOffset()){
        scrollTo(0,mScroller.getCurrY());
        postInvalidate();
    }
}
```

　　程序运行效果如下图所示。

{% asset_img cvg.gif [自定义ViewGroup] %}

项目地址→[CustomViewGroup](https://github.com/Blankj/CustomViewGroup)

* * *

原文地址[自定义ViewGroup(Android群英传)][passage_url]

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。

[passage_url]: http://blankj.com//2016/07/20/自定义ViewGroup(Android群英传)/