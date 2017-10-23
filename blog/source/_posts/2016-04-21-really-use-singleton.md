---
title: 你真的会写单例吗
date: 2016-04-21 10:08:35
categories:
  - Android
  - 设计模式
tags:
  - 设计模式
  - 单例
---

提到单例模式，相信都不会陌生，今天对其进行总结。

以下是单例模式的特点：

1. 单例类只能有一个实例。
2. 单例类必须自己自己创建自己的唯一实例。
3. 单例类必须给所有其他对象提供这一实例。

种类的话不好说有几类，因为要考虑到是否在多线程下运行，下面来介绍主要的几类：

<!-- more -->
### 懒汉类

```java
//懒汉式单例类.在第一次调用的时候实例化自己 
public class Singleton {
    private Singleton() {
    }

    private static Singleton single = null;

    //静态工厂方法 
    public static Singleton getInstance() {
        if (single == null) {
            single = new Singleton();
        }
        return single;
    }
}
```

懒汉么，所以在多线程下会失效，所以下面介绍三种懒汉的升级版来适应多线程

- 在getinstance前加上synchronized（同步），但这导致的是每次getInstance都会去同步，消耗资源。

```java
public class Singleton {
    private Singleton() {
    }

    private static Singleton single = null;

    // 静态工厂方法
    public static synchronized Singleton getInstance() {
        if (single == null) {
            single = new Singleton();
        }
        return single;
    }
}
```

- 双重检查锁，它是在以上做的修改，判断两次空，所以只有在第一次调用的时候会同步，避免了每次同步资源的消耗，注意 `volatile` 关键字。

```java
public class Singleton {
    private Singleton() {
    }

    private volatile static Singleton singleton = null; // 声明成 volatile 

    //静态工厂方法
    public static Singleton getInstance() {
        if (singleton == null) {
            synchronized (Singleton.class) {
                if (singleton == null) {
                    singleton = new Singleton();
                }
            }
        }
        return singleton;
    }
}
```

- 内部静态类，这种我觉得是最好的，既实现了线程安全，也避免了同步带来的性能影响。

```java
public class Singleton {
    private static class LazyHolder {
        private static final Singleton INSTANCE = new Singleton();
    }

    private Singleton() {
    }

    public static Singleton getInstance() {
        return LazyHolder.INSTANCE;
    }
}
```

### 饿汉类

饿汉式是典型的空间换时间，当类装载的时候就会创建类的实例，不管你用不用，先创建出来，然后每次调用的时候，就不需要再判断，节省了运行时间。

``` java
//饿汉式单例类.在类初始化时，已经自行实例化
public class Singleton {
    private Singleton() {
    }

    private static final Singleton single = new Singleton();

    //静态工厂方法
    public static Singleton getInstance() {
        return single;
    }
}
```

这种也是我比较喜欢的，因为简单易懂，但当实现了Serializable接口后，反序列化时单例会被破坏，实现Serializable接口需要重写readResolve，才能保证其反序列化依旧是单例：

``` java
private Object readResolve() throws ObjectStreamException { 
    return single; 
}  
```

### 枚举类

``` java
public enum Singleton {
    INSTANCE;

    public void whateverMethod() {
    }
}
```

这种方式是Effective Java作者Josh Bloch 提倡的方式，它不仅能避免多线程同步问题，而且还能防止反序列化重新创建新的对象，可谓是很坚强的壁垒啊，不过，个人认为由于1.5中才加入enum特性，用这种方式写不免让人感觉生疏，在实际工作中，我也很少看见有人这么写过。

以上就是常用的单例模式，一般的情况下，我会使用饿汉式，只有在要明确实现lazy loading效果时才会使用内部静态类，另外，如果涉及到反序列化创建对象时我会试着使用枚举的方式来实现单例，不过，我一直会保证我的程序是线程安全的。