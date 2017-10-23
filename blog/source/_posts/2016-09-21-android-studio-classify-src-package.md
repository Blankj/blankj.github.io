---
title: Android Studio下对资源进行分包
date: 2016-09-21 11:29:46
categories:
  - Android
  - Android Tips
tags:
  - Android Tips
---

也许你曾经用过这个插件来对资源文件layout进行分类[https://github.com/dmytrodanylyk/folding-plugin](https://github.com/dmytrodanylyk/folding-plugin)，但如今随着AS版本的升高，该插件已经不再支持。

now，我来分享一种比这个插件更加优秀合理地对资源进行分包，让你的资源文件看起来简明有序。
先上效果图，如下所示：

{% asset_img classify_dir.png [效果图] %}

<!-- more -->

从图中可以看出，我们可以对每个页面的资源都进行具体分类，不只是layout，还有drawable及value，是不是心动了，赶紧照着如下配置试一试吧，别再让资源文件们“混为一潭”了。
方法很简单，配置我们的app文件夹下的build.gradle文件，比如我的

```
android {
    ...
    sourceSets {
        main {
            res.srcDirs =
                    [
                            'src/main/res/layouts',
                            'src/main/res',
                            'src/main/res/layouts/home',
                            'src/main/res/layouts/hot_sale',
                            'src/main/res/layouts/amuse',
                            'src/main/res/layouts/delicacy',
                            'src/main/res/layouts/food_management',
                            'src/main/res/layouts/settings',
                    ]
        }
    }
}
```

配置完之后，sync project一下就成功了。

***

**补充：发现有小伙伴试验不成功，好伐，是我疏漏了，文件夹是要自己创建的，因为自己创建的文件夹gradle不能解析为资源文件来使用，所以需要在build.gradle中进行配置，这样你们就懂了吧。**