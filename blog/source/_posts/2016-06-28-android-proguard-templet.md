---
title: Android Studio混淆模板及常用第三方混淆(看了都说好)
date: 2016-06-28 17:36:46
categories:
  - Android
  - Android Tips
tags:
  - Android Tips
  - 混淆
---

　　首先要在build.gradle中开启混淆，也就是 `minifyEnabled true` ，我用的build.gradle具体如下所示：

``` groovy
def releaseTime() {
    return new Date().format("yyyy.MM.dd", TimeZone.getTimeZone("UTC"))
}

android {
    ....

    buildTypes {
        release {
            // 混淆
            minifyEnabled true
            // Zipalign优化
            zipAlignEnabled true
            // 移除无用的resource文件
            shrinkResources true
            // 前一部分代表系统默认的android程序的混淆文件，该文件已经包含了基本的混淆声明，后一个文件是自己的定义混淆文件
            proguardFiles getDefaultProguardFile('proguard-android.txt'), 'proguard-rules.pro'
 
            applicationVariants.all { variant ->
                variant.outputs.each { output ->
                    def outputFile = output.outputFile
                    if (outputFile != null && outputFile.name.endsWith('.apk')) {
                        // 输出apk名称为appname-1.0-beta1-xxxx-xx-xx.apk
                        def fileName = "app-${defaultConfig.versionName}-beta1-${releaseTime()}.apk"
                        output.outputFile = new File(outputFile.parent, fileName)
                    }
                }
            }
        }
    }

    ....
}
```
<!--more-->
　　上面注释中写到proguard-android.txt文件，这是系统默认的混淆文件，具体在../sdk/tools/proguard/目录下，其中包含了android最基本的混淆，一般不需要改动，我们需要配置的是项目中app下的proguard-rules.pro文件，在此提供一份混淆模板。  

```
#############################################
#
# 对于一些基本指令的添加
#
#############################################
# 代码混淆压缩比，在0~7之间，默认为5，一般不做修改
-optimizationpasses 5
 
# 混合时不使用大小写混合，混合后的类名为小写
-dontusemixedcaseclassnames
 
# 指定不去忽略非公共库的类
-dontskipnonpubliclibraryclasses
 
# 这句话能够使我们的项目混淆后产生映射文件
# 包含有类名->混淆后类名的映射关系
-verbose
 
# 指定不去忽略非公共库的类成员
-dontskipnonpubliclibraryclassmembers
 
# 不做预校验，preverify是proguard的四个步骤之一，Android不需要preverify，去掉这一步能够加快混淆速度。
-dontpreverify
 
# 保留Annotation不混淆
-keepattributes *Annotation*,InnerClasses
 
# 避免混淆泛型
-keepattributes Signature
 
# 抛出异常时保留代码行号
-keepattributes SourceFile,LineNumberTable
 
# 指定混淆是采用的算法，后面的参数是一个过滤器
# 这个过滤器是谷歌推荐的算法，一般不做更改
-optimizations !code/simplification/cast,!field/*,!class/merging/*
 
 
#############################################
#
# Android开发中一些需要保留的公共部分
#
#############################################
 
# 保留我们使用的四大组件，自定义的Application等等这些类不被混淆
# 因为这些子类都有可能被外部调用
-keep public class * extends android.app.Activity
-keep public class * extends android.app.Appliction
-keep public class * extends android.app.Service
-keep public class * extends android.content.BroadcastReceiver
-keep public class * extends android.content.ContentProvider
-keep public class * extends android.app.backup.BackupAgentHelper
-keep public class * extends android.preference.Preference
-keep public class * extends android.view.View
-keep public class com.android.vending.licensing.ILicensingService
 
 
# 保留support下的所有类及其内部类
-keep class android.support.** {*;}
 
# 保留继承的
-keep public class * extends android.support.v4.**
-keep public class * extends android.support.v7.**
-keep public class * extends android.support.annotation.**
 
# 保留R下面的资源
-keep class **.R$* {*;}
 
# 保留本地native方法不被混淆
-keepclasseswithmembernames class * {
    native <methods>;
}
 
# 保留在Activity中的方法参数是view的方法，
# 这样以来我们在layout中写的onClick就不会被影响
-keepclassmembers class * extends android.app.Activity{
    public void *(android.view.View);
}
 
# 保留枚举类不被混淆
-keepclassmembers enum * {
    public static **[] values();
    public static ** valueOf(java.lang.String);
}
 
# 保留我们自定义控件（继承自View）不被混淆
-keep public class * extends android.view.View{
    *** get*();
    void set*(***);
    public <init>(android.content.Context);
    public <init>(android.content.Context, android.util.AttributeSet);
    public <init>(android.content.Context, android.util.AttributeSet, int);
}
 
# 保留Parcelable序列化类不被混淆
-keep class * implements android.os.Parcelable {
    public static final android.os.Parcelable$Creator *;
}
 
# 保留Serializable序列化的类不被混淆
-keepclassmembers class * implements java.io.Serializable {
    static final long serialVersionUID;
    private static final java.io.ObjectStreamField[] serialPersistentFields;
    !static !transient <fields>;
    !private <fields>;
    !private <methods>;
    private void writeObject(java.io.ObjectOutputStream);
    private void readObject(java.io.ObjectInputStream);
    java.lang.Object writeReplace();
    java.lang.Object readResolve();
}
 
# 对于带有回调函数的onXXEvent、**On*Listener的，不能被混淆
-keepclassmembers class * {
    void *(**On*Event);
    void *(**On*Listener);
}
 
# webView处理，项目中没有使用到webView忽略即可
-keepclassmembers class fqcn.of.javascript.interface.for.webview {
    public *;
}
-keepclassmembers class * extends android.webkit.webViewClient {
    public void *(android.webkit.WebView, java.lang.String, android.graphics.Bitmap);
    public boolean *(android.webkit.WebView, java.lang.String);
}
-keepclassmembers class * extends android.webkit.webViewClient {
    public void *(android.webkit.webView, jav.lang.String);
}
 
# 移除Log类打印各个等级日志的代码，打正式包的时候可以做为禁log使用，这里可以作为禁止log打印的功能使用
# 记得proguard-android.txt中一定不要加-dontoptimize才起作用
# 另外的一种实现方案是通过BuildConfig.DEBUG的变量来控制
#-assumenosideeffects class android.util.Log {
#    public static int v(...);
#    public static int i(...);
#    public static int w(...);
#    public static int d(...);
#    public static int e(...);
#}
 
#############################################
#
# 项目中特殊处理部分
#
#############################################
 
#-----------处理反射类---------------
 
 
 
#-----------处理js交互---------------
 
 
 
#-----------处理实体类---------------
# 在开发的时候我们可以将所有的实体类放在一个包内，这样我们写一次混淆就行了。
#-keep public class com.ljd.example.entity.** {
#    public void set*(***);
#    public *** get*();
#    public *** is*();
#}
 
 
#-----------处理第三方依赖库---------






```

　　下面介绍常用第三方混淆配置（已按字母排序，不定期更新）：  

```
# AndroidEventBus
-keep class org.simple.** { *; }
-keep interface org.simple.** { *; }
-keepclassmembers class * {
    @org.simple.eventbus.Subscriber <methods>;
}


# 百度地图（jar包换成自己的版本，记得签名要匹配）
-libraryjars libs/baidumapapi_v2_1_3.jar
-keep class com.baidu.** {*;}
-keep class vi.com.** {*;}
-keep class com.sinovoice.** {*;}
-keep class pvi.com.** {*;}
-dontwarn com.baidu.**
-dontwarn vi.com.**
-dontwarn pvi.com.**


# Bugly
-dontwarn com.tencent.bugly.**
-keep class com.tencent.bugly.** {*;}


# ButterKnife
-keep class butterknife.** { *; }
-dontwarn butterknife.internal.**
-keep class **$$ViewBinder { *; } 
-keepclasseswithmembernames class * {
    @butterknife.* <fields>;
}
-keepclasseswithmembernames class * {
    @butterknife.* <methods>;
}


# EventBus
-keepattributes *Annotation*
-keepclassmembers class ** {
    @org.greenrobot.eventbus.Subscribe <methods>;
}
-keep enum org.greenrobot.eventbus.ThreadMode { *; }


# Facebook
-keep class com.facebook.** {*;}
-keep interface com.facebook.** {*;}
-keep enum com.facebook.** {*;}


# FastJson
-dontwarn com.alibaba.fastjson.**
-keep class com.alibaba.fastjson.** { *; }
-keepattributes Signature
-keepattributes *Annotation*


# Fresco
-keep class com.facebook.fresco.** {*;}
-keep interface com.facebook.fresco.** {*;}
-keep enum com.facebook.fresco.** {*;}


# 高德相关依赖
# 集合包:3D地图3.3.2 导航1.8.0 定位2.5.0
-dontwarn com.amap.api.**
-dontwarn com.autonavi.**
-keep class com.amap.api.**{*;}
-keep class com.autonavi.**{*;}
# 地图服务
-dontwarn com.amap.api.services.**
-keep class com.map.api.services.** {*;}
# 3D地图
-dontwarn com.amap.api.mapcore.**
-dontwarn com.amap.api.maps.**
-dontwarn com.autonavi.amap.mapcore.**
-keep class com.amap.api.mapcore.**{*;}
-keep class com.amap.api.maps.**{*;}
-keep class com.autonavi.amap.mapcore.**{*;}
# 定位
-dontwarn com.amap.api.location.**
-dontwarn com.aps.**
-keep class com.amap.api.location.**{*;}
-keep class com.aps.**{*;}
# 导航
-dontwarn com.amap.api.navi.**
-dontwarn com.autonavi.**
-keep class com.amap.api.navi.** {*;}
-keep class com.autonavi.** {*;}


# Glide
-keep public class * implements com.bumptech.glide.module.GlideModule
-keep public enum com.bumptech.glide.load.resource.bitmap.ImageHeaderParser$** {
  **[] $VALUES;
  public *;
}


# Gson
-keepattributes Signature-keepattributes *Annotation*
-keep class sun.misc.Unsafe { *; }
-keep class com.google.gson.stream.** { *; }
# 使用Gson时需要配置Gson的解析对象及变量都不混淆。不然Gson会找不到变量。
# 将下面替换成自己的实体类
-keep class com.example.bean.** { *; }


# Jackson
-dontwarn org.codehaus.jackson.**
-dontwarn com.fasterxml.jackson.databind.**
-keep class org.codehaus.jackson.** { *;}
-keep class com.fasterxml.jackson.** { *; }


# 极光推送
-dontoptimize
-dontpreverify
-dontwarn cn.jpush.**
-keep class cn.jpush.** { *; }


# OkHttp
-keep class com.squareup.okhttp.** { *; }
-keep interface com.squareup.okhttp.** { *; }
-dontwarn com.squareup.okhttp.**


# OkHttp3
-dontwarn com.squareup.okhttp3.**
-keep class com.squareup.okhttp3.** { *;}
-dontwarn okio.**


# Okio
-dontwarn com.squareup.**  
-dontwarn okio.**  
-keep public class org.codehaus.* { *; }  
-keep public class java.nio.* { *; }


# OrmLite
-keepattributes *DatabaseField* 
-keepattributes *DatabaseTable* 
-keepattributes *SerializedName*  
-keep class com.j256.**
-keepclassmembers class com.j256.** { *; }
-keep enum com.j256.**
-keepclassmembers enum com.j256.** { *; }
-keep interface com.j256.**
-keepclassmembers interface com.j256.** { *; }


# Realm
-keep class io.realm.annotations.RealmModule
-keep @io.realm.annotations.RealmModule class *
-keep class io.realm.internal.Keep
-keep @io.realm.internal.Keep class * { *; }
-dontwarn javax.**
-dontwarn io.realm.**


# Retrofit
-dontwarn retrofit2.**
-keep class retrofit2.** { *; }
-keepattributes Signature
-keepattributes Exceptions


# Retrolambda
-dontwarn java.lang.invoke.*


# RxJava RxAndroid
-dontwarn sun.misc.**
-keepclassmembers class rx.internal.util.unsafe.*ArrayQueue*Field* {
    long producerIndex;
    long consumerIndex;
}
-keepclassmembers class rx.internal.util.unsafe.BaseLinkedQueueProducerNodeRef {
    rx.internal.util.atomic.LinkedQueueNode producerNode;
}
-keepclassmembers class rx.internal.util.unsafe.BaseLinkedQueueConsumerNodeRef {
    rx.internal.util.atomic.LinkedQueueNode consumerNode;
}


# Universal-Image-Loader-v1.9.5
-libraryjars libs/universal-image-loader-1.9.5-SNAPSHOT-with-sources.jar
-dontwarn com.nostra13.universalimageloader.**
-keep class com.nostra13.universalimageloader.** { *; }


# 微信支付
-dontwarn com.tencent.mm.**
-dontwarn com.tencent.wxop.stat.**
-keep class com.tencent.mm.** {*;}
-keep class com.tencent.wxop.stat.**{*;}


# 信鸽
-keep public class * extends android.app.Service
-keep public class * extends android.content.BroadcastReceiver
-keep class com.tencent.android.tpush.**  {* ;}
-keep class com.tencent.mid.**  {* ;}
-keepattributes *Annotation*


# 新浪微博
-keep class com.sina.weibo.sdk.* { *; }  
-keep class android.support.v4.* { *; }  
-keep class com.tencent.* { *; }  
-keep class com.baidu.* { *; }  
-keep class lombok.ast.ecj.* { *; }  
-dontwarn android.support.v4.**  
-dontwarn com.tencent.**s  
-dontwarn com.baidu.**  


# 讯飞语音
-dontwarn com.iflytek.**
-keep class com.iflytek.** {*;}


# xUtils3.0
-keepattributes Signature,Annotation
-keep public class org.xutils.** {
public protected *;
}
-keep public interface org.xutils.** {
public protected *;
}
-keepclassmembers class * extends org.xutils.** {
public protected *;
}
-keepclassmembers @org.xutils.db.annotation.* class * {;}
-keepclassmembers @org.xutils.http.annotation. class * {*;}
-keepclassmembers class * {
@org.xutils.view.annotation.Event ;
}


# 银联
-dontwarn com.unionpay.**
-keep class com.unionpay.** { *; }


# 友盟统计分析
-keepclassmembers class * { public <init>(org.json.JSONObject); }
-keepclassmembers enum com.umeng.analytics.** {
    public static **[] values();
    public static ** valueOf(java.lang.String);
}


# 友盟自动更新
-keepclassmembers class * { public <init>(org.json.JSONObject); }
-keep public class cn.irains.parking.cloud.pub.R$*{ public static final int *; }
-keep public class * extends com.umeng.**
-keep class com.umeng.** { *; }


# 支付宝钱包
-dontwarn com.alipay.**
-dontwarn HttpUtils.HttpFetcher
-dontwarn com.ta.utdid2.**
-dontwarn com.ut.device.**
-keep class com.alipay.android.app.IAlixPay{*;}
-keep class com.alipay.android.app.IAlixPay$Stub{*;}
-keep class com.alipay.android.app.IRemoteServiceCallback{*;}
-keep class com.alipay.android.app.IRemoteServiceCallback$Stub{*;}
-keep class com.alipay.sdk.app.PayTask{ public *;}
-keep class com.alipay.sdk.app.AuthTask{ public *;}
-keep class com.alipay.mobilesecuritysdk.*
-keep class com.ut.*
```

* * *

原文地址[Android Studio混淆模板及常用第三方混淆(看了都说好)][passage_url]

我的自媒体博客[Blankj小站](http://blankj.com/)，欢迎来逛逛。

[passage_url]: http://blankj.com/2016/06/28/Android-Studio混淆模板及常用第三方混淆(看了都说好)/