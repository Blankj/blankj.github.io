---
title: 利用EditText的RightDrawable切换密码显示
date: 2016-07-21 17:09:18
categories:
  - Android
  - 自定义View
tags:
  - 自定义View
---

　　相信大家都见过如下图所示的密码文本输入框，点击右方的图标便可切换为明文显示密码。  

{% asset_img etpw.png [密码文本输入框] %}

　　在安卓中，我们可以充分利用EditText的RightDrawable来实现这样的效果，同理一键清除也可实现，其效果图如下所示。  

{% asset_img etp.gif [EditTextPassword] %}

　　下面对其进行简单介绍，首先是布局文件，很简单，就一个ImageView和一个自定义的EditText，代码如下所示。  
<!--more-->
``` java
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:layout_marginLeft="8dp"
    android:layout_marginRight="8dp">
 
    <ImageView
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:layout_gravity="center_vertical"
        android:layout_marginRight="8dp"
        android:background="@drawable/locked"/>
 
 
    <blankj.edittextpassword.EditTextPassword
        android:id="@+id/etp_input"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:layout_gravity="center_vertical"
        android:background="@drawable/input"
        android:drawableRight="@drawable/eye_grey"
        android:singleLine="true"
        android:textColor="#ff7000"
        android:textCursorDrawable="@null"/>
 
</LinearLayout>
```

　　再是对EditText进行拓展，首先要获取到EditText的RightDrawable，代码如下所示。  

``` java
final int DRAWABLE_RIGHT = 2;
drawableRight = getCompoundDrawables()[DRAWABLE_RIGHT];
```

　　然后定义它的回调接口和外部访问接口函数，代码如下所示。  

``` java
public interface DrawableRightListener {
    void onDrawableRightClick();
}
public void setDrawableRightListener(DrawableRightListener listener) {
    this.mRightListener = listener;
}
```

　　在onTouchEvent()中判断是否点击到RightDrawable来设置是否启用回调，在此还需保存光标位置，再点击完之后进行复位，优化用户体验，具体代码如下所示。  

``` java
@Override
public boolean onTouchEvent(MotionEvent event) {
    if (drawableRight != null && mRightListener != null && event.getAction() == MotionEvent.ACTION_UP) {
        int x = (int) event.getX();
        int y = (int) event.getY();
        //判断点击是否落在rightDrawable中
        if (x > getWidth() - getTotalPaddingRight() && x < getWidth() && y > 0 && y < getHeight()) {
            //获取点击之前光标的位置
            int pos = getSelectionStart();
            //设置回调
            mRightListener.onDrawableRightClick();
            //恢复点击之前光标的位置
            this.setSelection(pos);
        }
    }
    return super.onTouchEvent(event);
}
```

　　最后就是在使用的地方进行定义回调函数的作用了，代码如下所示。  

``` java
inputETP.setDrawableRightListener(this);

//同样地可以在这可以实现其他的效果，比如一键清空
@Override
public void onDrawableRightClick() {
    if (mIsShow) {
        inputETP.setCompoundDrawablesWithIntrinsicBounds(0, 0, R.drawable.eye_grey, 0) ;
        inputETP.setInputType(InputType.TYPE_CLASS_TEXT | EditorInfo.TYPE_TEXT_VARIATION_PASSWORD);
    } else {
        inputETP.setCompoundDrawablesWithIntrinsicBounds(0, 0, R.drawable.eye_orange, 0) ;
        inputETP.setInputType(InputType.TYPE_TEXT_VARIATION_VISIBLE_PASSWORD);
    }
    mIsShow = !mIsShow ;
}
```

项目地址→[EditTextPassword](https://github.com/Blankj/EditTextPassword)

***

原文地址[利用EditText的RightDrawable切换密码显示][passage_url]  

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。  

[passage_url]: http://blankj.com/2016/07/21/利用EditText的RightDrawable切换密码显示/