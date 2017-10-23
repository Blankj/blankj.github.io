---
title: ListView常用优化技巧(Android群英传)
date: 2016-07-24 16:21:15
categories:
  - Android群英传
  - 4.ListView使用技巧
tags:
  - 4.ListView使用技巧
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。*** 

前言：ListView——列表，它作为一个非常重要的显示方式，不管是在Web中还是移动平台中，都是一个非常好的、不开或缺的展示信息的工具。在Android中，ListView控件接管了这一重担，在大量的场合下，我们都需要使用这个控件。虽然在Android 5.X时代，RecyclerView在很多地方都在逐渐取代ListView，但ListView的使用范围依然非常的广泛，它这万年老大哥的地位也不是轻易就能撼动的。下面就介绍一下ListView常用优化技巧。  <!--more-->

### 使用ViewHolder模式提高效率

　　ViewHolder模式是提高ListView效率的一个很重要的方法。ViewHolder模式充分利用了ListView的视图缓存机制，避免了每次在调用getView()的时候都去通过findViewById()实例化控件。据测试，使用ViewHolder将提高50%以上的效率。使用ViewHolder模式来优化ListView非常简单，只需要在自定义Adapter中定义一个内部类ViewHolder，并将布局中的控件作为成员变量，代码如下所示。  

``` java
public final class ViewHolder {
    public ImageView img;
    public TextView title;
}
```

　　接下来，只要在getView()方法中通过视图缓存机制来重用以缓存即可，完整的使用ViewHolder创建ListView Adapter的实例代码如下所示。  

``` java
import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.ImageView;
import android.widget.TextView;
 
import java.util.List;
 
/*********************************************
* author: Blankj on 2016/7/23 15:39
 * blog:   http://blankj.com
* e-mail: blankj@qq.com
*********************************************/
public class ViewHolderAdapter extends BaseAdapter {
 
    private List<String> mData;
    private LayoutInflater mInflater;
 
    public ViewHolderAdapter(Context context, List<String> data) {
        this.mData = data;
        mInflater = LayoutInflater.from(context);
    }
 
    @Override
    public int getCount() {
        return mData.size();
    }
 
    @Override
    public Object getItem(int position) {
        return mData.get(position);
    }
 
    @Override
    public long getItemId(int position) {
        return position;
    }
 
    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        ViewHolder viewHolder = null;
        // 判断是否缓存
        if (convertView == null) {
            viewHolder = new ViewHolder();
            // 通过LayoutInflater实例化布局
            convertView = mInflater.inflate(R.layout.viewholder_item, null);
            viewHolder.img = (ImageView) convertView.findViewById(R.id.imageView);
            viewHolder.title = (TextView) convertView.findViewById(R.id.textView);
            convertView.setTag(viewHolder);
        }else {
            // 通过tag找到缓存的布局
            viewHolder = (ViewHolder) convertView.getTag();
        }
        // 设置布局中控件要显示的视图
        viewHolder.img.setBackgroundResource(R.mipmap.ic_launcher);
        viewHolder.title.setText(mData.get(position));
        return convertView;
    }
 
    public final class ViewHolder {
        public ImageView img;
        public TextView title;
    }
}
```
　　效果很简单，这就是一个简单的ListView，如下图所示。  

{% asset_img use_vh_lv.jpg [使用ViewHolder的ListView] %}

### 设置项目间分隔线

　　ListView的各个项目之间，可以通过设置分隔线来进行区分，系统提供了divider和dividerHeight这样两个属性来帮助我们实现这一功能。通过这两个属性，也可以控制ListView之间的分隔线和它的高度。当然，分隔线不仅仅可以设置为一个颜色，同样也可以设置为一个图片资源，分隔线的使用代码如下所示。  

``` java
android:divider="@color/colorAccent"
android:dividerHeight="10dp"
```

　　以上代码所实行的效果如下图所示。  

{% asset_img lv_divider.jpg [ListView分隔线] %}

　　特殊情况下，当设置分隔线为如下代码时，就可以把分隔线设置为透明了。  

``` java
android:divider="@null"
```

### 隐藏ListView的滚动条

　　默认的ListView在滚动时，在右边会显示滚动条，指示当前滑动的位置，我们可以设置scrollbars属性，控制ListView的滚动条状态。特别地，当设置scrollbars属性为none的时候，ListView滚动或者不滚动，就都不会出现滚动条了，代码如下所示。  

``` java
android:divider="@null"
```

### 取消ListView的Item点击效果

　　当点击ListView中的一项时，系统默认会出现一个点击效果，在Android5.X上是一个波纹效果，而在Android5.X之下的版本则是一个改变背景颜色的效果，但可以通过修改listSelector属性来取消掉点击后的回馈效果，代码如下所示。  

``` java
android:listSelector="#00000000"
```

　　当然，也可以直接使用Android自带的透明色来实现这个效果，代码如下所示。  

``` java
android:listSelector="@android:color/transparent"
```

### 设置ListView需要显示在第几项

　　ListView以Item为单位进行显示，默认显示在第一个Item，当需要指定具体显示的Item时，可以通过如下代码来实现。  

``` java
mListView.setSelection(N);
```

　　其中N就是需要显示的第N个Item。

　　当然，这个方法类似scrollTo，是瞬间完成的移动。除此以外，还可以使用如下代码来实现平滑移动。  

``` java
mListView.smoothScrollBy(distance,duration);
mListView.smoothScrollByOffset(offset);
mListView.smoothScrollToPosition(index);
```

### 动态修改ListView

　　ListView中的数据在某些情况下是需要变化的，当然可以通过重新设置ListView的Adapter来更新ListView的显示，但这也就需要重新获取一下数据，相当于重新刷新创建的ListView，这样显然不是非常友好，而且效率也不会太高。因此，可以使用一个更简单的方法来实现ListView的动态修改，代码如下所示。  

``` java
mData.add("new");
mAdapter.notifyDataSetChanged();
```

　　当修改了传递给Adapter的映射List之后，只需要通过调用Adapter的notifyDataSetChanged()方法，通知ListView更改数据源即可完成对ListView的动态修改。不过，使用这个方法有一点需要注意的是，在使用mAdapter.notifyDataSetChanged()方法时，必须保证传进Adapter的数据List是同一个List而不能是其他对象，否则将无法实现该效果。下面这个实例就演示了如何动态地修改ListView。通过点击按钮，不断地给原有的List增加一个新的Item，并调用notifyDataSetChanged()方法来实现ListView的动态更新，完整代码如下所示。  

``` java
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.ListView;
 
import java.util.ArrayList;
 
public class MainActivity extends AppCompatActivity {
 
    ListView mListView;
    ViewHolderAdapter mAdapter;
    ArrayList<String> mData;
 
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        mListView = (ListView) findViewById(R.id.listView);
        mData = new ArrayList<>();
        for (int i = 0; i < 20; ++i) {
            mData.add(i + "");
        }
        mAdapter = new ViewHolderAdapter(this, mData);
        mListView.setAdapter(mAdapter);
        mAdapter.notifyDataSetChanged();
        for (int i = 0, len = mListView.getCheckedItemCount(); i < len; i++) {
            View view = mListView.getChildAt(i);
        }
    }
 
    public void btnAdd(View view) {
        mData.add("new");
        mAdapter.notifyDataSetChanged();
        mListView.setSelection(mData.size() - 1);
    }
}
```

　　实现的效果如下图所示。  

{% asset_img dynamic_update_ls.jpg [动态更新ListView] %}

### 遍历ListView中的所有Item

　　ListView作为一个ViewGroup，为我们提供了操纵子View的各种方法，最常用的就是通过getChildAt()来获取第i个子View，代码如下所示。  

``` java
for (int i = 0, len = mListView.getCheckedItemCount(); i < len; i++) {
    View view = mListView.getChildAt(i);
}
```

### 处理空ListView

　　ListView用于展示列表数据，但当列表中无数据时，ListView不会显示任何数据或提示，按照完善用户体验的需求，这里应该给以无数据的提示。幸好，ListView提供了一个方法——setEmptyView()，通过这个方法，我们可以给ListView设置一个在空数据下显示的默认提示。包含ListView的布局设置如下。  

``` java
<?xml version="1.0" encoding="utf-8"?>
<FrameLayout
    xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:tools="http://schemas.android.com/tools"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    tools:context="com.blankj.listviewskill.MainActivity">
 
    <ListView
        android:id="@+id/listView"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:divider="@null"
        android:listSelector="@android:color/transparent"
        android:paddingBottom="40dp"/>
 
    <ImageView
        android:id="@+id/empty_view"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:src="@mipmap/ic_launcher"/>
</FrameLayout>
```

　　在代码中，我们通过以下方式给ListView设置空数据时要显示的布局，代码如下所示。  

``` java
mListView.setEmptyView(findViewById(R.id.empty_view));
```

　　通过以上代码，就给ListView在空数据时显示了一张默认的图片，用来提示用户；而在有数据时，则不会显示。  

### ListView的滑动监听

　　ListView的滑动监听，是ListView中最重要的技巧，很多重写的ListView基本上都是在滑动事件的处理上下功夫，通过判断滑动事件进行不同的逻辑处理。而为了更佳精确地监听滑动事件，开发者通常还需要使用GestureDetector手势识别、VelocityTracker滑动速度检测等辅助类来完成更好的监听。这里介绍两种监听ListView滑动事件的方法，一个是通过OnTouchListener来实现监听，另一个是使用OnScrollListener来实现监听。  

#### OnTouchListener

　　OnTouchListener是View中的监听事件，通过监听ACTION_DOWN、ACTION_MOVE、ACTION_UP这三个事件发生时的坐标，就可以根据坐标判断用户滑动的方向，并在不同的事件中进行相应的逻辑处理，这种方式的使用代码如下所示。  

``` java
mListView.setOnTouchListener(new View.OnTouchListener() {
    @Override
    public boolean onTouch(View v, MotionEvent event) {
        switch (event.getAction()) {
            case MotionEvent.ACTION_DOWN:
                // 触摸时操作
                break;
            case MotionEvent.ACTION_MOVE:
                // 移动时操作
                break;
            case MotionEvent.ACTION_UP:
                // 离开时操作
                break;
        }
        return false;
    }
});
```

#### OnScrollListener

　　OnScrollListener是AbsListView中的监听事件，它封装了很多ListView相关的信息，使用起来也更加灵活。首先来看一下OnScrollListener的一般使用方法，代码如下所示。  

``` java
mListView.setOnScrollListener(new AbsListView.OnScrollListener() {
    @Override
    public void onScrollStateChanged(AbsListView view, int scrollState) {
        switch (scrollState) {
            case SCROLL_STATE_IDLE:
                // 滑动停止时
                Log.d("Test", "SCROLL_STATE_IDLE");
                break;
            case SCROLL_STATE_TOUCH_SCROLL:
                // 正在滚动
                Log.d("Test", "SCROLL_STATE_TOUCH_SCROLL");
                break;
            case SCROLL_STATE_FLING:
                // 手指抛动时，即手指用力滑动
                // 在离开后ListView由于惯性继续滑动
                Log.d("Test", "SCROLL_STATE_FLING");
                break;
        }
    }
    @Override
    public void onScroll(AbsListView view, int firstVisibleItem, int visibleItemCount, int totalItemCount) {
        // 滚动时一直调用
        Log.d("Test", "onScroll");
    }
```

　　OnScrollListener中有两个回调方法——onScrollStateChanged()和onScroll()。  

　　先来看第一个方法onScrollStateChanged()，这个方法根据它的参数scrollState来决定其回调的次数，scrollState有以下三种模式： 

- SCROLL_STATE_IDLE：滚动停止时。  
- SCROLL_STATE_TOUCH_SCROLL：正在滚动时。  
- SCROLL_STATE_FLING：手指抛动时，即手指用力滑动，在离开后ListView由于惯性继续滑动  

　　当用户没有做手指抛动的状态时，这个方法只会回调2次，否则会回调三次，差别就是手指抛动的这个状态。通常情况下，我们会在这个方法中通过不同的状态来设置一些标志Flag，来区分不同的滑动状态，供其他方法处理。  

　　下面再来看看onScroll()这个回调方法，它在ListView滚动时会一直回调，而方法中的后三个int类型的参数，则非常精确地显示了当前ListView滚动的状态，这三个参数如下所示。  

- firstVisibleItem：当前能看见的第一个Item的ID(从0开始)  
- visibleItemCount：当前能看见的Item的总数。  
- totalItemCount：整个ListView的Item总数。  

　　这里需要注意的是，当前能看见的Item数，包括没有显示完整的Item，即显示一小半的Item也包括在内了。通过这几个参数，可以很方便地进行一些判断，比如判断是否滚动到最后一行，就可以使用如下代码进行判断，当前可视的另一个Item的ID加上当前可视Item的和等于Item总数的时候，即滚动到了最后一行。  

``` java
if (firstVisibleItem + visibleItemCount == totalItemCount && totalItemCount > 0) {
    Log.d("Test", "滚动到最后一行");
}
```

　　再比如，可以通过如下代码来判断滚动的方向，代码如下所示。  

``` java
if(firstVisibleItem > lastVisibleItem){
    // 上滑
}else if(firstVisibleItem < lastVisibleItem){
    // 下滑
}
lastVisibleItem = firstVisibleItem;
```

　　通过一个成员变量lastVisibleItem来记录上次第一个可视的Item的ID并于当前的可视Item的ID进行比较，即可知道当前滚动的方向。  

　　要理解整个OnScrollListener，最好的方法还是在代码中添加Log，并打印出状态信息来进行分析学习。在以上代码中，已经添加了相应的Log，对照Log进行分析，会很快掌握OnScrollListener的用法。  

　　当然，ListView也给我们提供了一些封装的方法来获得当前可视的Item的位置等信息。  

``` java
// 获取可视区域内最后一个Item的id
mListView.getLastVisiblePosition();
// 获取可视区域内第一个Item的id
mListView.getFirstVisiblePosition();
```

项目地址→[ListViewSkill](https://github.com/Blankj/ListViewSkill)

* * *

原文地址[ListView常用优化技巧(Android群英传)][passage_url]  

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。  

[passage_url]: http://blankj.com/2016/07/24/ListView常用优化技巧(Android群英传)/