---
title: 自定义View(二)(Android群英传)
date: 2016-07-17 18:17:39
categories:
  - Android群英传
  - 3.Android控件架构与自定义控件详解
tags:
  - 3.Android控件架构与自定义控件详解
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***

上一篇[自定义View(一)(Android群英传)][pre_passage_url]中说的是**对现有控件进行拓展**，这篇介绍第二种自定义View方法，**创建复合控件**。

### 创建复合控件

　　创建复合控件可以很好地创建出具有重用功能的控件集合。这种方式通常需要继承一个合适的ViewGroup，再给它添加指定功能的控件，从而组合成新的复合控件。通过这种方式创建的控件，我们一般会给它指定一些可配置的属性，让它具有更强的拓展性。下面就以一个TopBar为示例，讲解如何创建复合控件。
<!-- more -->
　　我们知道为了应用程序风格的统一，很多应用程序都有一些共通的UI界面，比如下图中所示的TopBar这样一个标题栏。

{% asset_img topbar_full.jpg [界面上的TopBar] %}

　　通常情况下，这些界面都会被抽象出来，形成一个共通的UI组件。所有需要添加标题栏的界面都会引用这样一个TopBar，而不是每个界面都在布局文件中写这样一个TopBar。同时，设计者还可以给TopBar增加响应的接口，让调用者能够更加灵活地控制TopBar，这样不仅可以提高界面的复用率，更能在需要修改UI时，做到快速修改，而不需要对每个页面的标题栏都进行修改。

　　下面我们就来看看该如何创建一个这样的UI模板。首先，模板应该具有通用性与可定制性。也就是说，我们需要给调用者以丰富的接口，让他们可以更改模板中的文字、颜色、行为等信息，而不是所有的模板都一样，那样就失去了模板的意义。

#### 定义属性

　　为一个View提供可自定义的属性非常简单，只需要在res资源目录的values目录下创建一个attrs.xml的属性定义文件，并在该文件中通过如下代码定义相应的属性即可。

```java
<?xml version="1.0" encoding="utf-8"?> 
<resources>
    <declare-styleable name="TopBar">
        <attr name="_title" format="string" />
        <attr name="_titleTextSize" format="dimension" />
        <attr name="_titleTextColor" format="color" />
        <attr name="leftTextColor" format="color" />
        <attr name="leftBackground" format="reference|color" />
        <attr name="leftText" format="string" />
        <attr name="rightTextColor" format="color" />
        <attr name="rightBackground" format="reference|color" />
        <attr name="rightText" format="string" />
    </declare-styleable>
</resources>
```

　　我们在代码中通过<declare-styleable>标签声明了使用自定义属性，并通过name属性来确定引用的名称。最后，通过<attr>标签来声明具体的自定义属性，比如在这里定义了标题文字的字体、大小、颜色，左边按钮的文字颜色、背景、字体，右边按钮的文字颜色、背景、字体等属性，并通过format属性来指定属性的类型。这里需要注意的就是，有些属性可以是颜色属性，也可以是引用属性。比如按钮的背景，可以把它指定为具体的颜色，也可以把它指定为一张图片，所以使用“|”来分隔不同的属性——“reference|color”。

　　在确定好属性后，就可以创建一个自定义控件----TopBar，并让它继承自ViewGroup，从而组合一些需要的控件。这里为了简单，我们继承RelativeLayout。在构造方法中，通过如下所示代码来获取XML布局文件中自定义的那些属性，即与我们使用系统提供的那些属性一样。

```java
TypedArray ta = context.obtainStyledAttributes(attrs, R.styleable.TopBar);
```

　　系统提供了TypedArray这样的数据结构来获取自定义属性集，后面引用的styleable的TopBar，就是我们在XML中通过<declare-styleable name="TopBar">所指定的name名。接下来，通过TypedArray对象的getString()、getColor()等方法，就可以获取这些定义的属性值，代码如下所示。

``` java
// 通过这个方法，将你在attrs.xml中定义的declare-styleable的所有属性值存储到TypedArray中
TypedArray ta = context.obtainStyledAttributes(attrs, R.styleable.TopBar);
// 从TypedArray中取出对应的值来为要设置的属性赋值
mLeftTextColor = ta.getColor(R.styleable.TopBar_leftTextColor, 0);
mLeftBackground = ta.getDrawable(R.styleable.TopBar_leftBackground);
mLeftText = ta.getString(R.styleable.TopBar_leftText);
mRightTextColor = ta.getColor(R.styleable.TopBar_rightTextColor, 0);
mRightBackground = ta.getDrawable(R.styleable.TopBar_rightBackground);
mRightText = ta.getString(R.styleable.TopBar_rightText);
mTitleTextSize = ta.getDimension(R.styleable.TopBar__titleTextSize, 10);
mTitleTextColor = ta.getColor(R.styleable.TopBar__titleTextColor, 0);
mTitle = ta.getString(R.styleable.TopBar__title);
// 获取完TypedArray的值后，一般要调用recycle方法来避免重新创建时的错误
ta.recycle();
```

　　这里需要注意的是，当获取完所有的属性值后，需要调用TypedArray的recycle方法来完成资源的回收。

#### 组合控件

　　接下来，我们就可以开始组合控件了。UI模板TopBar实际上由三个控件组成，即左边的点击按钮mLeftButton，右边的点击按钮mRightButton和中间的标题栏mTitleView。通过动态添加控件的方式，使用addView()方法将这三个控件加入到定义的TopBar模板中，并给它们设置我们前面所获取到的具体的属性值，比如标题的文字颜色、大小等，代码如下所示。

 ``` java
// 为创建的组件元素赋值
// 值就来源于我们在引用的xml文件中给对应属性的赋值
mLeftButton.setTextColor(mLeftTextColor);
mLeftButton.setBackground(mLeftBackground);
mLeftButton.setText(mLeftText);
 
mRightButton.setTextColor(mRightTextColor);
mRightButton.setBackground(mRightBackground);
mRightButton.setText(mRightText);
 
mTitleView.setText(mTitle);
mTitleView.setTextColor(mTitleTextColor);
mTitleView.setTextSize(mTitleTextSize);
mTitleView.setGravity(Gravity.CENTER);
 
// 为组件元素设置相应的布局元素
mLeftLayoutParams = new LayoutParams(LayoutParams.WRAP_CONTENT, LayoutParams.MATCH_PARENT);
mLeftLayoutParams.addRule(RelativeLayout.ALIGN_PARENT_LEFT, TRUE);
// 添加到ViewGroup
addView(mLeftButton, mLeftLayoutParams);
 
mRightLayoutParams = new LayoutParams(LayoutParams.WRAP_CONTENT, LayoutParams.MATCH_PARENT);
mRightLayoutParams.addRule(RelativeLayout.ALIGN_PARENT_RIGHT, TRUE);
addView(mRightButton, mRightLayoutParams);
 
mTitleLayoutParams = new LayoutParams(LayoutParams.WRAP_CONTENT, LayoutParams.MATCH_PARENT);
mTitleLayoutParams.addRule(RelativeLayout.CENTER_IN_PARENT, TRUE);
addView(mTitleView, mTitleLayoutParams);
```

- 定义接口

　　在UI模板类中定义一个左右按钮点击的接口，并创建两个方法，分别用于左边按钮的点击和右边按钮的点击，代码如下所示。

``` java
// 接口对象，实现回调机制，在回调方法中
// 通过映射的接口对象调用接口中的方法
// 而不用去考虑如何实现，具体的实现由调用者去创建
public interface topbarClickListener {
    // 左按钮点击事件
    void leftClick();
 
    // 右按钮点击事件
    void rightClick();
}
```

- 暴露接口给调用者

　　在模板方法中，为左、右按钮增加点击事件，但不去实现具体的逻辑，而是调用接口中相应的点击方法，代码如下所示。

``` java
// 按钮的点击事件，不需要具体的实现，
// 只需调用接口的方法，回调的时候，会有具体的实现
mRightButton.setOnClickListener(new OnClickListener() {
    @Override
    public void onClick(View v) {
        mListener.rightClick();
    }
});
mLeftButton.setOnClickListener(new OnClickListener() {
    @Override
    public void onClick(View v) {
        mListener.leftClick();
    }
});
 
// 暴露一个方法给调用者来注册接口回调
// 通过接口来获得回调者对接口方法的实现
public void setOnTopbarClickListener(topbarClickListener mListener) {
    this.mListener = mListener;
}
```

- 实现接口回调

　　在调用者的代码中，调用者需要实现这样一个接口，并完成接口中的方法，确定具体的实现逻辑，并使用第二步中暴露的方法，将接口的对象传递进去，从而完成回调。通常情况下，可以使用匿名内部类的形式来实现接口中的方法，代码如下所示。

``` java
mTopBar.setOnTopbarClickListener(new MyTopBar.topbarClickListener() {
    @Override
    public void leftClick() {
        Toast.makeText(MainActivity.this,
                "left", Toast.LENGTH_SHORT)
                .show();
    }
    @Override
    public void rightClick() {
        Toast.makeText(MainActivity.this,
                "right", Toast.LENGTH_SHORT)
                .show();
    }
});
```

　　这里为了简单演示，只显示两个Toast来区分不同的按钮点击事件。除了通过接口回调的方式来实现动态的控制UI模板，同样可以使用公共方法来动态地修改UI模板中的UI，这样就进一步提高了模板的可定制性，代码如下所示。

``` java
/**
* 设置按钮的显示与否 通过id区分按钮，flag区分是否显示
*
* @param id   id
* @param flag 是否显示
*/
public void setButtonVisable(int id, boolean flag) {
    if (flag) {
        if (id == 0) {
            mLeftButton.setVisibility(View.VISIBLE);
        } else {
            mRightButton.setVisibility(View.VISIBLE);
        }
    } else {
        if (id == 0) {
            mLeftButton.setVisibility(View.GONE);
        } else {
            mRightButton.setVisibility(View.GONE);
        }
    }
}
```

　　通过如上所示代码，当调用者通过TopBar对象调用这个方法后，根据参数，调用者就可以了动态地控制按钮的显示，代码如下所示。

``` java
// 控制topbar上组件的状态
mTopBar.setButtonVisable(0, true);
mTopBar.setButtonVisable(1, false);
```

#### 引用UI模板

　　最后一步，自然是在需要使用的地方引用UI模板，在引用前，需要指定引用第三方控件的名字空间。在布局文件中，可以看到如下一行代码。

``` java
xmlns:android="http://schemas.android.com/apk/res/android"
```

　　这行代码就是在指定引用的名字空间xmlns，即xml namespace。这里指定了名字空间为“android”，因此在接下来使用系统属性的时候，才可以使用“android:”来引用Android的系统属性。同样地，如果要使用自定义的属性，那么就需要创建自己的名字空间，在Android Studio中，第三方的控件都使用如下代码来引入名字空间。

``` java
xmlns:custom="http://schemas.android.com/apk/res-auto"
```

　　这里我们将引入的第三方控件的名字空间取为custom，之后再XML文件中使用自定义的属性时，就可以通过这个名字空间来引用，代码如下所示。

``` java
<cmj.com.delsys.MyTopBar
    android:id="@+id/topBar"
    android:layout_width="match_parent"
    android:layout_height="40dp"
    custom:_title="自定义标题"
    custom:_titleTextColor="#123412"
    custom:_titleTextSize="10sp"
    custom:leftBackground="@color/colorPrimary"
    custom:leftText="Back"
    custom:leftTextColor="#ffffff"
    custom:rightBackground="@color/colorPrimary"
    custom:rightText="More"
    custom:rightTextColor="#ffffff"/>
```

　　使用自定义的View与系统原生的View最大的区别就是在申明控件时，需要指定完整的包名，而在引用自定义的属性时，需要使用自定义的xmlns名字。

　　再更进一步，如果将这个UI模板写到一个布局文件中，代码如下所示。

``` java
<cmj.com.delsys.MyTopBar
    android:id="@+id/topBar"
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:custom="http://schemas.android.com/apk/res-auto"
    android:layout_width="match_parent"
    android:layout_height="40dp"
    custom:_title="自定义标题"
    custom:_titleTextColor="#123412"
    custom:_titleTextSize="10sp"
    custom:leftBackground="@color/colorPrimary"
    custom:leftText="Back"
    custom:leftTextColor="#ffffff"
    custom:rightBackground="@color/colorPrimary"
    custom:rightText="More"
    custom:rightTextColor="#ffffff"/>
```

　　通过如上所示的代码，我们就可以在其他的布局文件中，直接通过<include>标签来引用这个UI模板View，代码如下所示。

``` java
<include layout="@layout/widget_topbar"/>
```

　　这样就更加满足了我们的模板需求。

　　运行程序后，显示效果如下图所示。

{% asset_img topbar_top.jpg [组合控件] %}

　　当调用公共方法setButtonVisable()来控制左右两个按钮的显示和隐藏的时候，效果显示如下图所示。

{% asset_img topbar_hide_right.jpg [隐藏右按钮] %}

项目地址→[MyTopBar](https://github.com/Blankj/MyTopBar)

* * *

原文地址[自定义View(二)(Android群英传)][passage_url]

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。

[pre_passage_url]: http://blankj.com/2016/05/29/自定义View(一)(Android群英传)/
[passage_url]: http://blankj.com/2016/07/17/自定义View(二)(Android群英传)/