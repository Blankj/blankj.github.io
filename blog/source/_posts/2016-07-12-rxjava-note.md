---
title: RxJava笔记
date: 2016-07-12 09:26:54
categories:
  - Android
  - 第三方库
tags:
  - 第三方库
  - RxJava
---

　　分享一篇入门RxJava很好的一篇文章[给 Android 开发者的 RxJava 详解](http://gank.io/post/560e15be2dca930e00da1083)，以下精髓摘自其中。

　　RxJava 有四个基本概念：Observable (可观察者，即被观察者)、 Observer (观察者)、 subscribe (订阅)、事件。Observable 和 Observer 通过 subscribe() 方法实现订阅关系，从而 Observable 可以在需要的时候发出事件来通知 Observer。

### 基本实现

　　RxJava 的基本实现主要有以下三点：
<!--more-->
1. 创建 Observer
可以是 `Observer` ，也可以是 `Subscriber` 。
2. 创建 Observable
`create()`: 是 RxJava 最基本的创造事件序列的方法。
`just(T...)`: 将传入的参数依次发送出来。
`from(T[]) / from(Iterable<? extends T>)`: 将传入的数组或 Iterable 拆分成具体对象后，依次发送出来。
3. Subscribe
`Observable.subscribe(Subscriber/Observer)`

### 线程控制

　　在不指定线程的情况下， RxJava 遵循的是线程不变的原则，即：在哪个线程调用 subscribe()，就在哪个线程生产事件；在哪个线程生产事件，就在哪个线程消费事件。如果需要切换线程，就需要用到 `Scheduler` （调度器）。

- `Schedulers.immediate()` : 直接在当前线程运行，相当于不指定线程。这是默认的 Scheduler 。
- `Schedulers.newThread()` : 总是启用新线程，并在新线程执行操作。
- `Schedulers.io()` : I/O 操作（读写文件、读写数据库、网络信息交互等）所使用的 Scheduler。行为模式和 newThread() 差不多，区别在于 io() 的内部实现是是用一个无数量上限的线程池，可以重用空闲的线程，因此多数情况下 io() 比 newThread() 更有效率。不要把计算工作放在 io() 中，可以避免创建不必要的线程。
- `Schedulers.computation()` : 计算所使用的 Scheduler。这个计算指的是 CPU 密集型计算，即不会被 I/O 等操作限制性能的操作，例如图形的计算。这个 Scheduler 使用的固定的线程池，大小为 CPU 核数。不要把 I/O 操作放在 computation() 中，否则 I/O 操作的等待时间会浪费 CPU。
- `Schedulers.trampoline()` :在当前线程中的工作放入队列中排队，并依次操作。
- `AndroidSchedulers.mainThread()` : 它指定的操作将在 Android 主线程运行。

　　有了这几个 Scheduler ，就可以使用 `subscribeOn()` 和 `observeOn()` 两个方法来对线程进行控制了。`subscribeOn()`: 指定 `subscribe()` 所发生的线程，即 `Observable.OnSubscribe` 被激活时所处的线程。或者叫做事件产生的线程。`observeOn()`: 指定 `Subscriber` 所运行在的线程。或者叫做事件消费的线程。

　　文字叙述总归难理解，上代码：

``` java
Observable.just(1, 2, 3, 4)
    .subscribeOn(Schedulers.io()) // 指定 subscribe() 发生在 IO 线程
    .observeOn(AndroidSchedulers.mainThread()) // 指定 Subscriber 的回调发生在主线程
    .subscribe(new Action1<Integer>() {
        @Override
        public void call(Integer number) {
            Log.d(tag, "number:" + number);
        }
    });
```

　　上面这段代码中，由于 `subscribeOn(Schedulers.io())` 的指定，被创建的事件的内容 1、2、3、4 将会在 IO 线程发出；而由于 `observeOn(AndroidScheculers.mainThread())` 的指定，因此 `subscriber` 数字的打印将发生在主线程 。事实上，这种在 `subscribe()` 之前写上两句 `subscribeOn(Scheduler.io())` 和 `observeOn(AndroidSchedulers.mainThread())` 的使用方式非常常见，它适用于多数的 『后台线程取数据，主线程显示』的程序策略。

　　当使用了多个 `subscribeOn()` 的时候，只有第一个 `subscribeOn()` 起作用，当使用多个 `observeOn()` 时候， `observeOn()` 指定的是它之后的操作所在的线程。

　　如果对准备工作的线程有要求，`onStart()` 就不适用了，因为它总是在 `subscribe` 所发生的线程被调用，而不能指定线程。要在指定的线程来做准备工作，可以使用 `doOnSubscribe()` 方法，`doOnSubscribe()` 之后有 `subscribeOn()` 的话，它将执行在离它最近的 `subscribeOn()` 所指定的线程。

### 变换

　　所谓变换，就是将事件序列中的对象或整个序列进行加工处理，转换成不同的事件或事件序列。
- map(): 事件对象的`一对一`的转化。

- flatMap(): 返回的是个 `Observable` 对象。