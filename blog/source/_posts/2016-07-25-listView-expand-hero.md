---
title: ListView常用拓展(Android群英传)
date: 2016-07-25 16:11:20
categories:
  - Android群英传
  - 4.ListView使用技巧
tags:
  - 4.ListView使用技巧
---

***内容是博主照着书敲出来的，博主码字挺辛苦的，转载请注明出处，后序内容陆续会码出。***  

　　ListView虽然使用广泛，但系统原生的ListView显然是不能满足用户在审美、功能上不断提高的需求。不过也不要紧，Android完全可以定制化，让我们非常方便地对原生ListView进行拓展、修改。于是，在开发者的创新下，ListView越来越丰富多彩，各种各样的基于原生ListView的拓展让人目不暇接。下面来看几个常用的ListView拓展。  
<!-- more -->
### 具有弹性的ListView

　　Android默认的ListView在滚动到顶端或者底端的时候，并没有很好的提示。在Android5.X中，Google为这样的行为只添加了一个半月形的阴影效果，如下图所示。  

{% asset_img top_shadow.jpg [顶部阴影效果] %}

　　而在iOS系统中，列表都是具有弹性效果的，即滚动到底端或者顶端后会继续往下或者往上滑动一段距离。不得不说，这样的设计的确更加的友好，虽然不知道Google为什么不模仿这样的设计，但我们可以自己修改ListView，让ListView也可以“弹性十足”。  

　　网上有很多通过重写ListView来实现弹性效果的方法，比如增加HeaderView或者使用ScrollView进行嵌套，方法有很多，不过这里可以使用一种非常简单的方法来实现这个效果。虽然不如那些方法可定制化高、效果丰富，但主要目的是让读者朋友们学会如何从源代码中找到问题的解决办法。  

　　我们在查看ListView源代码的时候可以发现，ListView中有一个控制滑动到边缘的处理方法，如下所示。  

``` java
protected boolean overScrollBy(int deltaX, int deltaY,
                               int scrollX, int scrollY,
                               int scrollRangeX, int scrollRangeY,
                               int maxOverScrollX, int maxOverScrollY,
                               boolean isTouchEvent)
```

　　可以看见这样一个参数：maxOverScrollY，注释中这样写道——Number of pixels to overscroll by in either direction along the Y axis。由此可以发现，虽然它的默认值是0，但其实只要修改这个参数的值，就可以让ListView具有弹性了！所以，既然我们不知道为什么Google不采用这样的修改，那我们就自己来修改一下吧。重写这个方法，并将maxOverScrollY改为设置的值——mMaxOverDistance，代码如下所示。  

``` java
@Override
protected boolean overScrollBy(int deltaX, int deltaY, int scrollX, int scrollY, int scrollRangeX, int scrollRangeY, int maxOverScrollX, int maxOverScrollY, boolean isTouchEvent) {
    return super.overScrollBy(deltaX, deltaY, scrollX, scrollY, scrollRangeX, scrollRangeY, maxOverScrollX, mMaxOverDistance, isTouchEvent);
}
```

　　这样，通过对这个值得修改，就实现了一个具有弹性的ListView了。效果如下图所示。  

{% asset_img flexible_ls.jpg [弹性ListView效果] %}

　　当然，为了能够满足多分辨率的需求，我们可以在修改maxOverScrollY值的时候，可以通过屏幕的density来计算具体的值，让不同分辨率的弹性距离基本一致，代码如下所示。  

``` java
private void initView() {
    DisplayMetrics metrics = mContext.getResources().getDisplayMetrics();
    float density = metrics.density;
    mMaxOverDistance = (int) (density * mMaxOverDistance);
}
```

### 自动显示、隐藏布局的ListView

　　相信通过Google+的朋友应该非常熟悉这样一个效果：当我们在ListView上滑动的时候，顶部的ActionBar或者Toolbar就会相应的隐藏或者显示。这样的效果一出现，各种App竞相模仿，不得不说，Google的应用一直都是Android设计的风向标。  
　　大家可以发现，在滚动前界面上加载了上方的标题栏和右下角的悬浮编辑按钮，如下图所示。  

{% asset_img before_scroll_ui.png [滚动前界面] %}

　　当用户向下滚动时，标题栏和悬浮按钮消失了，让用户有更大的空间去阅读，如下图所示。  

{% asset_img after_scroll_ui.png [滚动后界面] %}

　　下面我们就来仿照这个例子设计一个类似的效果。  

　　我们知道，让一个布局显示或者隐藏并带有动画效果，可以通过属性动画来很方便地实现，所以这个效果的关键就在于如何获得ListView的各种滑动事件。所以借助View的OnTouchListener接口来监听ListView的滑动，通过比较与上次坐标的大小，来判断滑动的方向，并通过滑动的方向来判断是否需要显示或隐藏对应的布局。在开始判断滑动事件之前，我们还要做一些准备工作，首先需要给ListView增加一个HeaderView，避免第一个Item被Toolbar遮挡，代码如下所示。  

``` java
View header = new View(this);
header.setLayoutParams(new AbsListView.LayoutParams(ViewGroup.LayoutParams.MATCH_PARENT,
        (int) getResources().getDimension(R.dimen.abc_action_bar_default_height_material)));
mListView.addHeaderView(header);
```

　　在代码中，通过使用abc_action_bar_default_height_material属性获取系统Actionbar的高度，并设置给HeaderView。另外，定义一个mTouchSlop变量来获取系统认为的最低滑动距离，即超过这个距离的移动，系统就将其定义为滑动状态了，对这个值得获取非常简单，代码如下所示。  

``` java
mTouchSlop = ViewConfiguration.get(this).getScaledTouchSlop();
```

　　有了前面的准备工作，下面我们就可以判断滑动的事件了，关键代码如下所示。  

``` java
@Override
public boolean onTouch(View v, MotionEvent event) {
    switch (event.getAction()) {
        case MotionEvent.ACTION_DOWN:
            mFirstY = event.getY();
            break;
        case MotionEvent.ACTION_MOVE:
            mCurrentY = event.getY();
            if (mCurrentY - mFirstY > mTouchSlop) {
                // down
                if (mShow) {
                    toolbarAnim(0);
                }
                mShow = !mShow;
            } else if (mCurrentY - mFirstY < mTouchSlop) {
                // up
                if (mShow) {
                    toolbarAnim(1);
                }
                mShow = !mShow;
            }
            break;
        case MotionEvent.ACTION_UP:
            break;
    }
    return false;
```

　　代码逻辑非常简单，只是通过滑动点的坐标改变大小，来判断移动的方向，并根据移动方向来执行不同的动画效果。  

　　有了前面的分析，实现这样一个效果就非常简单了，最后加上控制布局显示隐藏的动画，如下所示。  

``` java
private void toolbarAnim(int flag) {
    if (mAnimator != null && mAnimator.isRunning()) {
        mAnimator.cancel();
    }
    if (flag == 0) {
        mAnimator = ObjectAnimator.ofFloat(mToolbar,
                "translationY", mToolbar.getTranslationY(), 0);
    } else {
        mAnimator = ObjectAnimator.ofFloat(mToolbar,
                "translationY", mToolbar.getTranslationY(),
                -mToolbar.getHeight());
    }
    mAnimator.start();
}
```

　　动画也是最简单的位移属性动画。不过这里需要说一点题外话，这里使用了Toorbar这样一个新控件，Google已经推荐它来逐渐取代ActionBar了，因为它更加灵活。但是在使用的时候，一定要注意使用的theme一定是要NoActionBar的，不然会引起冲突。同时，不要忘记引入编译，代码如下。  

``` java
dependencies {
    compile fileTree(dir: 'libs', include: ['*.jar'])
    compile 'com.android.support:appcompat-v7:23.4.0'
}
```

　　运行程序后初始状态如下图所示，Toorbar显示在最上方。  

{% asset_img show_toolbar.png [Toolbar显示] %}

　　当向上滑动时，Toolbar隐藏，如下图所示。  

{% asset_img hide_toolbar.png [Toolbar隐藏] %}

　　再向下滑动时，Toolbar显示。  

### 聊天ListView

　　通常我们使用的ListView的每一项都具有相同的布局，所以展现出来的时候，除了数据不同，只要你不隐藏布局，其他的布局应该都是类似的。而我们熟知的QQ、微信等聊天App，在聊天界面，会展示至少两种布局，即收到的消息和自己发送的消息，其实这样的效果也是通过ListView来实现的，下面我们就来模仿一个聊天软件的聊天列表界面，其效果如下图所示。  

{% asset_img chat_lv.jpg [聊天界面ListView] %}

　　这样一个ListView与我们平时所使用的ListView最大的不同，就是它拥有两个不同的布局——收到的布局和发送的布局。要实现这样的效果，就需要拿ListView的Adapter“开刀”。  

　　在定义BaseAdapter的时候，需要去重写它的getView()方法，这个方法就是用来获取布局的，那么只需要在获取布局的时候，判断一下该获取哪一种布局就可以了。而且，ListView在设计的时候就已经考虑到了这种情况，所以它提供了两个方法，代码如下所示。  

``` java
@Override
public int getItemViewType(int position) {
    return super.getItemViewType(position);
}
@Override
public int getViewTypeCount() {
    return super.getViewTypeCount();
}
```

　　getItemViewType()方法用来返回第position个Item是何种类型，而getViewTypeCount()方法用来返回不同布局的总数。通过这两个方法，再结合getView()方法，就可以很轻松地设计出上面的聊天布局了。  

　　首先来实现两个布局——chat_item_itemin和chat_item_itemout。布局大同小异，只是方向上有区别。需要注意的是，显示聊天信息内容的TextView使用了9patch的图片，这种图片格式是Android中用来拉伸图片的，你可以把它想象成在某些方向上拉伸却不会失真、形变的图片就可以了，布局代码如下所示。由于in和out界面内容只是方向上的区别，这里只贴出一个布局的代码。  

``` java
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:gravity="center_vertical"
    android:orientation="horizontal"
    android:padding="10dp">
 
    <ImageView
        android:id="@+id/icon_in"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:src="@mipmap/ic_launcher" />
 
    <TextView
        android:id="@+id/text_in"
        android:layout_width="wrap_content"
        android:layout_height="wrap_content"
        android:background="@drawable/chatitem_in_bg"
        android:gravity="center"
        android:textSize="20sp" />
 
</LinearLayout>
```

　　同时，为了封装下聊天内容，便于在Adapter中获取数据信息，我们封装了一个Bean来保存聊天信息，代码如下所示。  

``` java
import android.graphics.Bitmap;
 
/*********************************************
* author: Blankj on 2016/7/25 14:01
 * blog:   http://blankj.com
* e-mail: blankj@qq.com
*********************************************/
public class ChatListViewBean {
 
    private int type;
    private String text;
    private Bitmap icon;
 
    public ChatListViewBean() {
    }
 
    public int getType() {
        return type;
    }
 
    public void setType(int type) {
        this.type = type;
    }
 
    public String getText() {
        return text;
    }
 
    public void setText(String text) {
        this.text = text;
    }
 
    public Bitmap getIcon() {
        return icon;
    }
 
    public void setIcon(Bitmap icon) {
        this.icon = icon;
    }
}
```

　　非常简单，我们只是声明了需要的信息并提供了get和set方法。  

　　接下来，需要来完成最重要的BaseAdapter了，同样使用ViewHolder模式来提高ListView的效率，并在getView()方法中进行布局类型的判断，从而确定使用哪种布局，代码如下所示。  

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
 * author:  Blankj on 2016/7/25 14:01
 * blog:    http://blankj.com
 * e-mail:  blankj@qq.com
*********************************************/
public class ChatListViewAdapter extends BaseAdapter {
 
    private List<ChatListViewBean> mData;
    private LayoutInflater mInflater;
 
    public ChatListViewAdapter(Context context, List<ChatListViewBean> data) {
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
    public int getItemViewType(int position) {
        ChatListViewBean bean = mData.get(position);
        return bean.getType();
    }
 
    @Override
    public int getViewTypeCount() {
        return 2;
    }
 
    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        ViewHolder viewHolder = null;
        if (convertView == null) {
            viewHolder = new ViewHolder();
            if (getItemViewType(position) == 0) {
                convertView = mInflater.inflate(R.layout.chat_item_itemin, null);
                viewHolder.icon = (ImageView) convertView.findViewById(R.id.icon_in);
                viewHolder.text = (TextView) convertView.findViewById(R.id.text_in);
            } else {
                convertView = mInflater.inflate(R.layout.chat_item_itemout, null);
                viewHolder.icon = (ImageView) convertView.findViewById(R.id.icon_out);
                viewHolder.text = (TextView) convertView.findViewById(R.id.text_out);
            }
            convertView.setTag(viewHolder);
        } else {
            viewHolder = (ViewHolder) convertView.getTag();
        }
        viewHolder.icon.setImageBitmap(mData.get(position).getIcon());
        viewHolder.text.setText(mData.get(position).getText());
        return convertView;
    }
 
    public final class ViewHolder {
        public ImageView icon;
        public TextView text;
    }
}
```

　　在以上代码中，通过在getView()中判断getItemType(position)的值来决定具体实例化哪一个布局，从而实现在一个ListView中多个布局内容的添加。最后，在测试的Activity里面添加了一些测试代码，来测试这个布局。  

``` java
import android.app.Activity;
import android.graphics.BitmapFactory;
import android.os.Bundle;
import android.widget.ListView;
 
import java.util.ArrayList;
import java.util.List;
 
/*********************************************
* author: Blankj on 2016/7/25 13:30
 * blog:   http://blankj.com
* e-mail: blankj@qq.com
*********************************************/
public class ChatListViewTest extends Activity {
 
    private ListView mListView;
 
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_chat);
        mListView = (ListView) findViewById(R.id.lv_chat);
 
        ChatListViewBean bean1 = new ChatListViewBean();
        bean1.setType(0);
        bean1.setIcon(BitmapFactory.decodeResource(getResources(),
                R.drawable.in_icon));
        bean1.setText("Hello how are you?");
 
        ChatListViewBean bean2 = new ChatListViewBean();
        bean2.setType(1);
        bean2.setIcon(BitmapFactory.decodeResource(getResources(),
                R.drawable.out_icon));
        bean2.setText("Fine thank you, and you?");
 
        ChatListViewBean bean3 = new ChatListViewBean();
        bean3.setType(0);
        bean3.setIcon(BitmapFactory.decodeResource(getResources(),
                R.drawable.in_icon));
        bean3.setText("I am fine, too");
 
        ChatListViewBean bean4 = new ChatListViewBean();
        bean4.setType(1);
        bean4.setIcon(BitmapFactory.decodeResource(getResources(),
                R.drawable.out_icon));
        bean4.setText("Bye bye");
 
        ChatListViewBean bean5 = new ChatListViewBean();
        bean5.setType(0);
        bean5.setIcon(BitmapFactory.decodeResource(getResources(),
                R.drawable.in_icon));
        bean5.setText("See you");
 
        List<ChatListViewBean> data = new ArrayList<ChatListViewBean>();
        data.add(bean1);
        data.add(bean2);
        data.add(bean3);
        data.add(bean4);
        data.add(bean5);
        mListView.setAdapter(new ChatListViewAdapter(this, data));
    }
}
```

　　在测试代码中，简单地添加了一些模拟的聊天内容，并将信息封装到设置的Bean对象中，最后运行程序，即可得到之前所示的聊天效果界面。  

### 动态改变ListView布局

　　通常情况下，如果要动态地改变点击Item的布局来达到一个Focus的效果，一般有两种方法。一种是将两种布局写在一起，通过控制布局的显示、隐藏，来达到切换布局的效果；另一种则是在getView()的时候，通过判断来选择加载不同的布局。两种方法各有利弊，关键还是看使用的场合。下面就以第二种方式，来演示一下这样的效果，程序运行后初始效果下图所示，第一个Item为默认Focus状态。  

{% asset_img app_init_status.jpg [程序初始状态] %}

　　当点击其他Item的时候，点击的Item变为Focus状态，其他Item还原，效果如下图所示。  

{% asset_img focus_change.jpg [Focus改变] %}

　　该效果实现的关键还是在于BaseAdapter。在这个实例中，通过如下所示的两个方法来给Item设置两种不同的布局——Focus和Normal。  

``` java
private View addFocusView(int i) {
    ImageView iv = new ImageView(mContext);
    iv.setImageResource(R.mipmap.ic_launcher);
    return iv;
}

private View addNormalView(int i) {
    LinearLayout layout = new LinearLayout(mContext);
    layout.setOrientation(LinearLayout.HORIZONTAL);
    ImageView iv = new ImageView(mContext);
    iv.setImageResource(R.drawable.in_icon);
    layout.addView(iv, new LinearLayout.LayoutParams(
            LinearLayout.LayoutParams.WRAP_CONTENT,
            LinearLayout.LayoutParams.WRAP_CONTENT));
    TextView tv = new TextView(mContext);
    tv.setText(mData.get(i));
    layout.addView(tv, new LinearLayout.LayoutParams(
            LinearLayout.LayoutParams.WRAP_CONTENT,
            LinearLayout.LayoutParams.WRAP_CONTENT));
    layout.setGravity(Gravity.CENTER);
    return layout;
}
```

　　在这两个方法中，可以根据Item位置的不同来设置不同的显示图片等信息，但这里为了方便，就统一只显示一张图片。  

　　下面回到BaseAdapter，在getView()方法中，通过判断点击的位置来改变相应的视图，代码如下所示。  

``` java
@Override
public View getView(int position, View convertView, ViewGroup parent) {
    LinearLayout layout = new LinearLayout(mContext);
    layout.setOrientation(LinearLayout.VERTICAL);
    if (mCurrentItem == position) {
        layout.addView(addFocusView(position));
    } else {
        layout.addView(addNormalView(position));
    }
    return layout;
}
```

　　在以上代码中，通过判断当前CurrentItem是否是点击的那个position，就可以动态控制显示的布局了。当然，仅仅这样是不够的，因为getView()是在初始化的时候调用，后面再点击Item的时候，并没有再次调用getView()。所以，必须要让ListView在点击后，再刷新一次。于是我们请出了notifyDataSetChanged()方法来帮助实现刷新布局的功能，代码如下所示。  

``` java
mListView.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> parent, View view,
                                    int position, long id) {
                mAdapter.setCurrentItem(position);
                mAdapter.notifyDataSetChanged();
            }
        });
```

项目地址→[ListViewExpandation](https://github.com/Blankj/ListViewExpandation)

* * *

原文地址[ListView常用拓展(Android群英传)][passage_url]  

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。  

[passage_url]: http://blankj.com/2016/07/25/ListView常用拓展(Android群英传)/