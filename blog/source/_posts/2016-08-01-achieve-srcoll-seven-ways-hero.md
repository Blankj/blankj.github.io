---
title: 实现滑动的七种方法(Android群英传)
date: 2016-08-01 00:44:28
categories:
  - Android群英传
  - 5.Android Scroll分析
tags:
  - 5.Android Scroll分析
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***  

　　当了解了Android坐标系和触控事件后，我们再来看看如何使用系统提供的API来实现动态地修改一个View的坐标，即实现滑动效果。而不管采用哪一种方式，其实现的思想基本是一致的，当触摸View时，系统记下当前触摸点坐标；当手指移动时，系统记下移动后的触摸点坐标，从而获取到相对于前一次坐标点的偏移量，并通过偏移量来修改View的坐标，这样不断重复，从而实现滑动过程。  

　　下面我们就通过一个实例，来看看在Android中该如何实现滑动效果。定义一个View，并置于一个LinearLayout中，实现一个简单的布局，代码如下所示。

``` java
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
              android:layout_width="match_parent"
              android:layout_height="match_parent"
              android:orientation="vertical">
 
    <com.blankj.achievescroll.DragView
        android:layout_width="100dp"
        android:layout_height="100dp"
        android:background="#ff09cfb1"/>
</LinearLayout>
```

　　我们的目的就是让这个自定义View随着手指在屏幕上的滑动而滑动。初始化时显示效果如下图所示。

{% asset_img layout_eg.jpg [示例布局] %}

<!-- more -->
### layout方法
　　我们知道，在View进行绘制时，会调用onLayout()方法来设置显示的位置。同样，可以通过修改View的left，top，right，bottom四个属性来控制View的坐标。与前面提供的模板代码一样，在每次回调onTouchEvent的时候，我们都来获取一下触摸点的坐标，代码如下所示。  
``` java
int x = (int) event.getX();
int y = (int) event.getY();
```

　　接着，在ACTION_DOWN事件中记录触摸点的坐标，代码如下所示。  

``` java
case MotionEvent.ACTION_DOWN:
    // 记录触摸点坐标
    lastX = x;
    lastY = y;
    break;
```

　　最后，可以在ACTION_MOVE事件中计算偏移量，并将偏移量作用到Layout的left，top，right，bottom基础上，增加计算出来的偏移量，代码如下所示。  

``` java
case MotionEvent.ACTION_MOVE:
    // 计算偏移量
    int offsetX = x - lastX;
    int offsetY = y - lastY;
    // 在当前left、top、right、bottom的基础上加上偏移量
    layout(getLeft() + offsetX,
            getTop() + offsetY,
            getRight() + offsetX,
            getBottom() + offsetY);
    break;
```
　　这样每次移动后，View都会调用Layout方法来对自己重新布局，从而达到移动View的效果。  

　　在上面的代码中，使用的是getX()、getY()方法来获取坐标值，即通过视图坐标来获取偏移量。当然，同样可以使用getRawX()、getRawY()来获取坐标，并使用绝对坐标来计算偏移量，代码如下所示。  

``` java
@Override
public boolean onTouchEvent(MotionEvent event) {
    int rawX = (int) event.getRawX();
    int rawY = (int) event.getRawY();
    switch (event.getAction()) {
        case MotionEvent.ACTION_DOWN:
            // 记录触摸点坐标
            lastX = rawX;
            lastY = rawY;
            break;
        case MotionEvent.ACTION_MOVE:
            // 计算偏移量
            int offsetX = rawX - lastX;
            int offsetY = rawY - lastY;
            // 在当前left、top、right、bottom的基础上加上偏移量
            layout(getLeft() + offsetX,
                    getTop() + offsetY,
                    getRight() + offsetX,
                    getBottom() + offsetY);
            // 重新设置初始坐标
            lastX = rawX;
            lastY = rawY;
            break;
    }
    return true;
}
```

　　使用绝对坐标系，有一点非常需要注意的地方，就是在每次执行完ACTION_MOVE的逻辑后，一定要重新设置初始坐标，这样才能准确地获取偏移量，两种方式的不同点一定要自己想清楚原因哦。  

### offsetLeftAndRight()与offsetTopAndBottom()

　　这个方法相当于系统提供的一个对左右、上下移动的API的封装。当计算出偏移量后，只需要使用如下代码就可以完成View的重新布局，效果与使用Layout方法一样，代码如下所示。  

``` java
// 同时对left和right进行偏移
offsetLeftAndRight(offsetX);
// 同时对top和bottom进行偏移
offsetTopAndBottom(offsetY);
```

　　这里的offsetX、offSetY与在Layout方法中计算offset方法一样，这里就不重复了。  

### LayoutParams

　　LayoutParams保存了一个View的布局参数。因此可以在程序中，通过改变LayoutParams来动态地修改一个布局的位置参数，从而达到改变View位置的效果。我们可以很方便地在程序中使用getLayoutParams()来获取一个View的LayouParams。当然，计算偏移量的方法与在Layout方法中计算offset也是一样。当获取到偏移量之后，就可以通过setLayoutParams来改变其LayoutParams，代码如下所示。  

``` java
LinearLayout.LayoutParams layoutParams = (LinearLayout.LayoutParams) getLayoutParams();
layoutParams.leftMargin = getLeft() + offsetX;
layoutParams.topMargin = getTop() + offsetY;
setLayoutParams(layoutParams);
```

　　不过这里需要注意的是，通过getLayoutParams()获取LayoutParams时，需要根据View所在父布局的类型来设置不同的类型，比如这里将View放在LinearLayout中，那么就可以使用LinearLayout.LayoutParams。类似地，如果在RelativeLayout中，就要使用RelativeLayout.LayoutParams。当然，这一切的前提是你必须要有一个父布局，不然系统不法获取LayoutParams。  

　　在通过改变LayoutParams来改变一个View的位置时，通常改变的是这个View的Margin属性，所以除了使用布局的LayoutParams之后，还可以使用ViewGroup.MarginLayoutParams来实现这样一个功能，代码如下所示。  

``` java
ViewGroup.MarginLayoutParams layoutParams = (ViewGroup.MarginLayoutParams) getLayoutParams();
layoutParams.leftMargin = getLeft() + offsetX;
layoutParams.topMargin = getTop() + offsetY;
setLayoutParams(layoutParams);
```

　　我们可以发现，使用ViewGroup.MarginLayoutParams更加的方便，不需要考虑父布局的类型，当然他们的本质都是一样的。  

### scrollTo与scrollBy

　　在一个View中，系统提供了scrollTo、scrollBy两种方式来改变一个View的位置。这两个方法的区别非常好理解，与英文的To与By的区别类似，scrollTo(x, y)表示移动到一个具体的坐标点(x, y)，而scrollBy(dx, dy) 表示移动的增量为dx、dy。  

　　与前面几种方式不同，在获取偏移量后使用scrollBy来移动View，代码如下所示。  

``` java
int offsetX = x - lastX;
int offsetY = y - lastY;
scrollBy(offsetX, offsetY);
```

　　但是，当我们拖动View的时候，你会发现View并没有移动！难道是我们方法写错了吗？其实，方法没有写错，View也确实移动了，只是它移动的并不是我们想要移动的东西。scrollTo、scrollBy方法移动的都View的content，即让View的内容移动，如果在ViewGroup中使用scrollTo、scrollBy方法，那么移动的将是所有子View，但如果在View中使用，那么移动的将是View的内容，例如TextView，content就是它的文本；ImageView，content就是它的drawable对象。  

　　相信通过上面的分析，读者朋友应该知道为什么不能在View中使用这个两个方法来拖动这个View了。那么我们就该View所在的ViewGroup中来视同scrollBy方法，移动它的子View，代码如下所示。  

``` java
((View) getParent()).scrollBy(offsetX, offsetY);
```

　　但是，当再次拖动View的时候，你会发现View虽然移动了，但却在乱动，并不是我们想要的跟随触摸点的移动而移动。这里需要先了解一下视图移动的一些知识。大家在理解这个问题的时候，不妨这样想象手机屏幕是一个中空的盖板，盖板下面是一个巨大的画布，也就是我们想要显示的视图。当把这个盖板盖在画布上的某一处时，透过中间空的矩形，我们看见了手机屏幕上的视图，而画布在其他地方的视图，则被盖板盖住了无法看见。我们的视图与这个例子非常类似，我们没有看见视图，并不代表它就不存在，有可能只是在屏幕外面而已。当调用scrollBy方法时，可以想象为外面的盖板在移动，这么说比较抽象，来看一个具体的例子，如下图所示。  

{% asset_img understand_scrollBy.png [理解] %}

　　在上图中，中间的矩形相当于屏幕，即可视区域。后面的content就相当于画布，代表视图。大家可以看到，只有视图的中间部分目前是可视的，其他部分都不可见。在可见区域中，我们设置了一个Button，它的坐标是(20,10)。  

　　下面使用scrollBy方法，将盖板(屏幕、可视区域)，在水平方向上向X轴正方向(右方)平移20，在竖直方向上向Y轴正方向(下方)平移10，那么平移后的可视区域如下图所示。  

{% asset_img after_move_visiable_area.png [移动之后的可视区域] %}

　　我们可以发现，虽然设置scrollBy(20, 10)，偏移量均为X轴、Y轴正方向上的正数，但是在屏幕的可视区域内，Button却向X轴、Y轴负方向上移动了。这就是因为参考系选择的不同，而产生的不同效果。  

　　通过上面的分析可以发现，如果将scrollBy中的参数dx和dy设置为正数，那么content将向坐标轴负方向移动；如果将scrollBy中的参数dx和dy设置为负数，那么content将向坐标轴正方向移动。因此回到前面的例子，要实现跟随手指移动而滑动的效果，就必须将偏移量改为负值，代码如下所示。  

``` java
int offsetX = x - lastX;
int offsetY = y - lastY;
((View) getParent()).scrollBy(-offsetX, -offsetY);
```

　　再去试验一下，大家就可以发现，效果与前面几种方式的效果相同了。类似地，在使用绝对坐标时，也可以通过scrollTo方法来实现这一效果。  

### Scroller

　　既然提到了scrollTo、scrollBy方法，就不得不再来说一说Scroller类。Scroller类与scrollTo、scrollBy方法十分相似，有着千丝万缕的联系。那么它们之间具体有什么区别呢？要解答这个问题，首先来看一个小例子。假如要完成这样一个效果：通过点击按钮，让一个ViewGroup的子View向右移动100个像素。问题看似非常简单，只要在按钮的点击事件中使用前面讲的scrollBy方法设置下偏移量不就可以了吗？的确，通过这样一个方法可以让ViewGroup中的子View平移。但是读者朋友可以发现，不管使用scrollTo还是scrollBy方法，子View的平移都是瞬间发生的，在事件执行的时候平移就已经完成了，这样的效果会让人感觉非常突兀。Google建议使用自然的过度动画来实现移动效果，当然也要遵循这一原则。因此，Scroller类就这样应运而生了，通过Scroller类可以实现平滑移动的效果，而不再是瞬间完成的移动。  

　　说到Scroller类的实现原理，其实它与前面使用scrollTo和scrollBy方法来实现子View跟随手指移动的原理基本类似。虽然scrollBy方法是让子View瞬间从某点移动到另一个点，但是由于在ACTION_MOVE事件中不断获取手指移动的微小的偏移量，这样就将一段距离划分成了N个非常小的偏移量。虽然在每个偏移量里面，通过scrollBy方法进行了瞬间移动，但是在整体上却可以获得一个平滑移动的效果。这个原理与动画的实现原理基本类似，他们都是利用了人眼的视觉暂留特性。  

　　下面我们就演示一下如何使用Scroller类实现平滑移动。在这个实例中，同样让子View跟随手指的滑动而滑动，但是在手指离开屏幕时，让子View平滑地移动到初始位置，即屏幕左上角。一般情况下，使用Scroller类需要如下三个步骤。  

- 初始化Scroller

　　首先，通过它的构造方法来创建一个Scroller对象，代码如下所示。  
``` java
// 初始化Scroller
mScroller = new Scroller(context);
```

- 重写computeScroll()方法，实现模拟滚动

　　下面我们需要重写computeScroll()方法，它是使用Scroller类的核心，系统在绘制View的时候会在draw()方法中调用该方法。这个方法实际上就是使用scrollTo方法。再结合Scroller对象，帮助获取到当前的滚动值。我们可以通过不断地瞬间移动一个小的距离来实现整体上的平滑移动效果。通常情况下，computeScroll的代码可以利用如下模板代码来实现。  

``` java
@Override
public void computeScroll() {
    super.computeScroll();
    // 判断Scroller是否执行完毕
    if (mScroller.computeScrollOffset()) {
        ((View) getParent()).scrollTo(
                mScroller.getCurrX(),
                mScroller.getCurrY());
        // 通过重绘来不断调用computeScroll
        invalidate();
    }
}
```

　　Scroller类提供了computeScrollOffset()方法来判断是否完成了整个滑动，同时也提供了getCurrX()、getCurrY()方法来获取当前滑动坐标。在上面的代码中，唯一需要注意的是invalidate()方法，因为只能在computeScroll()方法中获取模拟过程中的scrollX和scrollY坐标。但computeScroll()方法是不会自动调用的，只能通过invalidate()→draw()→computeScroll()来间接调用computeScroll()方法，所以需要在模板代码中调用invalidate()方法，实现循环获取scrollX和scrollY的目的。而当模拟过程结束后，scroller.computeScrollOffset()方法会返回false，从而中断循环，完成整个平滑移动过程。  

- startScroll开启模拟过程

　　最后，万事俱备只欠东风。我们在需要使用平滑移动事件中，使用Scroller类的startScroll()方法来开启平滑移动过程。startScroll()方法具有两个重载方法。  

 - `public void startScroll(int startX, int startY, int dx, int dy, int duration)`  
 - `public void startScroll(int startX, int startY, int dx, int dy)`  

　　可以看到他们的区别就是一个具有指定的持续时长，而另一个没有。这个非常好理解，与在动画中设置durarion和使用默认的显示时长是一个道理。而其他四个坐标，则与它们的命名含义相同，就是起始坐标与偏移量。在获取坐标时，通常可以使用getScrollX()和getScrollerY()方法来获取父视图中content所滑动到的电的坐标，不过要注意的是这个值的正负，它与在scrollBy和scrollTo中讲解的情况是一样的。  

　　通过上面三个步骤，我们就可以使用Scroller类来实现平滑移动了，下面回到实例中，在构造方法中初始化Scroller对象，并重写View的computeScroll()方法。最后，需要监听手指离开屏幕的事件，并在该事件中通过调用startScroll()方法完成平滑移动。那么要监听手指离开屏幕的事件，只需要在onTouchEvent中增加一个ACTION_UP监听选项即可，代码如下所示。  

``` java
case MotionEvent.ACTION_UP:
    // 手指离开时，执行滑动过程
    View viewGroup = ((View) getParent());
    mScroller.startScroll(
            viewGroup.getScrollX(),
            viewGroup.getScrollY(),
            -viewGroup.getScrollX(),
            -viewGroup.getScrollY());
    invalidate();
```

　　在startScroll()方法中，我们获取子View移动的距离——getScrollX()、getScrollY()，并将偏移量设置为其相反数，从而将子View滑动到原来位置。这里需要注意的还是invalidate()方法，需要使用这个方法来通知View进行重绘，从而来调用conputeScroll()的模拟过程。当然，也可以给startScroll()方法增加一个duration的参数来设置滑动的持续时长。  

### 属性动画

　　为视图增加位移动画，视图进行位移偏移后，利用视图动画在松手后视图回到原处，具体代码如下所示。  

``` java
@Override
public boolean onTouchEvent(MotionEvent event) {
    int x = (int) event.getX();
    int y = (int) event.getY();
    switch (event.getAction()) {
        case MotionEvent.ACTION_DOWN:
            // 记录触摸点坐标
            lastX = x;
            lastY = y;
            break;
        case MotionEvent.ACTION_MOVE:
            // 计算偏移量
            int offsetX = x - lastX;
            int offsetY = y - lastY;// 同时对left和right进行偏移
            offsetLeftAndRight(offsetX);
            // 同时对top和bottom进行偏移
            offsetTopAndBottom(offsetY);
            break;
        case MotionEvent.ACTION_UP:
            // 手指离开时，执行滑动过程
            ObjectAnimator animator1 = ObjectAnimator.ofFloat(this, "translationX", -getLeft());
            ObjectAnimator animator2 = ObjectAnimator.ofFloat(this, "translationY", -getTop());
            AnimatorSet set = new AnimatorSet();
            set.playTogether(animator1, animator2);
            set.start();
            break;
    }
    return true;
}
```

### ViewDragHelper

　　Google在其support库中为我们提供了DrawerLayout和SlidingPaneLayout两个布局来帮助开发者实现侧边栏滑动的效果。这两个新的布局，大大方便了我们创建自己的滑动布局界面。然而，这两个功能强大的布局背后，却隐藏着一个鲜为人知却功能强大的类——ViewDragHelper。通过ViewDragHelper，基本可以实现各种不同的滑动、拖放需求，因此这个方法也是各种滑动方案中的终极绝招。  

　　ViewDragHelper虽然功能强大，但其使用方法也是这次最复杂的。读者朋友需要在理解ViewDragHelper基本使用方法的基础上，通过不断练习来掌握它的技巧。下面通过一个实例，来演示一下如何使用ViewDragHelper创建一个滑动布局。在这个例子中，准备实现类似QQ滑动侧边栏的布局，初始时显示内容界面，当用户手指滑动超过一段距离时，内容界面侧滑显示菜单界面，整个过程如下图所示。  

{% asset_img init_and_slip_show_menu.png [初始状态即滑动展开菜单界面] %}

　　下面来看具体的代码是如何实现的。  

- 初始化ViewDragHelper

　　首先，自然是需要初始化ViewDragHelper。ViewDragHelper通常定义在一个ViewGroup的内部，并通过其静态工厂方法进行初始化，代码如下所示。  

``` java
mViewDragHelper = ViewDragHelper.create(this, callback);
```

　　它的第一个参数是要监听的View，通常需要是一个ViewGroup，即parentView；第二个参数是一个Callback回调，这个回调就是整个ViewDragHelper的逻辑核心，后面再来详细讲解。  

- 拦截事件

　　接下来，要重写事件拦截方法，将事件传递给ViewDragHelper进行处理，代码如下所示。  

``` java
@Override
public boolean onInterceptTouchEvent(MotionEvent ev) {
    return super.onInterceptTouchEvent(ev);
}
@Override
public boolean onTouchEvent(MotionEvent event) {
    // 将触摸事件传递给ViewDragHelper，此操作必不可少
    mViewDragHelper.processTouchEvent(event);
    return true;
}
```

　　这一点我们在讲Android事件机制的时候已经进行了详细讲解，这里就不再重复了。  

- 处理computeScroll

　　没错，使用ViewDragHelper同样需要重写下computeScroll()方法，因为ViewDragHelper内部也是通过Scroller来实现平滑移动的。通常情况下，可以使用如下所示的模板代码。  

``` java
@Override
public void computeScroll() {
    if (mViewDragHelper.continueSettling(true)) {
        ViewCompat.postInvalidateOnAnimation(this);
    }
}
```

- 处理回调Callback

　　下面就是最关键的Callback实现，通过如下所示代码来创建一个ViewDragHelper.Callback。  

``` java
private ViewDragHelper.Callback callback = new ViewDragHelper.Callback() {
    @Override
    public boolean tryCaptureView(View child, int pointerId) {
        return false;
    }
};
```

　　IDE自动帮我们重写了一个方法——tryCaptureView()。通过这个方法，我们可以指定在创建ViewDragHelper时，参数parentView中哪一个子View可以被移动，例如在这个实例中自定义了一个ViewGroup，里面定义了两个子View——MenuView和MainView，当指定如下代码时，则只有MainView是可以被拖动的。  

``` java
// 何时开始检测触摸事件
@Override
public boolean tryCaptureView(View child, int pointerId) {
    // 如果当前触摸的child是mMainView时开始检测
    return mMainView == child;
}
```

　　下面来看具体的滑动方法——clampViewPositionVertical()和clampViewPositionHorizontal()，分别对应垂直和水平方向上的滑动。如果要实现滑动效果，那么这两个方法是必须要重写的。因为它默认的返回值是0，即不发生滑动。当然，如果只重写clampViewPositionVertical()或clampViewPositionHorizontal()中的一个，那么就只会实现该方向上的滑动效果了，代码如下所示。  

``` java
@Override
public int clampViewPositionVertical(View child, int top, int dy) {
    return top;
}
@Override
public int clampViewPositionHorizontal(View child, int left, int dx) {
    return left;
}
```

　　clampViewPositionVertical(View child, int top, int dy)中的参数top，代表在垂直方向上child移动的距离，而dy则表示比较前一次的增量。同理，clampViewPositionHorizontal(View child, int left, int dx)也是类似的含义。通常情况下，只需要返回top和left即可，但当需要更加精确地计算padding等属性的时候，就需要对left进行一些处理，并返回合适大小的值。  

　　仅仅是通过重写上面的这三个方法，就可以实现一个最基本的滑动效果了。当用手拖动MainView的时候，它就可以跟随手指的滑动而滑动了，代码如下所示。  

``` java
private ViewDragHelper.Callback callback = new ViewDragHelper.Callback() {
    // 何时开始检测触摸事件
    @Override
    public boolean tryCaptureView(View child, int pointerId) {
        // 如果当前触摸的child是mMainView时开始检测
        return mMainView == child;
    }
    @Override
    public int clampViewPositionVertical(View child, int top, int dy) {
        return 0;
    }
    @Override
    public int clampViewPositionHorizontal(View child, int left, int dx) {
        return left;
    }
};
```

　　下面继续来优化这个实例。在讲解Scroller类时，曾实现了这样一个效果——在手指离开屏幕后，子View滑动回初始位置。当时我们是通过监听ACTION_UP事件，并通过调用Scroller类来实现的，这里使用ViewDragHelper来实现这样的效果。在ViewDragHelper.Callback中，系统提供了这样的方法——onViewReleased()，通过重写这个方法，可以非常简单地实现当手指离开屏幕后实现的操作。当然，这个方法内部是通过Scroller类来实现的，这也是前面重写computeScroll()方法的原因，这部分代码如下所示。  

``` java
// 拖动结束后调用
@Override
public void onViewReleased(View releasedChild, float xvel, float yvel) {
    super.onViewReleased(releasedChild, xvel, yvel);
    // 手指抬起后缓慢移动到指定位置
    if (mMainView.getLeft() < 500) {
        // 关闭菜单
        // 相当于Scroller的startScroll方法
        mViewDragHelper.smoothSlideViewTo(mMainView, 0, 0);
    } else {
        // 打开菜单
        mViewDragHelper.smoothSlideViewTo(mMainView, 300, 0);
    }
    ViewCompat.postInvalidateOnAnimation(DragViewGroup.this);
}
```

　　设置让MainView移动后左边距小于500像素的时候，就是用smoothSlideViewTo()方法来将MainView还原到初始状态，即坐标为(0, 0)的点。而当其左边距大于500的时候，则将MainView移动到(300, 0)坐标，即显示MenuView。读者朋友可以发现如下所示的这两行代码，与在使用Scroller类的时候使用的startScroll()方法是不是非常像呢？  

``` java
// ViewDragHelper
mViewDragHelper.smoothSlideViewTo(mMainView, 0, 0);
ViewCompat.postInvalidateOnAnimation(DragViewGroup.this);

// Scroll
mScroller.startScroll(x, y, dx, dy);
invalidate();
```

　　通过前面一步步的分析，现在要实现类似QQ侧滑菜单的效果，是不是就非常简单了呢？下面自定义一个ViewGroup来完成整个实例的编写。滑动的处理部分前面已经讲解过了，在自定义ViewGroup的onFInishInflate()方法中，按顺序将子View分别定义成MenuView和MainView，并在onSizeChanged()方法中获得View的宽度。如果你需要根据View的宽度来处理滑动后的效果，就可以使用这个值来进行判断。这部分代码如下所示。  

``` java
@Override
protected void onFinishInflate() {
    super.onFinishInflate();
    mMenuView = getChildAt(0);
    mMainView = getChildAt(1);
}
@Override
protected void onSizeChanged(int w, int h, int oldw, int oldh) {
    super.onSizeChanged(w, h, oldw, oldh);
    mWidth = mMenuView.getMeasuredWidth();
}
```

　　最后，整个通过ViewDragHelper实现QQ侧滑功能的完整代码参考项目地址即可。  

　　当然，这里只是非常简单地模拟了QQ侧滑菜单这个功能。ViewDragHelper的很多强大功能还没能够得到展示。在ViewDragHelper.Callback中，系统定义了大量的监听事件来帮助我们吹各种事件，下面就列举一些事件。  

- onViewCaptured()

　　这个事件在用户触摸到View后调用。  

- onViewDragStateChanged()

　　这个事件在拖拽状态改变时回调，比如idle，dragging等状态。  

- onViewPositionChanged()

　　这个事件在位置改变时回调，常用于滑动时更改scale进行缩放等效果。  

　　使用ViewDragHelper可以帮助我们非常好地处理程序中的滑动效果。但同时ViewDragHelper的使用也比较复杂，需要开发者对事件拦截、滑动处理都有比较清楚的认识。所以建议初学者循序渐进，在掌握前面几种解决方案的基础上，再来学习ViewDragHelper，以实现更加丰富的滑动效果。  

项目地址→[AchieveScroll](https://github.com/Blankj/AchieveScroll)

* * *

原文地址[实现滑动的七种方法(Android群英传)][passage_url]  

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。  

[passage_url]: http://blankj.com/2016/08/01/实现滑动的七种方法(Android群英传)/