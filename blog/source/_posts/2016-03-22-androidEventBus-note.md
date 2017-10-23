---
title: AndroidEventBus笔记
date: 2016-03-22 16:03:07
categories:
  - Android
  - 第三方库
tags:
  - 第三方库
  - AndroidEventBus
---

EventBus最直接的好处就是解耦，但因为内部实现是反射，小项目无所谓，大项目的话性能不是很好。  

本次要介绍的是AndroidEventBus，而不是greenrobot的EventBus，其使用注解，使用方便，但效率比不上EventBus。订阅函数支持tag(类似广播接收器的Action)使得事件的投递更加准确，能适应更多使用场景。其github项目地址为：[androideventbus](https://github.com/hehonghui/AndroidEventBus/blob/master/README-ch.md)下面介绍其用法。  
<!--more-->
首先在build.gradle中加入依赖

```java
dependencies {
    compile 'org.simple:androideventbus:1.0.5.1'// 可以在官网查看最新版本
}
```
　　  
在Android中我们可以在onCreate方法中调用EventBus的register(Object subscriber) 注册订阅者，具体如下：  

```java
@Override
protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_main);
    tryBtn = (Button) findViewById(R.id.btn_main_try);
    showTV = (TextView) findViewById(R.id.tv_main_show);
    tryBtn.setOnClickListener(new View.OnClickListener() {
        @Override
        public void onClick(View v) {
            Intent intent = new Intent(getApplicationContext(), SecondActivity.class);
            startActivity(intent);
        }
    });

    //注册EventBus
    EventBus.getDefault().register(this);
}
```

注册者中通过Subscriber注解来标识事件接收对象中的接收方法。

```java
@Subscriber(tag = "my_tag", mode = ThreadMode.MAIN)
public void onEvent(FirstEvent event) {
    String msg = "onEvent收到了消息：" + event.getMsg();
    showTV.setText(msg);
    Toast.makeText(this, msg, Toast.LENGTH_LONG).show();
}
```

tag可用来区分post过来的消息标识，根据tag值可处理不同的消息。  

mode有三种。  

1. **ThreadMode.MAIN:**执行在UI线程。  
2. **ThreadMode.POST:**post函数在哪个线程执行,该函数就执行在哪个线程。  
3. **ThreadMode.ASYNC:**执行在一个独立的线程。  
　　  
调用EventBus的unregister(Object subscriber) 方法，取消注册的订阅者：

```java
@Override
protected void onDestroy() {
    super.onDestroy();
    //注销EventBus
    EventBus.getDefault().unregister(this);
}
```

以上三步一般都写在同一个类中。  

下面介绍怎么把消息传过来，很简单，只需在其他需要调用的地方post(Object event) 即可，消息便会进入队列。

```java
@Override
protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_second);
    firstEventBtn = (Button) findViewById(R.id.btn_second_first_event);
    firstEventBtn.setOnClickListener(new View.OnClickListener() {
        @Override
        public void onClick(View v) {
            //发送消息
            EventBus.getDefault().post(new FirstEvent("FirstEvent btn clicked"), "my_tag");
        }
    });
}
```

**需要强调的是：post的参数一定要是对象（引用类型），int、byte、long都是不可以的，实在需要就得转换为Integer、Byte和Long类型。**  
　　  
FirstEvent类

```java
public class FirstEvent {
    private String mMsg;
    public FirstEvent(String msg) {
        mMsg = msg;
    }
    public String getMsg(){
        return mMsg;
    }
}
```

好了，就是这么简单。