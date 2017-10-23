---
title: Android开发之版本统一规范
date: 2016-09-21 19:12:09
categories:
  - Android
  - Android Tips
tags:
  - Android Tips
---

Android开发存在着众多版本的不同，比如compileSdkVersion、minSdkVersion、targetSdkVersion以及项目中依赖第三方库的版本，不同的module及不同的开发人员都有不同的版本，所以需要一个统一版本规范的文件，现在我就来介绍一种方式：配置config.gradle文件。
在项目根目录，也就是跟app同一目录下的地方新建config.gradle文件，如下图所示

{% asset_img build_config.png [config.gradle] %}
<!-- more -->
如果了解groovy的话，那么阅读以下代码肯定是小菜一碟了，不了解的话，看一下之后其实也很好懂。

``` groovy
ext {
    signingConfig = [
            storePassword: "xxx",
            keyAlias     : "xxx",
            keyPassword  : "xxx"
    ]

    android = [
            compileSdkVersion: 24,
            buildToolsVersion: "24.0.2",
            minSdkVersion    : 17,
            targetSdkVersion : 24,
            versionName      : "1.0.0",
            versionCode      : 1
    ]

    depsVersion = [
            support    : "24.2.1",
            retrofit   : "2.1.0",
            okhttp     : "3.3.1",
            agera      : "1.1.0",
            butterKnife: "8.4.0",
    ]

    deps = [
            // ------------- Android -------------
            supportV4          : "com.android.support:support-v4:${depsVersion.support}",
            supportV7          : "com.android.support:appcompat-v7:${depsVersion.support}",
            recyclerviewV7     : "com.android.support:recyclerview-v7:${depsVersion.support}",
            cardviewV7         : "com.android.support:cardview-v7:${depsVersion.support}",
            design             : "com.android.support:design:${depsVersion.support}",
            percent            : "com.android.support:percent:${depsVersion.support}",

            // ------------- Three Parts -------------
            butterknife        : "com.jakewharton:butterknife:${depsVersion.butterKnife}",
            butterknifeCompiler: "com.jakewharton:butterknife-compiler:${depsVersion.butterKnife}",

            rxandroid          : "io.reactivex:rxandroid:1.2.1",
            rxjava             : "io.reactivex:rxjava:1.1.6",

            retrofit           : "com.squareup.retrofit2:retrofit:2.1.0",

            okhttp             : "com.squareup.okhttp3:okhttp:3.4.1",

            androideventbus    : "org.simple:androideventbus:1.0.5.1",

            picasso            : "com.squareup.picasso:picasso:2.5.2",

            fresco             : "com.facebook.fresco:fresco:0.13.0",

            gson               : "com.google.code.gson:gson:2.7",

            // ------------- Test dependencies -------------
            junit              : "junit:junit:4.12",
            truth              : "com.google.truth:truth:0.28",
            robolectric        : "org.robolectric:robolectric:3.1.2",
            openglApi          : "org.khronos:opengl-api:gl1.1-android-2.1_r1",
            compiletesting     : "com.google.testing.compile:compile-testing:0.9",
            autoservice        : "com.google.auto.service:auto-service:1.0-rc2",
            autocommon         : "com.google.auto:auto-common:0.6",
    ]
}
`</pre>

有了这个规范，那么我们在app下的build.gradle文件就可以这样来引用了

<pre>`android {
    compileSdkVersion rootProject.ext.android.compileSdkVersion
    buildToolsVersion rootProject.ext.android.buildToolsVersion

    defaultConfig {
        applicationId    "xxx"
        minSdkVersion    rootProject.ext.android.minSdkVersion
        targetSdkVersion rootProject.ext.android.targetSdkVersion
        versionCode      rootProject.ext.android.versionCode
        versionName      rootProject.ext.android.versionName
    }

    signingConfigs {
        myConfig {
            storeFile     file("../sign/kmkey")
            storePassword rootProject.ext.signingConfig.storePassword
            keyAlias      rootProject.ext.signingConfig.keyAlias
            keyPassword   rootProject.ext.signingConfig.keyPassword
        }
    }
    ...
}

dependencies {
    compile     fileTree(dir: 'libs', include: ['*.jar'])
    compile     rootProject.ext.deps.supportV7
    testCompile rootProject.ext.deps.junit
}
```

是不是一劳永逸了，今后修改版本只需要修改根目录下的config.gradle文件即可，希望可以对你们的Android开发规范有所帮助。

***

PS：还需要在项目根目录的build.gradle的最上方添加一句引用`apply from: "config.gradle"`